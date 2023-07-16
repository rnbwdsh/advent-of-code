import std/strutils
import std/sequtils
import std/algorithm
import std/tables
import aocd

proc cmp(a, b: (char, int)): int =
  if a[1] == b[1]:
    return a[0].int - b[0].int
  return b[1] - a[1]

proc validate(name: string, checksum: string): bool =
  var ct = name.toCountTable
  var ct2 = pairs(ct).toSeq
  ct2.sort(cmp)
  return ct2.mapIt(it[0]).join("").startsWith(checksum)

proc solve(inp: string, partB: bool): int =
  for line in inp.splitLines:
    var
      arr = line.split("-")
      last = arr.pop
      checksum = last[^6..^2]
      id = last[0..^8].parseInt
      name = arr.join("")

    if partB:
      for a in arr:
        var decoded = a.mapIt(((it.int - 'a'.int + id) mod 26 +
            'a'.int).char).join("")
        if decoded == "northpole":
          return id
    elif validate(name, checksum):
      result += id


assert solve("""aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]""", false) == 123+987+404
assert solve("""ghkmaihex-hucxvm-lmhktzx-267[hmxka]""", true) == 267
let inp = requestLevel(4)
echo solve(inp, false)
echo solve(inp, true)
