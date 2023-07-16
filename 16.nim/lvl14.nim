import std/strutils
import std/tables
import std/md5
import std/strformat
import aocd


#[proc plotBoolMatrix(matrix: Matrix[bool]) =
  for row in matrix.rows:
    echo row.mapIt(if it: "#" else: ".").join("")
  echo ""]#
var cache: Table[string, string]


proc getMD5Cached(s: string, partB: bool): string =
  if not partB:
    result = s.getMD5
  else:
    if cache.contains(s):
      return cache[s]
    result = s
    for i in 0..2016:
      result = result.getMD5
    cache[s] = result

proc checkInner(inp: string, i: int, t: char, partB: bool): bool =
  for j in i+1..i+1000:
    let hash2 = (inp & $j).getMD5Cached(partB)
    let needle = fmt"{t}{t}{t}{t}{t}"
    if needle in hash2:
      return true

proc solve(inp: string, partB: bool): int =
  var startT: Table[int, char]
  var cntKeys = 0
  for i in 0..1_000_000:
    # echo toHash
    let hash = (inp & $i).getMD5Cached(partB)
    for p in 0 ..< len(hash) - 2:
      if hash[p] == hash[p+1] and hash[p+1] == hash[p+2]:
        var ci = checkInner(inp, i, hash[p], partB)
        if ci:
          inc cntKeys
          if cntKeys == 64:
            return i
        break


echo solve("abc", false)
echo solve("abc", true)
let inp = requestLevel(14)
echo solve(inp, false)
echo solve(inp, true)
