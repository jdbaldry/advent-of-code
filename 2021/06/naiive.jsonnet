local input = importstr 'input.txt';
local input = importstr 'test.txt';

// data LanternFish = Int
// data Fish = [LanternFish]

local fish = std.map(std.parseInt, std.split(std.rstripChars(input, '\n'), ','));

// tick :: Fish -> Fish
local tick(fish) = std.foldr(
  function(lf, acc)
    if lf == 0 then
      [6, 8] + acc
    else
      [lf - 1] + acc,
  fish,
  [],
)
;

// tickN :: Int -> Fish -> Fish
local tickN(n) = function(fish)
  std.foldr(
    function(_, acc) tick(acc),
    std.range(1, n),
    fish
  )
;

// gen :: Int -> [Int]
local gen(n) =
  std.map(
    function(i) std.length(tickN(n)(i)),
    std.makeArray(8, function(i) [i]),
  )
;

// fishCount :: Int -> Fish -> Int
local fishCount(generation) = function(fish)
  local genX = gen(generation);
  std.foldr(
    function(i, acc) genX[i] + acc,
    fish,
    0,
  )
;

std.map(
  function(i) fishCount(i)([0, 1, 2, 3, 4, 5, 6, 7]),
  std.range(0, 100),
)
