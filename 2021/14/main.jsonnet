local input = importstr 'input.txt';
local input = importstr 'test.txt';

local pc = import 'parser-combinators.libsonnet';
local util = import 'util.libsonnet',
      lines = util.lines,
      last = util.last
;

local template = std.stringChars(lines(input)[0]);
local ruleP = pc.seq([pc.capture(pc.n(pc.upper, 2)), pc.string(' -> '), pc.capture(pc.upper)]);
local rules = std.foldr(
  function(line, acc)
    local captured = last(ruleP(line, pc.init)).captured;
    { [captured[0]]: captured[1] } + acc,
  lines(input)[2:],
  {}
)
;

local offset(arr) = local l = std.length(arr); std.makeArray(l - 1, function(i) arr[i + 1]);

// step :: [Char] -> [Char]
local step(template) =
  local zip(a, b) =
    local aux(a, b, acc) =
      if std.length(a) == 0 then
        acc
      else
        aux(offset(a), offset(b), acc + [[a[0], b[0]]]) tailstrict
    ;
    if std.length(a) < std.length(b) then
      aux(a, b, [])
    else
      aux(b, a, [])
  ;
  local l = std.length(template);
  std.foldr(
    function(i, acc)
      local a = template[i], b = if i < l - 1 then template[i + 1], index = std.join('', [a, b]);
      [a, if index in rules then rules[index], b] + acc,
    std.range(0, l - 1),
    [],
  ) tailstrict
;

local stepN(n) = function(template)
  if n == 0 then
    template
  else
    stepN(n - 1)(step(template))
;

std.map(
  function(rule) '%s -> %s' % [rule, std.join('', std.prune(stepN(5)(std.stringChars(rule))))],
  [last(std.objectFields(rules))],
)
