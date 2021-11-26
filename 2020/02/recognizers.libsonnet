// A recognizer is a parser that can recognize a valid string.
// Each recognizer has the signature:
// Recognizer :: (String, Int) -> [Int]
// Where the resulting array is the index after each parsed character.
// References:
// - https://qntm.org/combinators
// - https://en.wikipedia.org/wiki/Parser_combinator

// empty is a recognizer that will recognize the empty string ''.
local empty(str, i) = [i];

// never is a recognizer that always returns empty results.
local never(str, i) = [];

// satisfies returns a recognizer that recognizes any character for which
// the predicate 'pred' is satisfied.
local satisfies(pred=function(c) true) = function(str, i)
  local length = std.length(str);
  if length == 0 || i > length - 1 || !pred(str[i]) then
    []
  else [i + 1]
;

// any recognizes any character.
local any = satisfies();

// char returns a recognizer for char 'c'.
local char(c) = satisfies(function(_c) c == _c);

// a is a recognizer that can recognize the character 'a' in the string 'str'
// at index 'i'.
local a = char('a');

// b is a recognizer that can recognize the character 'b' in the string 'str'
// at index 'i'.
local b = char('b');

// digit is a recognizer for digits 0-9.
local digit = satisfies(function(c) '0' <= c && c <= '9');

// lower is a recognizer for lowercase alphabetic characters a-z.
local lower = satisfies(function(c) 'a' <= c && c <= 'z');

// alternate combines the results of two parsers.
local alternate(p1, p2) = function(str, i) p1(str, i) + p2(str, i);

// concat concatenates two parsers, feeding the results from the first into the
// input of the second.
local concat(p1, p2) = function(str, i)
  local parsed = p1(str, i);
  local l = std.length(parsed);
  if l == 0 then
    []
  else
    parsed + p2(str, parsed[l - 1])
;

// star is the Kleene star and applies the same parser zero or more times.
local star(p) = function(str, i)
  local parsed = p(str, i);
  local l = std.length(parsed);
  if l == 0 then
    parsed
  else
    parsed + star(p)(str, parsed[l - 1])
;

// seq concatenates all parsers in arr.
local seq(arr) = std.foldr(function(p, acc) concat(p, acc), arr, never);

// valid applies the parser 'p' and returns true iff the string 'str'
// is recognized.
local valid(p, str) = std.length(p(str, 0)) == std.length(str);

std.join(
  '\n',
  [
    "a('ab', 0)                   => %s" % [a('ab', 0)],
    "b('ab', 1)                   => %s" % [b('ab', 1)],
    "alternate(a, b)('ab', 0)     => %s" % [alternate(a, b)('ab', 0)],
    "alternate(a, b)('ab', 1)     => %s" % [alternate(a, b)('ab', 1)],
    "concat(a, b)('ab', 0)        => %s" % [concat(a, b)('ab', 0)],
    "star(a)('ab', 0)             => %s" % [star(a)('ab', 0)],
    "star(a)('aaaaa', 0)          => %s" % [star(a)('aaaaa', 0)],
    "valid(concat(a, b), 'ab')    => %s" % [valid(concat(a, b), 'ab')],
    "valid(star(a), 'aaaaa')      => %s" % [valid(star(a), 'aaaaa')],
    "seq([a, b, b])('abb', 0)     => %s" % [seq([a, b, b])('abb', 0)],
    "valid(seq([a, b, b]), 'abb') => %s" % [valid(seq([a, b, b]), 'abb')],
    "valid(digit, '0')            => %s" % [valid(digit, '0')],
  ] +
  [
    "valid(seq([digit, char('-'), digit, char(' '), lower, char(':'), char(' '), star(lower)]), '1-3 a: abcde'), => %s"
    % [valid(seq([digit, char('-'), digit, char(' '), lower, char(':'), char(' '), star(lower)]), '1-3 a: abcde')],
  ]
)
