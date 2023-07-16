import std/strutils
import std/algorithm
import aocd

proc mergeRanges(records: seq[(uint, uint)]): seq[(uint, uint)] =
  var r = records
  r.sort
  var i = 0
  while i < len(r):
    var (startCurr, finCurr) = r[i]
    for j in i+1..<len(r):
      var (startNext, finNext) = r[j]
      if startNext <= finCurr:
        finCurr = max(finCurr, finNext)
        r[j] = (0.uint, 0.uint)
    if startCurr > 0 or finCurr > 0:
      result.add((startCurr, finCurr))
    inc i

proc solve(inp: string, partB: bool): uint =
  var records: seq[(uint, uint)]
  for line in inp.splitLines:
    let ab = line.split("-")
    records.add((ab[0].parseUInt, ab[1].parseUInt+1)) # make ends exclusives
  let records2 = records.mergeRanges
  if partB:
    result = 4294967296.uint
    for (s, f) in records2:
      result -= f - s
  else:
    return records2[0][1]


assert solve("""5-8
0-2
4-7""", false) == 3
let inp = requestLevel(20)
echo solve(inp, false)
echo solve(inp, true)
