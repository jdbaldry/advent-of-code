local input = importstr 'input.txt';
// local input = std.join('\n', ['1721', '979', '366', '299', '675', '1456']);

local lines = std.filter(function(line) std.length(line) > 0, std.split(input, '\n'));
local expenses = std.map(std.parseInt, lines);

local sum(numbers=[]) =
  std.foldr(
    function(number, acc) acc + number,
    numbers,
    0
  )
;

local sumsTo(n) = function(numbers) sum(numbers) == n;

local product(numbers=[]) =
  std.foldr(
    function(number, acc) acc * number,
    numbers,
    1
  )
;

local cartesianProduct(arr_a=[], arr_b=[]) =
  if std.length(arr_a) == 0 || std.length(arr_b) == 0 then
    []
  else
    std.foldr(
      function(a, acc) acc + std.map(
        function(b) [a, b],
        arr_b
      ),
      arr_a,
      []
    )
;

local cartesianProductN(arrs) =
  if std.length(arrs) <= 1 then
    arrs
  else if std.length(arrs) == 2 then
    cartesianProduct(arrs[0], arrs[1])
  else
    std.flattenArrays(
      std.map(
        function(e)
          std.map(
            function(rest) [e] + std.flattenArrays([rest]),
            cartesianProductN(arrs[1:]),
          ),
        arrs[0],
      )
    )
;

local while(cond=function(state) true, do=function(state) state, state) =
  if !cond(state) then
    state
  else while(cond, do, do(state)) tailstrict;

local last = std.length(expenses) - 1;

null

// Failed attempts:
//
// [
//   product(
//     std.filter(
//       sumsTo(2020),
//       cartesianProductN([expenses, expenses])
//     )[0]
//   ),
//   // Way too expensive in RAM.
//   product(
//     std.filter(
//       sumsTo(2020),
//       cartesianProductN([expenses, expenses, expenses])
//     )[0]
//   ),
// ]
//
// [
//   while(
//     function(s) s.product == null && (s.i != 0 || s.j != 0),
//     function(s)
//       {
//         i: if s.j != 0 then s.i else s.i - 1,
//         j: if s.j == 0 then last else s.j - 1,
//         product: if expenses[s.i] + expenses[s.j] == 2020 then expenses[s.i] * expenses[s.j],
//       },
//     {
//       i: last,
//       j: last,
//       k: last,
//       product: null,
//     }
//   ).product,
//   // This exhausts the Go stack size limit.
//   //
//   // runtime: goroutine stack exceeds 1000000000-byte limit
//   // runtime: sp=0xc01ea80990 stack=[0xc01ea80000, 0xc03ea80000]
//   // fatal error: stack overflow

//   // runtime stack:
//   // runtime.throw(0x5a1370, 0xe)
//   // 	runtime/panic.go:1117 +0x72
//   // runtime.newstack()
//   // 	runtime/stack.go:1069 +0x7ed
//   // runtime.morestack()
//   // 	runtime/asm_amd64.s:458 +0x8f

//   // goroutine 1 [running]:
//   // github.com/google/go-jsonnet.(*interpreter).evaluate(0xc000076360, 0x5ef390, 0xc0001dad00, 0x0, 0x0, 0x0, 0x0, 0x0)
//   while(
//     function(s) s.product == null && (s.i != 0 || s.j != 0 || s.k != 0),
//     function(s)
//       {
//         i: if s.j != 0 || s.k != 0 then s.i else s.i - 1,
//         j: if s.k != 0 then s.j else if s.j == 0 then last else s.j - 1,
//         k: if s.k == 0 then last else s.k - 1,
//         product: if expenses[s.i] + expenses[s.j] + expenses[s.k] == 2020 then expenses[s.i] * expenses[s.j] * expenses[s.k],
//       },
//     {
//       i: last,
//       j: last,
//       k: last,
//       product: null,
//     }
//   ).product,
// ]



