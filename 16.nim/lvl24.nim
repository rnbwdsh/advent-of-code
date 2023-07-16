import std/strutils
import std/sequtils
import std/tables
import std/enumerate
import std/random
import neo
import aocd

type XYDist = Table[(int, int), int]
type StartXY = Table[int, (int, int)]

proc chr2int(c: char): int = c.ord - '0'.ord
proc hash2(x, y: int): int16 = (x + y * 256).int16

proc makeMatrix(inp: string): (Matrix[char], StartXY) =
  let lines = inp.splitlines
  var m = zeros(lines[0].len, lines.len, char)
  var starts: Table[int, (int, int)]
  for y, line in enumerate(lines):
    for x, c in enumerate(line):
      m[x, y] = c
      if c != '.' and c != '#':
        starts[chr2int(c)] = (x, y)
  return (m, starts)

proc calculateDistMatrix(m: Matrix[char], starts: StartXY): XYDist =
  for s, (x, y) in starts.pairs:
    var seen: set[int16]
    var distFromStart = 1
    var states = @[(x, y)]
    seen.incl(hash2(x, y))
    while states.len > 0:
      var nextStates: seq[(int, int)]
      for (x, y) in states:
        for (xn, yn) in [(x+1, y), (x-1, y), (x, y+1), (x, y-1)]:
          let ha = hash2(xn, yn)
          if ha in seen:
            continue
          if m[xn, yn] != '#':
            nextStates.add((xn, yn))
            seen.incl(ha)
            if m[xn, yn] != '.':
              let t = m[xn, yn].chr2int
              result[(s, t)] = distFromStart
      inc distFromStart
      states = nextStates

proc calculateShortestPath(dist: XYDist, starts: int, partB: bool): int =
  result = 10_000
  for _ in 0..100_000:
    var path = (1..starts).toSeq
    path.shuffle
    var pathDist = dist[(0, path[0])]
    for i in 0..<starts-1:
      pathDist = pathDist + dist[(path[i], path[i+1])]
    if partB:
      pathDist += dist[(path[^1], 0)]
    result = min(result, pathDist)

proc solve(inp: string, partB: bool): int =
  var (m, starts) = makeMatrix(inp)
  var distMatrix = calculateDistMatrix(m, starts)
  return calculateShortestPath(distMatrix, len(starts)-1, partB)


assert solve("""###########
#0.1.....2#
#.#######.#
#4.......3#
###########""", false) == 14
let inp = requestLevel(24)
echo solve(inp, false)
echo solve(inp, true)
