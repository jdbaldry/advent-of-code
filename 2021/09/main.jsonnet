local input = importstr 'input.txt';
local input = importstr 'test.txt';

local util = import '../../lib/util.libsonnet',
      sum = util.sum
;

local notEmpty(str) = std.length(str) != 0;
local heights = std.map(function(line) std.map(std.parseInt, std.stringChars(line)), std.filter(notEmpty, std.split(input, '\n')));

local display(heights) = std.join('\n', std.map(function(row) std.join('', std.map(function(p) '%3s' % [p], row)), heights));

// partOne :: [[Int]] -> Int
local partOne(heights) =
  // adjacent :: [[Int]] -> Int -> Int -> [Int]
  local adjacent(heights) =
    assert std.length(heights) > 0 : 'heights must have at least one row';
    function(x) function(y)
      local maxY = std.length(heights) - 1;
      local maxX = std.length(heights[0]) - 1;
      std.prune([
        if x != 0 then heights[y][x - 1],
        if x != maxX then heights[y][x + 1],
        if y != 0 then heights[y - 1][x],
        if y != maxY then heights[y + 1][x],
      ])
  ;

  // lessThan :: [Int] -> Int -> Bool
  // lessThan returns true iff x is less than all ys.
  local lessThan(ys) = function(x) std.foldr(function(y, acc) acc && x < y, ys, true);

  sum(
    std.flattenArrays(
      std.mapWithIndex(
        function(y, hs)
          std.mapWithIndex(
            function(x, h)
              if lessThan(adjacent(heights)(x)(y))(h) then
                h + 1
              else 0,
            hs
          ),
        heights,
      )
    )
  )
;

// partTwo :: [[Int]] -> Int
local partTwo(heights) =
  // If monus is truncated subtraction, this is truncated addition where truncation happens at 'limit'.
  // https://en.wikipedia.org/wiki/Monus
  // inverseMonus :: Int -> Int -> Int -> Int
  local inverseMonus(limit) = function(x) function(addend) local sum = x + addend; if sum >= limit then limit else sum;
  local monus(minuend) = function(subtrahend) local difference = minuend - subtrahend; if difference < 0 then 0 else difference;
  // incOne :: Int -> Int
  local incOne(x) = inverseMonus(10)(1)(x);
  // decOne :: Int -> Int
  local decOne(x) = monus(x)(1);
  // inc :: [[Int]] -> [[Int]]
  local inc(heights) =
    std.makeArray(
      std.length(heights), function(y)
        std.makeArray(std.length(heights[y]), function(x) incOne(heights[y][x]))
    )
  ;
  // dec :: [[Int]] -> [[Int]]
  local dec(heights) =
    std.makeArray(
      std.length(heights), function(y)
        std.makeArray(std.length(heights[y]), function(x) decOne(heights[y][x]))
    )
  ;
  // Set all points lower than the point at heights[y][x] to the max value.
  // normalize :: Int -> Int -> [[Int]] -> [[Int]]
  local normalizeUp(x) = function(y) function(heights)
    local point = heights[y][x];
    std.makeArray(
      std.length(heights), function(y1)
        std.makeArray(std.length(heights[y1]), function(x1) if heights[y1][x1] < point then 9 else heights[y1][x1])
    )
  ;
  // Set all points lower than the point at heights[y][x] to the min value.
  // normalize :: Int -> Int -> [[Int]] -> [[Int]]
  local normalizeDown(x) = function(y) function(heights)
    local point = heights[y][x];
    std.makeArray(
      std.length(heights), function(y1)
        std.makeArray(std.length(heights[y1]), function(x1) if heights[y1][x1] < point then 0 else heights[y1][x1])
    )
  ;
  // incWhile :: ([[Int]] -> Bool) -> [[Int]] -> [[Int]]
  local incWhile(pred=function(heights) true, heights) = if pred(heights) then incWhile(pred, inc(heights)) else heights;
  // decWhile :: ([[Int]] -> Bool) -> [[Int]] -> [[Int]]
  local decWhile(pred=function(heights) true, heights) = if pred(heights) then decWhile(pred, dec(heights)) else heights;
  // notMaxed :: Int -> Int -> [[Int]] -> ([[Int]] -> Bool)
  local notMaxed(x) = function(y) function(heights) heights[y][x] != 9;
  // notZero :: Int -> Int -> [[Int]] -> ([[Int]] -> Bool)
  local notZero(x) = function(y) function(heights) heights[y][x] != 0;
  // std.length(std.filter(function(x) x != 10, std.flattenArrays(incWhile(notMaxed(1)(0), normalize(1)(0)(heights)))))
  // normalize(1)(0)(heights)
  // display(decWhile(notZero(0)(0), normalize(0)(0)(heights)))

  local from(x) = function(y) function(heights)
    local point = heights[y][x];
    std.makeArray(
      std.length(heights), function(y1)
        std.makeArray(std.length(heights[y1]), function(x1) (std.abs(x - x1) + std.abs(y - y1)) * (std.abs(x - x1) + std.abs(y - y1)))
    )
  ;

  local basin(x) = function(y) function(heights)
    local point = heights[y][x];
    std.makeArray(
      std.length(heights), function(y1)
        std.makeArray(std.length(heights[y1]), function(x1) std.abs(x - x1) + std.abs(y - y1))
    )
  ;
  local x = 9, y = 0;
  std.join(
    '\n\n',
    [
      display(
        heights
      ),
      display(
        from(x)(y)(heights)
      ),
      display(
        basin(x)(y)(from(x)(y)(heights))
      ),
    ]
  )
;

//partOne(heights),
partTwo(heights)
