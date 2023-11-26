local input = importstr 'input.txt';
// local input = importstr 'test.txt';

local pc = import '../../lib/parser-combinators.libsonnet';

// data Point = (Int, Int)
// data Segment = (Point, Point)
// data EncodedPoint = String
// data Board = { encodedPoint : Int }

// x :: Point -> Int
local x(p) = p[0];

// minX :: Segment -> Int
local minX(s) = std.min(x(s[0]), x(s[1]));

// maxX :: Segment -> Int
local maxX(s) = std.max(x(s[0]), x(s[1]));

// y :: Point -> Int
local y(p) = p[1];

// minY :: Segment -> Int
local minY(s) = std.min(y(s[0]), y(s[1]));

// maxY :: Segment -> Int
local maxY(s) = std.max(y(s[0]), y(s[1]));

// zip :: [a] -> [a] -> [(a, a)]
local zip(xs) = function(ys) [[xs[i], ys[i]] for i in std.range(0, std.length(xs) - 1)];

// points :: Segment -> [Point]
local points(s) =
  local comp(a) = function(b) if a == b then 0 else if a < b then 1 else -1;
  local x1 = x(s[0]);
  local y1 = y(s[0]);
  local x2 = x(s[1]);
  local y2 = y(s[1]);
  local dx = comp(x1)(x2);
  local dy = comp(y1)(y2);
  [[x1 + n * dx, y1 + n * dy] for n in std.range(0, std.max(std.abs(x2 - x1), std.abs(y2 - y1)))]
;

// last :: [a] -> a
local last(arr) =
  local l = std.length(arr);
  if l == 0 then [] else arr[l - 1]
;

local lines = std.filter(function(line) std.length(line) != 0, std.split(input, '\n'));

// stringToSegment :: String -> [(Int, Int) (Int Int)]
local stringToSegment(str) =
  local captured = std.map(std.parseInt, last(pc.seq([
    pc.capture(pc.star(pc.digit)),
    pc.char(','),
    pc.capture(pc.star(pc.digit)),
    pc.string(' -> '),
    pc.capture(pc.star(pc.digit)),
    pc.char(','),
    pc.capture(pc.star(pc.digit)),
  ])(str, pc.init)).captured);
  [captured[0:2], captured[2:]]
;


// horizontal :: Segment -> Bool
local horizontal(s) = minY(s) == maxY(s);

// vertical :: Segment -> Bool
local vertical(s) = minX(s) == maxX(s);

// cardinal :: Segment -> Bool
local cardinal(s) = horizontal(s) || vertical(s);

// cross :: Point -> Point -> Int
local cross(p1) = function(p2) x(p1) * y(p2) - y(p1) * x(p2);

// sub :: Point -> Point -> Int
local sub(p1) = function(p2) [x(p1) - x(p2), y(p1) - y(p2)];

// onSegment  :: Point -> Segment -> Bool
local onSegment(p) = function(l)
  x(p) <= std.max(x(l[0]), x(l[1])) && x(p) >= std.min(x(l[0]), x(l[1]))
  &&
  y(p) <= std.max(y(l[0]), y(l[1])) && y(p) >= std.min(y(l[0]), y(l[1]))
;

// orientation :: Point -> Point -> Point -> Int
local orientation(p1) = function(p2) function(p3)
  local o = (y(p2) - y(p1)) * (x(p3) - x(p2))
            -
            (x(p2) - x(p1)) * (y(p3) - y(p2));
  if o == 0 then 0 else if o < 0 then -1 else 1
;

// colinear :: Segment -> Segment -> Bool
local colinear(s1) = function(s2)
  local o1 = orientation(s1[0])(s1[1])(s2[0]);
  local o2 = orientation(s1[0])(s1[1])(s2[1]);
  local o3 = orientation(s1[1])(s2[1])(s1[0]);
  local o4 = orientation(s1[1])(s2[1])(s1[1]);

  (o1 == 0 && onSegment(s2[0])(s1))
  ||
  (o2 == 0 && onSegment(s2[1])(s1))
  ||
  (o3 == 0 && onSegment(s1[0])(s2))
  ||
  (o4 == 0 && onSegment(s1[1])(s2))
