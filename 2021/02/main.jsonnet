local input = importstr 'input.txt';
// local input = importstr 'test.txt';

local pc = import '../../2020/02/parser-combinators.libsonnet';

local commands = std.filter(function(line) line != '', std.split(input, '\n'));

local parse(parser=function(str, state) state, converter=function(str) str) = function(str, prev)
  local state = parser(str, prev);
  local length = std.length(state);
  assert length != 0 : 'unable to parse command, error in state: %s' % [prev];
  {
    value: converter(state[length - 1].parsed),
    index: state[length - 1].index,
  }
;

local state = [pc.init[0] { line: 1 }];
local parseCommand(command, line) =
  local alternates = parse(pc.alternates([
    pc.string('down'),
    pc.string('forward'),
    pc.string('up'),
  ]))(command, [state[0] { line: line }]);
  local direction = alternates.value;
  local index = alternates.index;

  local space = parse(pc.char(' '))(command, [state[0] { index: index, line: line }]);
  local index = space.index;

  local digits = parse(pc.star(pc.digit), std.parseInt)(command, [state[0] { index: index, line: line }]);
  local distance = digits.value;
  local index = digits.index;
  assert index == std.length(command) : 'unparsed command: "%s"' % [command[index:]];
  { direction: direction, distance: distance }
;
local broken = [
  local travelled = std.foldr(
    function(command, acc)
      local line = acc[std.length(acc) - 1][0] + 1;
      local parsed = parseCommand(command, line);

      acc { [parsed.direction]+: parsed.distance },
    commands,
    {}
  );
  (travelled.down - travelled.up) * travelled.forward,
  local position = { aim: 0, horizontal: 0, depth: 0 };
  local final = std.foldl(
    function(acc, command)
      local line = acc[std.length(acc) - 1][0] + 1;
      local parsed = parseCommand(command, line);

      if parsed.direction == 'down' then
        acc { aim+: parsed.distance }
      else if parsed.direction == 'up' then
        acc { aim: acc.aim - parsed.distance }
      else if parsed.direction == 'forward' then
        acc { horizontal: acc.horizontal + parsed.distance, depth: acc.depth + acc.aim * parsed.distance }
      else
        error 'unrecognized direction "%s" on line %d' % [parsed.direction, line],
    commands,
    position,
  );
  final.horizontal * final.depth,
];
[]
