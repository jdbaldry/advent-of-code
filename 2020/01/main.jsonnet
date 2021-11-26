local input = importstr 'input.txt';
// local input = std.join('\n', ['1721', '979', '366', '299', '675', '1456']);

local lines = std.filter(function(line) std.length(line) > 0, std.split(input, '\n'));
local expenses = std.map(std.parseInt, lines);

// product returns the product of the numbers in 'numbers'.
local product(numbers=[]) =
  std.foldr(
    function(number, acc) acc * number,
    numbers,
    1
  )
;

// sumOfK returns the first tuple of length 'k' where all values of the tuple are in the set
// of 'numbers' and the sum of those numbers is equal to 'sum'.
// If no tuple is found, it returns null.
local sumOfK(numbers, sum, k) =
  if k == 1 then
    if std.member(numbers, sum) then [sum] else null
  else
    local addends =
      std.prune(std.map(
        function(n)
          local complement = sumOfK(numbers, sum - n, k - 1);
          if complement != null then complement + [n],
        numbers
      ));
    if std.length(addends) != 0 then addends[0]
;

[
  product(sumOfK(expenses, 2020, 2)),
  product(sumOfK(expenses, 2020, 3)),
]
