local pc = import 'parser-combinators.libsonnet';

local input = importstr 'input.txt';
//local input = std.join('\n', ['1-3 a: abcde', '1-3 b: cdefg', '2-9 c: ccccccccc']);

local lines = std.filter(function(line) std.length(line) > 0, std.split(input, '\n'));

// parse parses a string with 'parser' and converts it to the desired type with 'converter'.
local parse(parser=function(str, state) state, converter=function(str) str) = function(str, prev)
  local state = parser(str, prev);
  local length = std.length(state);
  assert length != 0 : 'unable to parse str, error in state: %s' % [state];

  {
    value: converter(state[length - 1].parsed),
    index: state[length - 1].index,
  }
;

local parseLine(line) =
  local digits = parse(pc.star(pc.digit), std.parseInt)(line, pc.init);
  local min = digits.value;
  local index = digits.index;

  local any = parse(pc.any)(line, pc.init { index: index });
  local index = any.index;

  local digits = parse(pc.star(pc.digit), std.parseInt)(line, pc.init { index: index });
  local max = digits.value;
  local index = digits.index;

  local any = parse(pc.any)(line, pc.init { index: index });
  local index = any.index;

  local lower = parse(pc.lower)(line, pc.init { index: index });
  local char = lower.value;
  local index = lower.index;

  local any2 = parse(pc.n(pc.any, 2))(line, pc.init { index: index });
  local index = any2.index;

  local lowers = parse(pc.star(pc.lower))(line, pc.init { index: index });
  local string = lowers.value;

  {
    min: min,
    max: max,
    char: char,
    string: string,
  }
;

local parsed = std.map(parseLine, lines);
[
  std.foldr(
    function(line, acc)
      local count = std.count(std.stringChars(line.string), line.char);
      if line.min <= count && count <= line.max then acc + 1 else acc,
    parsed,
    0,
  ),
  std.foldr(
    function(line, acc)
      if (
           line.string[line.min - 1] == line.char
           ||
           line.string[line.max - 1] == line.char
         )
         &&
         !(
           line.string[line.min - 1] == line.char
           &&
           line.string[line.max - 1] == line.char
         ) then
        acc + 1
      else acc,
    parsed,
    0,
  ),
]
