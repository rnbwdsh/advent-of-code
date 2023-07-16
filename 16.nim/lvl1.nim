import std/strutils
import std/sets
import aocd

let directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]

proc solve(inp: string, partB: bool): int =
  var facing, x, y = 0
  var seen = toHashSet([(x, y)])
  for line in inp.split(", "):
    let dir = line[0]
    let dist = line.stri[1..^1].parseInt
    facing = (facing + (if dir == 'R': 1 else: -1) + 4) mod 4
    for _ in 0..<dist:
      x += directions[facing][0]
      y += directions[facing][1]
      if partB and seen.containsOrIncl((x, y)):
        return x.abs + y.abs
  return x.abs + y.abs


assert solve("R2, L3", false) == 5
assert solve("R8, R4, R4, R8", true) == 4
let inp = requestLevel(1)
echo solve(inp, false)
echo solve(inp, true)
