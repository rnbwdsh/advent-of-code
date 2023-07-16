import std/strutils
import aocd

type Disk* = object
  diskIdx, nrPos, pos0: int

proc checkRun(t0: int, disks: seq[Disk]): bool =
  for d in disks:
    var opening = ((d.diskIdx + t0 + d.pos0)) mod d.nrPos
    if opening != 0:
      return false
  return true

proc solve(inp: string, partB: bool): int =
  var disks: seq[Disk]
  for line in inp.splitLines:
    var parts = line.replace("#", "").replace(".", "").split
    disks.add(Disk(diskIdx: parts[1].parseInt, nrPos: parts[3].parseInt,
        pos0: parts[^1].parseInt))
  if partB:
    disks.add(Disk(diskIdx: len(disks)+1, nrPos: 11, pos0: 0))
  for i in 0..100_000_000:
    if checkRun(i, disks):
      return i


assert solve("""Disc #1 has 5 positions; at time=0, it is at position 4.
Disc #2 has 2 positions; at time=0, it is at position 1.""", false) == 5
let inp = requestLevel(15)
echo solve(inp, false)
echo solve(inp, true)
