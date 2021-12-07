local state = [{
  captured: [],
  index: 0,
  parsed: '',
}];

// empty parses the empty string ''.
local empty(_, prev) = prev;

// never never parses anything.
local never(_, _) = [];

// satisfies returns a parser that recognizes any character which
// satisfies the predicate 'pred'.
local satisfies(pred=function(c) true) = function(str, prev)
  local l = std.length(str);
  local _l = std.length(prev);
  local i = prev[_l - 1].index;
  local p = prev[_l - 1].parsed;
  local c = str[i];
  if l == 0 || i > l - 1 || !pred(c) then
    []
  else [prev[_l - 1] { index: i + 1, parsed: p + c }]
;


// concat concatenates two parsers, feeding the results from the first into the
// input of the second.
local concat(p1, p2) = function(str, prev)
  local states = p1(str, prev);
  local l = std.length(states);
  if l <= 0 then
    []
  else
    states + p2(str, [states[l - 1]])
;

// seq concatenates all parsers in 'ps'.
local seq(ps=[]) = std.foldr(function(p, acc) concat(p, acc), ps, never);

// star is the Kleene star and applies the same parser zero or more times.
local star(p) = function(str, prev)
  local states = p(str, prev);
  local l = std.length(states);
  if l == 0 then
    states
  else
    states + star(p)(str, [states[l - 1]])
;

// plus applies the parser one or more times.
local plus(p) = concat(p, star(p));

// alternate combines the results of two parsers.
local alternate(p1, p2) = function(str, prev) p1(str, prev) + p2(str, prev);

// alternates combines the result of all parsers in array 'ps'.
local alternates(ps) = std.foldr(function(p, acc) alternate(acc, p), ps, empty);

// n applies the same parser 'k' times.
local n(p, k) = seq([p for i in std.range(1, k)]);

// capture captures the full results of parser p and drops intermediate states.
local capture(p) = function(str, prev)
  local states = p(str, prev);
  local l = std.length(states);
  local _l = std.length(prev);
  if l == 0 then
    []
  else
    local alreadyParsed = std.length(prev[_l - 1].parsed);
    [states[l - 1] {
      captured: prev[_l - 1].captured + [super.parsed[alreadyParsed:]],
    }]
;

// ignore drops the parsed result of 'p' but retains the index state.
local ignore(p) = function(str, prev)
  local states = p(str, prev);
  local l = std.length(prev);
  local _l = std.length(states);
  if _l == 0 then
    prev
  else
    std.makeArray(l, function(i) if i == l - 1 then prev[i] { index: states[_l - 1].index } else prev[i])
;

// any parses any character.
local any = satisfies();

// char parses character 'c'.
local char(c) = satisfies(function(_c) _c == c);

// string parses the string 'str'.
local string(str) = seq(std.map(char, std.stringChars(str)));

// set parses any character in set 's'.
local set(s) = std.foldr(function(c, acc) alternate(acc, char('c')), s, empty);

// digit parses digits 0-9.
local digit = satisfies(function(c) '0' <= c && c <= '9');

// lower parses lowercase alphabetic characters a-z.
local lower = satisfies(function(c) 'a' <= c && c <= 'z');

// ws matches any whitespace character.
local ws = satisfies(function(c) c == ' ' || c == '\n' || c == '\t' || c == '\r');

{
  // Usage examples.
  // $ jsonnet -Se '(import "parser-combinators.libsonnet").examples'
  examples: std.join(
    '\n\n',
    [
      "any('a', state)\n=> %s" % [any('a', state)],
      "alternate(char('a'), string('ab'))('ab', state)\n=> %s" % [alternate(char('a'), string('ab'))('ab', state)],
      "alternates([string('ab'), string('bc'), string('cd')])('cd', state)]\n=> %s" % [alternates([string('ab'), string('bc'), string('cd')])('cd', state)],
      "concat(char('b'), capture(star(any)))('baaaaa', state)\n=> %s" % [concat(char('b'), capture(star(any)))('baaaaa', state)],
      "concat(any, any)('ab')\n=> %s" % [concat(any, any)('ab', state)],
      "concat(lower, ignore(char(',')))('a,', state)\n=> %s" % [concat(lower, ignore(char(',')))('a,', state)],
      "star(any)('aaa', state)\n=> %s" % [star(any)('aaa', state)],
      "seq([char('a'), char('b')])('ab', state)\n=> %s" % [seq([char('a'), char('b')])('ab', state)],
    ]
  ),
  // Initial parser state.
  init:: state,

  // Combinators.
  alternate(p1, p2):: alternate(p1, p2),
  alternates(ps):: alternates(ps),
  capture(p):: capture(p),
  concat(p1, p2):: concat(p1, p2),
  ignore(p):: ignore(p),
  n(p, k):: n(p, k),
  plus(p):: plus(p),
  seq(ps):: seq(ps),
  star(p):: star(p),

  // Parsers.
  any:: any,
  char(c):: char(c),
  digit:: digit,
  lower:: lower,
  set(s):: set(s),
  string(str):: string(str),
  ws:: ws,
}
