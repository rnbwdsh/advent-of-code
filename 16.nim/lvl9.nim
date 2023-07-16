import std/strutils
import std/sequtils
import aocd

proc solveA(inp: string): string =
  var pos = 0
  while pos < len(inp):
    if inp[pos] == '(':
      let
        posClose = inp.find(')', pos)
        ab = inp[pos+1..<posClose].split('x').map(parseInt)
        repeatEndPos = min(len(inp), posClose + 1 + ab[0])
        toRepeat = inp[posClose + 1..<repeatEndPos]
      for _ in 0..<ab[1]:
        result &= toRepeat
      pos = posClose + ab[0]

    else:
      result &= inp[pos]
    inc pos

proc solveB(inp: string): int =
  var i = inp
  while '(' in i:
    i = solveA(i)
  return len(i)


assert solveA("""ADVENT""") == "ADVENT"
assert solveA("""A(1x5)BC""") == "ABBBBBC"
assert solveA("""(3x3)XYZ""") == "XYZXYZXYZ"
assert solveA("""A(2x2)BCD(2x2)EFG""") == "ABCBCDEFEFG"
assert solveA("""(6x1)(1x3)A""") == "(1x3)A"
let inp = requestLevel(9)
echo solveA(inp).len
echo solveB(inp)
