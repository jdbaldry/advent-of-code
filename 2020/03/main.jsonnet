local input = importstr 'input.txt';
// local input = std.join(
//   '\n',
//   [
//     '..##.......',
//     '#...#...#..',
//     '.#....#..#.',
//     '..#.#...#.#',
//     '.#...##..#.',
//     '..#.##.....',
//     '.#.#.#....#',
//     '.#........#',
//     '#.##...#...',
//     '#...##....#',
//     '.#..#...#.#',
//   ]
// )
// ;

local map = std.filter(function(line) std.length(line) > 0, std.split(input, '\n'));

// traverse returns the number of trees encountered when traversing the map
// travelling 'right' steps across, and 'down' steps down.
local traverse(map, right, down) =
  local length = std.length(map);
  local aux(i, j, acc) =
    if i >= length then
      acc
    else
      aux(
        i + down,
        (j + right) % std.length(map[i]),
        if map[i][j] == '#' then acc + 1 else acc
      ) tailstrict
  ;
  aux(0, 0, 0)
;

[
  traverse(map, 3, 1),
  std.foldr(
    function(slope, acc) acc * traverse(map, slope.right, slope.down),
    [{ right: 1, down: 1 }, { right: 3, down: 1 }, { right: 5, down: 1 }, { right: 7, down: 1 }, { right: 1, down: 2 }],
    1,
  ),
]
