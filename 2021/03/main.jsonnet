local input = importstr 'input.txt';
// local input = importstr 'test.txt';

local numbers = std.filter(function(line) std.length(line) != 0, std.split(input, '\n'));

local parseBinary(str) =
  local length = std.length(str);
  std.foldr(
    function(pair, acc) acc + std.pow(2, pair[0]) * std.parseInt(pair[1]),
    std.makeArray(length, function(i) [length - 1 - i, str[i]]),
    0,
  )
;

local length = std.length(numbers);
local digits = std.length(numbers[0]);

local counts =
  std.foldr(
    function(number, acc)
      std.mapWithIndex(function(i, e) acc[i] + std.parseInt(e), std.stringChars(number)),
    numbers,
    std.makeArray(digits, function(i) 0),
  )
;

local diags =
  std.foldr(
    function(n, acc) std.mapWithIndex(
      function(i, m) m + n[i],
      acc,
    ),
    std.mapWithIndex(
      function(i, n)
        [
          std.pow(2, digits - 1 - i) * if n > length / 2 then 1 else 0,
          std.pow(2, digits - 1 - i) * if n < length / 2 then 1 else 0,
        ],
      counts,
    ),
    [0, 0]
  )
;

local product(xs) = std.foldr(function(x, acc) acc * x, xs, 1);

local findO2(numbers) =
  local aux(numbers, i) =
    local length = std.length(numbers);
    if length == 0 then
      error 'must have at least one number'
    else if length == 1 then
      numbers[0]
    else
      local ones = std.foldr(function(number, acc) if number[i] == '1' then acc + 1 else acc, numbers, 0);
      local mostCommon = if ones >= length / 2 then '1' else '0';
      aux(std.filter(function(number) number[i] == mostCommon, numbers), i + 1);
  aux(numbers, 0)
;

local findCO2(numbers) =
  local aux(numbers, i) =
    local length = std.length(numbers);
    if length == 0 then
      error 'must have at least one number'
    else if length == 1 then
      numbers[0]
    else
      local ones = std.foldr(function(number, acc) if number[i] == '1' then acc + 1 else acc, numbers, 0);
      local mostCommon = if ones >= length / 2 then '1' else '0';
      aux(std.filter(function(number) number[i] != mostCommon, numbers), i + 1);
  aux(numbers, 0)
;

local compose(f, g) = function(x) f(g(x));

[
  product(diags),
  local O2 = compose(parseBinary, findO2)(numbers);
  local CO2 = compose(parseBinary, findCO2)(numbers);
  product([O2, CO2]),
]
