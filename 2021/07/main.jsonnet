local input = importstr 'input.txt';
// local input = importstr 'test.txt';

// max :: [Int] -> Int
local max(ns) =
  std.foldr(
    function(i, acc) if ns[i] > acc then ns[i] else acc,
    std.range(0, std.length(ns) - 1),
    -(std.pow(2, 53)) + 1,
  )
;

// min :: [Int] -> Int
local min(ns) =
  std.foldr(
    function(i, acc) if ns[i] < acc then ns[i] else acc,
    std.range(0, std.length(ns) - 1),
    std.pow(2, 53) - 1,
  )
;

// data Crab = Int
local crabs = std.map(std.parseInt, std.split(std.rstripChars(input, '\n'), ','));

local distributionSize = max(crabs) + 1;
local initialDistribution = std.makeArray(distributionSize, function(i) 0);

// linearCost :: Int -> Int
local linearCost(distance) = distance;

// triangularCost :: Int -> Int
local triangularCost(distance) = (distance / 2) * (distance + 1);

// weight :: (Int -> Int) -> Int -> [Int] -> Int
local weight(costFn) = function(i) function(distribution)
  std.foldr(
    function(j, acc) distribution[j] * costFn(std.abs(j - i)) + acc,
    std.range(0, distributionSize - 1),
    0,
  )
;

// distribute :: [Int] -> [Int]
local distribute(crabs) =
  std.foldr(
    function(crab, acc) std.makeArray(distributionSize, function(i) if crab == i then acc[i] + 1 else acc[i]),
    crabs,
    initialDistribution,
  )
;

// weights :: (Int -> Int) -> [Int] -> [Int]
local weights(costFn) = function(distribution)
  std.map(
    function(i) weight(costFn)(i)(distribution),
    std.range(0, distributionSize - 1),
  )
;

// minWithIndex :: [Int] -> (Int, Int)
local minWithIndex(weights) =
  std.foldr(
    function(i, acc) if weights[i] < acc[1] then [i, weights[i]] else acc,
    std.range(0, distributionSize - 1),
    [-1, std.pow(2, 53) - 1],
  )
;

[
  minWithIndex(weights(linearCost)(distribute(crabs))),
  minWithIndex(weights(triangularCost)(distribute(crabs))),
]
