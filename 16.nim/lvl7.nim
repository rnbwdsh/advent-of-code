import std/strutils
import std/re
import std/sequtils
import std/math
import aocd

let inBrackets = re(r"\[(.*?)\]")

proc abaMatch(inp: string): bool =
  for i in 0 ..< inp.len - 3:
    if inp[i] == inp[i+3] and inp[i+1] == inp[i+2] and inp[i] != inp[i+1]:
      return true
  return false

proc abaFindAll(inp: string): seq[string] =
  for i in 0 ..< inp.len - 2:
    if inp[i] == inp[i+2] and inp[i] != inp[i+1]
      if inp[i] != ' ' and inp[i+1] != ' ':
        result.add(inp[i..i+2])

proc lineMatchA(line: string): int =
  if abaMatch(line):
    for match in findAll(line, inBrackets):
      if len(match) > 4 and abaMatch(match):
        return 0
    return 1

proc lineMatchB(line: string): int =
  var inside = findAll(line, inBrackets)
  var innerAba: seq[string] = @[]
  var lc = line
  for i in inside:
    lc = lc.replace(i, " ")
    for a in abaFindAll(i):
      innerAba.add(a[1] & a[0] & a[1])
  for a in abaFindAll(lc):
    if innerAba.contains(a):
      return 1

proc solve(inp: string, b: bool): int =
  inp.splitLines.map(if b: lineMatchB else: lineMatchA).sum


assert solve("""abba[mnop]qrst
abcd[bddb]xyyx
aaaa[qwer]tyuiioxxoj[asdfgh]zxcvbn""", false) == 2
assert solve("""aba[bab]xyz
xyx[xyx]xyx
aaa[kek]eke
zazbz[bzb]cdb""", true) == 3
let inp = requestLevel(7)
echo solve(inp, false)
echo solve(inp, true)
