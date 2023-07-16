import std/strutils
import std/algorithm
import aocd

proc expand(a: string): string =
  return a & "0" & a.reversed.join("").replace("0", "2").replace("1",
      "0").replace("2", "1")

proc checksum(a: string): string =
  for i in 0..<(len(a) div 2):
    result &= $(a[i*2] == a[i*2+1]).int

proc solve(inp: string, targetLen: int): string =
  result = inp
  while result.len < targetLen:
    result = result.expand
  result = result[0..<targetLen].checksum
  while result.len mod 2 == 0:
    result = result.checksum


assert solve("""10000""", 20) == "01100"
let inp = requestLevel(16)
echo solve(inp, 272)
echo solve(inp, 35651584)
