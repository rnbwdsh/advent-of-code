import std/strutils
import aocd
import neo
import neo_ops

let directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]
let lookup_array = matrix(@[
    @[0, 0, 1, 0, 0],
    @[0, 2, 3, 4, 0],
    @[5, 6, 7, 8, 9],
    @[0, 10, 11, 12, 0],
    @[0, 0, 13, 0, 0],
  ]).t

proc clamp(x: int, upperBound: int): int = max(0, min(x, upperBound))

proc solveA(inp: string): string =
  var x, y = 1
  var code = ""
  for line in inp.splitLines:
    for c in line:
      var dirIdx = "DRUL".find(c)
      x = clamp(x + directions[dirIdx][0], 2)
      y = clamp(y + directions[dirIdx][1], 2)
    code &= $(1 + x + y * 3)
  return code

proc solveB(inp: string): string =
  var x = 0
  var y = 2
  var code = ""
  for line in inp.splitLines:
    for c in line:
      var dirIdx = "DRUL".find(c)
      let xn = clamp(x + directions[dirIdx][0], 4)
      let yn = clamp(y + directions[dirIdx][1], 4)
      if lookup_array[xn, yn] != 0:
        (x, y) = (xn, yn)
    let lookedUp = lookup_array[x, y]
    code &= lookedUp.toHex[^1]
  return code


assert solveA("ULL\nRRDDD\nLURDL\nUUUUD") == "1985"
assert solveB("ULL\nRRDDD\nLURDL\nUUUUD") == "5DB3"
let inp = requestLevel(2)
echo solveA(inp)
echo solveB(inp)
