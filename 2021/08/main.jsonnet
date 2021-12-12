local input = importstr 'input.txt';
// local input = importstr 'test.txt';

local util = import '../../lib/util.libsonnet',
      cut = util.cut,
      lines = util.lines,
      sortStr = util.sortStr,
      sum = util.sum,
      words = util.words
;

local ls = lines(input);

// data Pattern = Ord String
// data Display = String
// data ParsedLine = ([Pattern], [Digit])
// [([Pattern], [Digit])]
local parsedLines =
  std.map(
    function(l)
      local split = std.split(l, '|');
      [std.sort(words(split[0]), keyF=function(e) std.length(e)), words(split[1])],
    ls
  );

// patterns :: ParsedLine -> [Pattern]
local patterns(parsedLine) = parsedLine[0];
// display :: ParsedLine -> [Digit]
local display(parsedLine) = parsedLine[1];

// one :: [Pattern] -> Pattern
local one(ps) = ps[0];
local four(ps) = ps[2];
local seven(ps) = ps[1];
local eight(ps) = ps[9];
// twoThreeFive are the patterns of length 5.
local twoThreeFive(ps) = ps[3:6];
// zeroSixNine are the patterns of length 6.
local zeroSixNine(ps) = ps[6:9];

// three has both segments of one present.
local three(ps) =
  local three = std.filter(function(p) std.length(cut(one(ps))(p)) == 0, twoThreeFive(ps));
  assert std.length(three) == 1 : 'incorrect deduction of three';
  three[0]
;

// twoFive are the patterns of length 5 that are not three.
local twoFive(ps) =
  local twoFive = std.filter(function(p) p != three(ps), twoThreeFive(ps));
  assert std.length(twoFive) == 2 : 'incorrect deduction of twoFive';
  twoFive
;

// nine shares all segments with three.
local nine(ps) =
  local nine = std.filter(function(p) std.length(cut(p)(three(ps))) == 1, zeroSixNine(ps));
  assert std.length(nine) == 1 : 'incorrect deduction of nine';
  nine[0]
;

// zeroSix are the patterns of length 6 that are not nine.
local zeroSix(ps) =
  local zeroSix = std.filter(function(p) p != nine(ps), zeroSixNine(ps));
  assert std.length(zeroSix) == 2 : 'incorrect deduction of zeroSix';
  zeroSix
;

// five shares all segments with nine.
local five(ps) =
  local five = std.filter(function(p) std.length(cut(p)(nine(ps))) == 0, twoFive(ps));
  assert std.length(five) == 1 : 'incorrect deduction of five';
  five[0]
;

// two is not five.
local two(ps) =
  local two = std.filter(function(p) p != five(ps), twoFive(ps));
  assert std.length(two) == 1 : 'incorrect deduction of two';
  two[0]
;

// zero is the pattern of length 6, that is not nine and that has both segments of one.
local zero(ps) =
  local zero = std.filter(function(p) std.length(cut(one(ps))(p)) == 0, zeroSix(ps)); assert std.length(zero) == 1 : 'incorrect deduction of zero';
  zero[0]
;

// six is not zero.
local six(ps) =
  local six = std.filter(function(p) p != zero(ps), zeroSix(ps));
  assert std.length(six) == 1 : 'incorrect deduction of six';
  six[0]
;

[
  std.length(
    std.filter(
      function(digit)
        local l = std.length(digit);
        l == 2 || l == 3 || l == 4 || l == 7,
      std.flatMap(display, parsedLines)
    )
  ),

  sum(
    std.map(
      function(pl)
        local ps = patterns(pl);
        local table =
          {
            [sortStr(zero(ps))]: 0,
            [sortStr(one(ps))]: 1,
            [sortStr(two(ps))]: 2,
            [sortStr(three(ps))]: 3,
            [sortStr(four(ps))]: 4,
            [sortStr(five(ps))]: 5,
            [sortStr(six(ps))]: 6,
            [sortStr(seven(ps))]: 7,
            [sortStr(eight(ps))]: 8,
            [sortStr(nine(ps))]: 9,
          };
        std.parseInt(std.join('', std.map(function(digit) '%s' % table[sortStr(digit)], display(pl)))),
      parsedLines
    )
  ),
]
