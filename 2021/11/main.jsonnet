local input = importstr 'input.txt';
// local input = importstr 'test.txt';
local util = import '../../lib/util.libsonnet',
      lines = util.lines
;

// data Point = (Int, Int)
// data Octopus = Int
// data Octopuses = [[Octopus]]
local octopuses = std.map(function(line) std.map(std.parseInt, std.stringChars(line)), lines(input));

// display :: Octopuses -> String
local display(oss) = std.join('\n', [''] + std.map(function(os) std.join(' ', std.map(function(o) '%3d' % [o], os)), oss));

// adjacent returns all adjacent points including diagonally adjacent points.
// adjacent :: Octopuses -> Point -> [Point]
local adjacent(oss) =
  assert std.length(oss) > 0 : 'oss must have at least one row';
  function(p)
    local x = p[0];
    local y = p[1];
    local maxY = std.length(oss) - 1;
    local maxX = std.length(oss[0]) - 1;
    std.prune([
      if x != 0 && y != 0 then [x - 1, y - 1],
      if x != 0 && y != maxY then [x - 1, y + 1],
      if x != maxX && y != 0 then [x + 1, y - 1],
      if x != maxX && y != maxY then [x + 1, y + 1],
      if x != 0 then [x - 1, y],
      if y != maxY then [x, y + 1],
      if x != maxX then [x + 1, y],
      if y != 0 then [x, y - 1],
    ])
;

// update :: Octopuses -> Point -> (Point -> Point) -> Octopuses
local update(oss) = function(p) function(f=function(p) p)
  local x = p[0], y = p[1];
  std.makeArray(
    std.length(oss),
    function(y1) std.makeArray(
      std.length(oss[y]),
      function(x1) if x == x1 && y == y1 then f(oss[y1][x1]) else oss[y1][x1],
    )
  )
;

// incOne :: Octopus -> Octopus
local incOne(o) = o + 1;
// readyOne :: Octopus -> Bool
local readyOne(o) = o == 10;
// resetOne :: Octopus -> Octopus
local resetOne(o) = if o < 0 then 0 else o;

// inc :: Octopuses -> Octopuses
local inc(oss) =
  std.makeArray(
    std.length(oss),
    function(y) std.makeArray(
      std.length(oss[y]),
      function(x) incOne(oss[y][x])
    )
  )
;

// flash :: Point -> Octopuses -> Octopuses
local flash(p, oss) =
  local x = p[0], y = p[1];
  if oss[y][x] >= 10 then
    local updated = update(oss)(p)(function(_) -100);
    local incremented =
      std.foldr(
        function(p, acc) update(acc)(p)(incOne),
        adjacent(updated)(p),
        updated
      )
    ;
    std.foldr(
      function(p, acc) flash(p, acc),
      adjacent(incremented)(p),
      incremented
    )
  else
    oss
;

// reset :: Octopuses -> Octopuses
local reset(oss) =
  std.makeArray(
    std.length(oss),
    function(y) std.makeArray(
      std.length(oss[y]),
      function(x) resetOne(oss[y][x])
    )
  )
;

// step :: Octopuses -> Octopuses
local step(oss) =
  // ready :: Octopuses -> [Point]
  local ready(oss) =
    std.prune(std.flattenArrays(std.makeArray(
      std.length(oss),
      function(y) std.makeArray(
        std.length(oss[y]),
        function(x) if readyOne(oss[y][x]) then [x, y] else null
      )
    )))
  ;
  local incremented = inc(oss);
  local flashed = std.foldr(flash, ready(incremented), incremented);
  reset(flashed)
;

// flashed :: Octopuses -> Int
local flashed(oss) = std.length(std.filter(function(o) o == 0, std.flattenArrays(oss)));

local stepN(n) = function(oss)
  if n == 0 then
    oss
  else stepN(n - 1)(step(oss))
;

local partOne() = std.foldr(
  function(_, acc)
    local stepped = step(acc[1]);
    [flashed(stepped) + acc[0], stepped],
  std.range(1, 100),
  [0, octopuses]
)[0]
;

local partTwo() =
  local all = std.length(std.flattenArrays(octopuses));
  local aux(n, acc) =
    if flashed(acc) == all then
      n
    else
      aux(n + 1, step(acc))
  ;
  aux(0, octopuses)
;

[
  partOne(),
  partTwo(),
]
