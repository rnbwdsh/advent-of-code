import std/strutils
import std/sequtils
import aocd

proc solveA(inp: string): int =
  let nrElfs = inp.parseInt
  var elfs = newSeqWith(nrElfs, true)
  while elfs.count(true) > 1:
    for i in 0..<nrElfs:
      if elfs[i]:
        result = i + 1
        # echo i
        for j in 1..<nrElfs:
          let tpos = (i+j) mod nrElfs
          if elfs[tpos]:
            elfs[tpos] = false
            break

proc solveB(inp: string): int =
  let nrElfs = inp.parseInt
  var i = 1
  while i * 3 < nrElfs:
    i *= 3
  return nrElfs - i


assert solveA("5") == 3
let inp = requestLevel(19)
echo solveA(inp)
echo solveB(inp)