;

// intersect :: Segment -> Segment -> Bool
// https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
local intersect(s1) = function(s2)
  local o1 = orientation(s1[0])(s1[1])(s2[0]);
  local o2 = orientation(s1[0])(s1[1])(s2[1]);
  local o3 = orientation(s1[1])(s2[1])(s1[0]);
  local o4 = orientation(s1[1])(s2[1])(s1[1]);

  // General case.
  (o1 != o2 && o3 != o4)
  ||
  // Special cases.
  (o1 == 0 && onSegment(s2[0])(s1))
  ||
  (o2 == 0 && onSegment(s2[1])(s1))
  ||
  (o3 == 0 && onSegment(s1[0])(s2))
  ||
  (o4 == 0 && onSegment(s1[1])(s2))
;

// intersections :: [Segment] -> Int
local intersections(ss) =
  if std.length(ss) <= 1 then
    0
  else
    local head = ss[0];
    local tail = ss[1:];
    std.foldr(
      function(v, acc)
        local intersects = intersect(head)(v);
        if intersects then acc + 1 else acc,
      tail,
      0
    ) + intersections(tail)
;

// overlap :: Segment -> Segment -> Int
// If segments intersect but are not colinear, there is exactly one point of overlap.
// If they intersect and are colinear, they may overlap by a number of points.
local overlap(s1) = function(s2)
  if intersect(s1)(s2) then
    if colinear(s1)(s2) then
      if horizontal(s1) then
        local min1 = std.min(x(s1[0]), x(s1[1]));
        local max1 = std.max(x(s1[0]), x(s1[1]));
        local min2 = std.min(x(s2[0]), x(s2[1]));
        local max2 = std.max(x(s2[0]), x(s2[1]));
        local diff = std.min(max1, max2) - std.max(min1, min2);
        if diff < 0 then 0 else if diff == 0 then 1 else diff
      else  // vertical
        local min1 = std.min(y(s1[0]), y(s1[1]));
        local max1 = std.max(y(s1[0]), y(s1[1]));
        local min2 = std.min(y(s2[0]), y(s2[1]));
        local max2 = std.max(y(s2[0]), y(s2[1]));
        local diff = std.min(max1, max2) - std.max(min1, min2);
        if diff < 0 then 0 else if diff == 0 then 1 else diff
    else 1
  else 0
;

// overlaps :: [Segment] -> Int
local overlaps(ss) =
  if std.length(ss) <= 1 then
    0
  else
    local head = ss[0];
    local tail = ss[1:];
    std.foldr(
      function(s, acc)
        acc + overlap(head)(s),
      tail,
      0
    ) + overlaps(tail)
;


// pointToEncodedPoint :: Point -> EncodedPoint
local pointToEncodedPoint(p) = '%s,%s' % p;

local board = {};

// plot :: Board -> [Segment] -> Board
local plot(board) = function(ss)
  std.foldr(
    function(s, board)
      if horizontal(s) then
        local y = minY(s);
        local min = minX(s);
        local max = maxX(s);
        board + {
          [pointToEncodedPoint([x, y])]+: 1
          for x in std.range(min, max)
        }
      else if vertical(s) then
        local x = minX(s);
        local min = minY(s);
        local max = maxY(s);
        board + {
          [pointToEncodedPoint([x, y])]+: 1
          for y in std.range(min, max)
        }
      else
        board + {
          [pointToEncodedPoint(p)]+: 1
          for p in points(s)
        },
    ss,
    board,
  )
;

// atLeastOneOverlap :: Board -> Int
local atLeastOneOverlap(board) =
  std.length(
    std.filter(function(field) board[field] > 1, std.objectFields(board))
  )
;

[
  atLeastOneOverlap(
    plot(board)(
      std.filter(
        cardinal,
        std.map(stringToSegment, lines)
      ),
    ),
  ),
  atLeastOneOverlap(
    plot(board)(
      std.map(stringToSegment, lines)
    )
  ),
]
