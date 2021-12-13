local input = importstr 'input.txt';
// local input = importstr 'test.txt';

local util = import '../../lib/util.libsonnet',
      lines = util.lines,
      stack = util.stack
;

// chunkOpen is the set of characters that open a new chunk.
local chunkOpen = ['[', '(', '{', '<'];

// chunkClose is the mapping of opening chunk characters to their
// respective closing character.
local chunkClose =
  {
    '[': ']',
    '(': ')',
    '{': '}',
    '<': '>',
  }
;

// data LexerState = LexerState {
//   input :: [Char],
//   i :: Int,
//   chunks :: Stack,
//   err :: Null | String
// }
local lexerState =
  {
    new(input): {
      input: input,
      i: 0,
      chunks: stack.new(100),
      err: null,
    },
  }
;

// lexC lexes a single character of a chunk.
// lexC :: LexerState -> LexerState
local lexC(prev) =
  local input = prev.input, i = prev.i;
  local chunks = prev.chunks;
  local err = prev.err;

  assert i >= 0 && i < std.length(input) : 'invalid index into input characters';

  local c = input[i];
  local want = if stack.size(chunks) != 0 then chunkClose[stack.peek(chunks)];

  if c == '\n' || c == ' ' || c == '\t' || c == '\r' then
    prev { i+: 1 }
  else
    local ls =
      if i == std.length(input) - 1 then
        if stack.size(chunks) == 0 then
          prev { err: 'EOF' }
        else
          prev { err: 'incomplete chunk' }
      else
        prev
    ;
    if std.member(chunkOpen, c) then
      ls { i+: 1, chunks: stack.push(chunks, c) }
    else if c == want then
      ls { i+: 1, chunks: stack.pop(chunks)[0] }
    else
      ls { err: 'invalid chunk char' }
;

// lex lexes until `err` is not `null`.
// Err can be 'EOF' which indicates successful lexing, otherwise it will
// represent the lexing error.
local lex(prev) =
  local err = prev.err;
  if err != null then
    prev
  else
    lex(lexC(prev))
;


local partOne() =
  // score is a mapping of incorrect chunk characters to their score.
  local score =
    {
      ')': 3,
      ']': 57,
      '}': 1197,
      '>': 25137,
    }
  ;
  std.foldr(
    function(c, acc) score[c] + acc,
    std.prune(std.map(
      function(line)
        local lexed = lex(lexerState.new(line));
        if lexed.err == 'invalid chunk char' then
          lexed.input[lexed.i],
      lines(input)
    )),
    0
  )
;

local partTwo() =
  // score is a mapping from closing chunk char to its autocorrect score.
  local score =
    {
      ')': 1,
      ']': 2,
      '}': 3,
      '>': 4,
    }
  ;
  // calculateScore calculates the score for the completion of the
  // incomplete chunks in `chunks`.
  local calculateScore(chunks) =
    std.foldr(
      function(c, acc) score[chunkClose[c]] + 5 * acc,
      chunks,
      0
    )
  ;
  // middle returns the middle element of a sorted array.
  // It only ever returns a single element and so does not work for arrays
  // of even length.
  // middle :: [a] -> a
  local middle(arr) =
    local l = std.length(arr);
    assert l % 2 != 0 : 'only arrays of odd length are supported';
    arr[std.floor(l / 2)]
  ;
  // scores are the ordered scores for the auto-complete of each
  // incomplete line of chunks.
  local scores =
    std.sort(
      std.map(
        calculateScore,
        std.prune(std.map(
          function(line)
            local lexed = lex(lexerState.new(line));
            if lexed.err == 'incomplete chunk' then
              lexed.chunks.stack,
          lines(input)
        )),
      )
    )
  ;
  middle(scores)
;

[
  partOne(),
  partTwo(),
]
