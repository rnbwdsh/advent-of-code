import std/strutils
import std/sequtils
import std/nre
import aocd

type Node* = object
  x, y, size, used, avail: int

proc solve(inp: string, partB: bool): int =
  var
    lines = inp.splitlines[2..^1]
    number = re"\d+"
    nodes: seq[Node]
    maxX, maxY: int
    emptyX, emptyY: int
  for line in lines:
    let numbers = line.findAll(number).map(parseInt)
    let n = Node(x: numbers[0], y: numbers[1], size: numbers[2], used: numbers[
        3], avail: numbers[4])
    nodes.add(n)
    if n.used == 0:
      (emptyX, emptyY) = (n.x, n.y)
    maxX = n.x + 1
    maxY = n.y + 1
  if partB:
    # move up from emptyY to 0
    # then from emptyX to 0
    # then up to maxY
    # then 5 steps per rotation to move the data down
    return emptyY + emptyX + (maxY+2) * 6 + 1
  for n1 in nodes:
    for n2 in nodes:
      if n1 != n2 and n1.used > 0 and n1.used <= n2.avail:
        inc result

let inp = requestLevel(22)
echo solve(inp, false)
echo solve(inp, true)
