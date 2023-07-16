import std/strutils
import neo
import aocd
const MAX_INT = 0x7fffffff

proc solve(inp: string, b: bool): string =
  var counters = zeros(8, 256, int)
  for line in inp.splitLines:
    for i, c in line:
      counters[i, c.ord] += 1
  for i in 0..7:
    if b:
      var minPos = MAX_INT
      var minVal = MAX_INT
      var row = counters.row(i)
      for j, x in row:
        if x > 0 and x < minVal:
          minVal = x
          minPos = j
      if minPos < MAX_INT:
        result = result & $(minPos).chr
    else:
      var mi = counters.row(i).maxIndex;
      if mi.val > 0:
        result = result & $(counters.row(i).maxIndex).i.chr

let data = """eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar"""
assert solve(data, false) == "easter"
assert solve(data, true) == "advent"
let inp = requestLevel(6)
echo solve(inp, false)
echo solve(inp, true)
