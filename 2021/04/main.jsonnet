local input = importstr 'input.txt';
// local input = importstr 'test.txt';

local pc = import '../../2020/02/parser-combinators.libsonnet';

// last :: [a] -> a
local last(arr) =
  local l = std.length(arr);
  if l == 0 then [] else arr[l - 1]
;

// split :: String -> String -> [String]
local split(str, splitstr) =
  local splits = std.stringChars(splitstr);
  local l = std.length(str);
  local _l = std.length(splits);
  local aux(str, splits, i, j, k, acc) =
    if j == l - 1 then
      acc + [str[i:]]
    else if k == _l - 1 && str[j] == splits[k] then
      aux(str, splits, j + 1, j + 1, 0, acc + [str[i:j]]) tailstrict
    else if str[j] == splits[k] then
      aux(str, splits, i, j + 1, k + 1, acc) tailstrict
    else
      aux(str, splits, i, j + 1, 0, acc) tailstrict
  ;
  aux(str, splits, 0, 0, 0, [])
;

local data = split(input, '\n\n');

local draw =
  pc.seq([
    pc.star(
      pc.concat(
        pc.capture(pc.plus(pc.digit)),
        pc.char(',')
      )
    ),
    pc.capture(pc.plus(pc.digit)),
    pc.ws,
  ])(data[0], pc.init)
;

local draws = std.length(draw);

local boards =
  std.map(
    function(line)
      std.map(
        function(row) std.map(function(n) [std.parseInt(n), false], std.filter(function(str) std.length(str) != 0, std.split(row, ' '))),
        std.filter(function(str) std.length(str) != 0, std.split(line, '\n'))
      ),
    data[1:],
  )
;

// Square :: [Int Bool]
// [ number marked ]
//
// Line :: [Square]
//
// Board :: [Line]
// [
//   [ [n m] [n m] [n m] [n m ] [n m] ]
//   [ [n m] [n m] [n m] [n m ] [n m] ]
//   [ [n m] [n m] [n m] [n m ] [n m] ]
//   [ [n m] [n m] [n m] [n m ] [n m] ]
//   [ [n m] [n m] [n m] [n m ] [n m] ]
// ]

// transpose :: [[a]] -> [[a]]
local transpose(arr) =
  assert std.isArray(arr) : 'not array: `transpose`';
  if std.length(arr) == 0 then
    []
  else if std.length(arr[0]) == 0 then
    transpose(arr[1:])
  else
    local x = arr[0][0];
    local xs = arr[0][1:];
    local xss = arr[1:];
    [
      [x] + [xs[0] for xs in xss if std.length(xs) > 0],
    ]
    + transpose(
      [xs] + [xs[1:] for xs in xss if std.length(xs) > 0]
    )
;

// rows :: Board -> [Line]
local rows(board) = board;

// cols :: Board -> [Line]
local cols(board) = transpose(board);

// completeLine :: Line -> Bool
local completeLine(line) = std.foldr(function(square, acc) acc && square[1], line, true);

// anyCompleteLine :: [Line] -> Bool
local anyCompleteLine(lines) = std.foldr(function(line, acc) acc || completeLine(line), lines, false);

// wins :: Board -> Bool
local wins(board) = anyCompleteLine(rows(board)) || anyCompleteLine(cols(board));

local board = [
  [[1, true], [2, false], [3, false], [4, false], [5, false]],
  [[6, true], [7, true], [8, true], [9, true], [10, false]],
  [[11, false], [12, true], [13, true], [14, true], [15, false]],
]
;

// markSquare :: Int -> Square -> Square
local markSquare(n) = function(square) if square[0] == n then [n, true] else square;

// markLine :: Int -> Line -> Line
local markLine(n) = function(line) std.map(markSquare(n), line);

// markBoard :: Int -> Board -> Board
local markBoard(n) = function(board) std.map(markLine(n), rows(board));

// play :: Int -> [Board] -> [Board]
local play(n) = function(boards) std.map(
  markBoard(n),
  boards,
)
;

local partOne = std.foldl(
  function(acc, n)
    local winners = acc[1];
    if std.length(winners) != 0 then
      [acc[0], winners, acc[2]]
    else
      local boards = std.map(function(board) markBoard(n)(board), acc[0]);
      local winners = std.filter(wins, boards);
      [boards, winners, n],
  std.map(std.parseInt, draw[draws - 1].captured),
  [boards, [], -1]
);

local partTwo = std.foldl(
  function(acc, n)
    local prev = acc[0];
    local last = acc[1];
    local draw = acc[2];
    if std.length(prev) == 0 then
      [prev, last, draw]
    else
      local boards = std.filter(function(board) !wins(board), std.map(function(board) markBoard(n)(board), prev));
      if std.length(boards) == 0 then
        [boards, [markBoard(n)(prev[0])], n]
      else
        [boards, last, n],
  std.map(std.parseInt, draw[draws - 1].captured),
  [boards, [], -1]
);

// sumSquares :: [Square] -> Int
local sumSquares(squares) = std.foldr(function(square, acc) acc + square[0], squares, 0);

// unmarkedSquare :: Square -> Bool
local unmarkedSquare(square) = !square[1];

// unmarked :: Board -> [Int]
local unmarked(board) =
  std.filter(unmarkedSquare, std.join([], board))
;

[
  sumSquares(unmarked(partOne[1][0])) * partOne[2],
  sumSquares(unmarked(partTwo[1][0])) * partTwo[2],
]
