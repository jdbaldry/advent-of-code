// notEmpty :: String -> Bool
local notEmpty(str) = std.length(str) != 0;
// lines :: String -> [String]
local lines(str) = std.filter(notEmpty, std.split(str, '\n'));
// words :: String -> [String]
local words(str) = std.filter(notEmpty, std.split(str, ' '));
// cut :: String -> String -> String
local cut(str) = function(cutset)
  std.foldr(
    function(c, str) std.join('', std.filter(function(_c) _c != c, std.stringChars(str))),
    std.stringChars(cutset),
    str
  )
;
// sortStr :: String -> String
local sortStr(str) = std.join('', std.sort(std.stringChars(str)));
// sum :: [Number] -> Number
local sum(ns) = std.foldr(function(n, acc) n + acc, ns, 0) tailstrict;
// last :: [a] -> a
local last(arr) = arr[std.length(arr) - 1];

local stack = {
  new(size):: {
    size: size,
    stack: [],
  },
  push(stack, e)::
    assert std.length(stack.stack) < stack.size : 'stack overflow';
    stack { stack+: [e] },
  pop(stack)::
    local l = std.length(stack.stack);
    assert l >= 1 : 'stack underflow';
    local e = stack.stack[l - 1];
    [stack { stack: stack.stack[:l - 1] }, e],
  peek(stack)::
    local l = std.length(stack.stack);
    assert l >= 1 : 'stack underflow';
    stack.stack[l - 1],
  size(stack):: std.length(stack.stack),
};
{
  // String functions.
  cut(str):: cut(str),
  lines(str):: lines(str),
  notEmpty(str):: notEmpty(str),
  sortStr(str):: sortStr(str),
  words(str):: words(str),

  // Number functions.
  sum(ns):: sum(ns),

  // Array functions.
  last(arr):: last(arr),

  stack:: stack,
}
