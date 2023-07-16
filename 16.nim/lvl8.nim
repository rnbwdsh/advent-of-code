import std/strutils
import std/sequtils

import neo
import aocd

const WIDTH = 50
const HEIGHT = 6

proc printField(field: Matrix[int]): int =
  for row in field.rows:
    let niceRow = row.mapIt(if it == 1: '#' else: ' ').join("")
    echo niceRow
    result += niceRow.count('#')
  echo ""

proc unpack2(s2: seq[int]): (int, int) = (s2[0], s2[1])

proc solve(inp: string): int =
  var field = zeros(HEIGHT, WIDTH, int)
  for line in inp.splitLines:
    if line.startsWith("rect"):
      let (w, h) = unpack2(line[5..^1].split("x").mapIt(it.strip.parseInt))
      for y in 0..h-1:
        for x in 0..w-1:
          field[y, x] = 1

    elif line.startsWith("rotate row"):
      let (y, by) = unpack2(line[13..^1].split(" by ").mapIt(it.strip.parseInt))
      for _ in 0..by-1:
        var last = field[y, WIDTH-1]
        for x in 0..WIDTH-1:
          let tmp = field[y, x]
          field[y, x] = last
          last = tmp

    elif line.startsWith("rotate column"):
      let (x, by) = unpack2(line[16..^1].split(" by ").mapIt(it.strip.parseInt))
      for _ in 0..by-1:
        var last = field[HEIGHT-1, x]
        for y in 0..HEIGHT-1:
          let tmp = field[y, x]
          field[y, x] = last
          last = tmp

    result = printField(field)


let inp = requestLevel(8)
assert solve("""rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1""") == 6
echo $(solve(inp))
