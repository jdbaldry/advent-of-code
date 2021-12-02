local xs = import 'twothousand.libsonnet';

function(method)
  if method == 'makeArray' then
    // $ time jsonnet --tla-str method=makeArray arrays.jsonnet > /dev/null
    // real 0m6.218s
    // user 0m9.221s
    // sys	0m0.263s
    std.map(
      function(_) std.makeArray(std.length(xs), function(i) if i == std.length(xs) - 1 then 0 else xs[i + 1]),
      std.range(0, 1000)
    )
  else if method == 'slice' then
    // $ time jsonnet --tla-str method=slice arrays.jsonnet > /dev/null
    // real 0m19.903s
    // user 0m52.744s
    // sys	0m0.472s
    std.map(
      function(_) xs[1:] + [0],
      std.range(0, 1000)
    )
  else error 'unrecognized method %s' % [method]
