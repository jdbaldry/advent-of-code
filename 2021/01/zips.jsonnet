local inp = import 'twothousand.libsonnet';

function(method)
  if method == 'arraycomp' then
    [[x, x] for x in inp]
  else if method == 'desugared' then
    std.join(
      [],
      std.makeArray(std.length(inp), function(i) local x = inp[i]; [x, x])
    )
  else if method == 'map' then
    std.map(function(x) [x, x], inp)
  else if method == 'recursive' then
    local zip(inp) =
      if std.length(inp) == 0 then
        []
      else
        local x = [[inp[0], inp[0]]];
        x + zip(inp[1:]);
    zip(inp)
  else if method == 'rec tailstrict' then
    local zip(inp) =
      if std.length(inp) == 0 then
        []
      else
        local x = [[inp[0], inp[0]]];
        x + zip(inp[1:]);
    zip(inp) tailstrict
  else if method == 'acc' then
    local zip(xs, ys) =
      local aux(xs, ys, acc) =
        if std.length(xs) == 0 || std.length(ys) == 0 then
          acc
        else
          aux(xs[1:], ys[1:], acc + [[xs[0], ys[0]]]);
      aux(xs, ys, [])
    ;
    zip(inp, inp)
  else if method == 'acc tailstrict' then
    local zip(xs, ys) =
      local aux(xs, ys, acc) =
        if std.length(xs) == 0 || std.length(ys) == 0 then
          acc
        else
          aux(xs[1:], ys[1:], acc + [[xs[0], ys[0]]]);
      aux(xs, ys, [])
    ;
    zip(inp, inp) tailstrict
  else error 'unrecognized method %s' % method
