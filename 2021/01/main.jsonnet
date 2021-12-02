// TODO: 2021/01/part-two can be solved without a sliding window.
// When checking if B + C + D > A + B + C, you only actually need
// to compare D > A since B + C is common to both.
local input = importstr 'input.txt';
// local input = importstr 'test.txt';

// depths is the parsed input of measured depths.
local depths = std.map(std.parseInt, std.filter(function(line) line != '', std.split(input, '\n')));

// zip zips the elements of two arrays.
local zip(xs, ys) = [[xs[i], ys[i]] for i in std.range(0, std.length(xs) - 1)];

// zip3 zips the elements of three arrays.
local zip3(xs, ys, zs) = [[xs[i], ys[i], zs[i]] for i in std.range(0, std.length(xs) - 1)];

// sum sums xs.
local sum(xs) = std.foldr(function(x, acc) acc + x, xs, 0);

// lshift returns a new array with elements shifted to the left 'k' times.
// The right hand side is padded with 'k' zeros.
local lshift(xs, k) =
  local length = std.length(xs);
  std.makeArray(length, function(i) if i >= length - k then 0 else xs[i + k])
;

// windows is the sum of a three measurement sliding windows of depths.
local windows = std.map(sum, zip3(
  depths,
  lshift(depths, 1),
  lshift(depths, 2),
))
;

// countIncreases counts the number of times the x in xs is greater than the previous x.
// That is when xs[i] > xs[i-1].
local countIncreases(xs) =
  std.foldr(
    function(pair, acc) if pair[1] > pair[0] then acc + 1 else acc,
    zip(xs, lshift(xs, 1)),
    0,
  )
;

[
  countIncreases(depths),
  countIncreases(windows),
]
