import std/strutils
import std/sequtils
import std/algorithm
import aocd
import neo
import neo_ops

proc solveA(inp: string): int =
  for line in inp.splitLines:
    var a = line.splitWhitespace.map(parseInt)
    a.sort
    echo a
    if a[0] + a[1] > a[2]:
      inc result

proc solveB(inp: string): int =
  var arrArr = inp.splitLines.mapIt(it.splitWhitespace.map(parseInt))
  var mat = matrix(arrArr).t
  for av in mat.reshape(mat.ld, 3).rows:
    var a: seq[int] = av.toSeq
    a.sort
    if a[0] + a[1] > a[2]:
      inc result


assert solveA("5  10  25") == 0
assert solveB("""101 301 501
102 302 502
103 303 503
201 401 601
202 402 602
203 403 603""") == 6
let inp = requestLevel(3)
echo solveA(inp)
echo solveB(inp)
