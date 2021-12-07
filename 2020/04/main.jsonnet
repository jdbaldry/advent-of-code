local pc = import '../02/parser-combinators.libsonnet';

local input = importstr 'input.txt';
// local input = importstr 'test.txt';

local passport = {
  // Birth Year
  byr: null,
  // Issue Year
  iyr: null,
  // Expiration Year
  eyr: null,
  // Height
  hgt: null,
  // Hair Color
  hcl: null,
  // Eye Color
  ecl: null,
  // Passport ID
  pid: null,
  // Country ID
  cid: null,
};

local byrP = pc.seq([pc.string('byr'), pc.char(':'), pc.n(pc.digit, 4)]);
local iyrP = pc.seq([pc.string('iyr'), pc.char(':'), pc.n(pc.digit, 4)]);
local eyrP = pc.seq([pc.string('eyr'), pc.char(':'), pc.n(pc.digit, 4)]);
local hgtP = pc.seq([pc.string('hgt'), pc.char(':'), pc.start(pc.digit), pc.alternate(pc.string('in'), pc.string('cm'))]);
local hclP = pc.seq([pc.string('hcl'), pc.char('#'), pc.set(std.stringChars('0123456789abcdef'))]);
local eclP = pc.seq([
  pc.string('ecl'),
  pc.alternates(std.map(pc.string, ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'othb'])),
]);
local pidP = pc.seq([pc.string('pid'), pc.n(pc.digit, 9)]);
local cidP = pc.seq([pc.string('cid'), pc.start(pc.digit)]);

local passportsP =
  pc.star(pc.concat(
    pc.alternates([byrP, iyrP, eyrP, hgtP, hclP, eclP, pidP, cidP]),
    pc.star(pc.set([
      pc.char(' '),
      pc.char('\n'),
    ]))
  ))
;

// local passports =
//   std.foldl(
//     function(acc, line) if line == '' then
//       { current: passport, all: [acc.current] + acc.all }
//     else
//       {
//         current: std.foldr(
//           function(field, acc) acc { [field[0]]: field[1] },
//           std.map(function(field) std.split(field, ':'), std.split(line, ' ')),
//           acc.current,
//         ),
//         all: acc.all,
//       },
//     lines,
//     { current: passport, all: [] },
//   ).all
// ;


// local validate(passport) =
//   std.foldr(function(field, acc) (field == 'cid' || passport[field] != null) && acc, std.objectFields(passport), true);

// std.foldr(
//   function(passport, acc) if validate(passport) then acc + 1 else acc,
//   passports,
//   0,
// )

passportsP(input, pc.init)
