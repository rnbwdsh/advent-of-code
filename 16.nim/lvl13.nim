import std/strutils
import std/sets
import std/bitops
import std/sugar
import neo
import aocd

const SIZE = 80

proc isWallAt(n, x, y: int): bool = bool(countSetBits(x*x + 3*x + 2*x*y + y +
    y*y + n) mod 2)

#[proc plotBoolMatrix(matrix: Matrix[bool]) =
  for row in matrix.rows:
    echo row.mapIt(if it: "#" else: ".").join("")
  echo ""]#

proc solve(inp: string, xt, yt: int, partB: bool): int =
  let n = inp.parseInt
  var m = makeMatrix(SIZE, SIZE, (i, j) => isWallAt(n, i, j))
  # plotBoolMatrix(m)
  var positions = @[(1, 1)].toHashSet
  var seen = positions
  var npos: HashSet[(int, int)]
  let limit = if partB: 50 else: 100
  for dist in 1..limit:
    npos.clear
    for (x, y) in positions:
      for (xd, yd) in [(0, 1), (1, 0), (-1, 0), (0, -1)]:
        let nx = x + xd
        let ny = y + yd
        let nt = (nx, ny)
        if nt == (xt, yt) and not partB:
          return dist
        if nx >= 0 and nx < SIZE and ny >= 0 and ny < SIZE and not m[nx, ny]:
          if not seen.contains(nt):
            npos.incl(nt)
            seen.incl(nt)
    positions = npos
  return seen.len


assert solve("10", 7, 4, false) == 11
assert solve("10", 7, 4, true) == 151
let inp = requestLevel(13)
echo solve(inp, 31, 39, false)
echo solve(inp, 31, 39, true)
