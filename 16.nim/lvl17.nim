import std/md5
import aocd

type Pathxy = (string, int, int)

const openLetters = "bcdef"

proc solve(inp: string, partB: bool): string =
  var paths = @[("", 0, 0)]
  for pathLen in 0..1000:
    var nextPaths: seq[Pathxy]
    for (path, x, y) in paths:
      let hash4 = (inp & path).getMD5
      if x == 3 and y == 3:
        if partB:
          result = $pathLen
          continue
        else:
          return path
      if hash4[0] in openLetters and y > 0:
        nextPaths.add((path & "U", x, y-1))
      if hash4[1] in openLetters and y < 3:
        nextPaths.add((path & "D", x, y+1))
      if hash4[2] in openLetters and x > 0:
        nextPaths.add((path & "L", x-1, y))
      if hash4[3] in openLetters and x < 3:
        nextPaths.add((path & "R", x+1, y))
    paths = nextPaths
    if len(paths) == 0:
      return


assert solve("vwbaicqe", false) == "DRDRULRDRD"
assert solve("kglvqrro", true) == "492"
let inp = requestLevel(17)
echo solve(inp, false)
echo solve(inp, true)
