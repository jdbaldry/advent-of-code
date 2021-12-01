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

// window is the sum of a three measurement sliding window of depths.
local window = std.map(sum, zip3(depths, depths[1:] + [0], depths[2:] + [0, 0]));

// countIncreases counts the number of times the x in xs is greater than the previous x.
// That is when x[i] > x[i-1].
local countIncreases(xs) =
  std.foldr(
    function(pair, acc) if pair[1] > pair[0] then acc + 1 else acc,
    zip(xs, xs[1:] + [0]),
    0,
  )
;

[
  countIncreases(depths),
  countIncreases(window),
]
