local input = importstr 'input.txt';
// local input = importstr 'test.txt';

// data Fish = { String: Int }
local init = {
  '0': 0,
  '1': 0,
  '2': 0,
  '3': 0,
  '4': 0,
  '5': 0,
  '6': 0,
  '7': 0,
  '8': 0,
};

local fish = std.foldr(
  function(str, acc) acc { [str]+: 1 },
  std.split(std.rstripChars(input, '\n'), ','),
  init,
)
;

// tick :: Fish -> Fish
local tick(fish) =
  {
    '0': fish['1'],
    '1': fish['2'],
    '2': fish['3'],
    '3': fish['4'],
    '4': fish['5'],
    '5': fish['6'],
    '6': fish['7'] + fish['0'],
    '7': fish['8'],
    '8': fish['0'],
  }
;

// tickN :: Int -> Fish -> Fish
local tickN(n) = function(fish)
  std.foldr(
    function(_, acc) tick(acc),
    std.range(1, n),
    fish
  )
;

// sum :: Fish -> Int
local sum(fish) = std.foldr(
  function(str, acc) fish[str] + acc,
  std.objectFields(fish),
  0,
)
;

[
  sum(tickN(80)(fish)),
  sum(tickN(256)(fish)),
]
