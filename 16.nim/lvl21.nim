import std/strutils
import std/sequtils
import std/algorithm
import std/random
import aocd

proc solveA(inp: string, instr: string): string =
  var i = inp.toSeq
  for line in instr.splitLines:
    let c = line.split(" ")
    if line.startsWith("swap position"):
      let x = c[2].parseInt
      let y = c[^1].parseInt
      (i[x], i[y]) = (i[y], i[x])
    if line.startsWith("swap letter"):
      let x = i.find(c[2][0])
      let y = i.find(c[^1][0])
      (i[x], i[y]) = (i[y], i[x])
    if line.startsWith("rotate left"):
      i.rotateLeft(c[2].parseInt)
    if line.startsWith("rotate right"):
      i.rotateLeft(-c[2].parseInt)
    if line.startsWith("rotate based on position of letter"):
      var idx = 1 + i.find(c[^1][0])
      if idx >= 5:
        inc idx
      i.rotateleft(-idx)
    if line.startsWith("reverse positions"):
      let x = c[2].parseInt
      let y = c[^1].parseInt
      i[x..y] = i[x..y].reversed
    if line.startsWith("move position"):
      let x = c[2].parseInt
      let y = c[^1].parseInt
      let letter = i[x]
      i.delete(x)
      i.insert(letter, y)
  return i.join("")

proc solveB(instr: string): string =
  while true:
    var trial = "abcdefgh"
    random.shuffle(trial)
    if solveA(trial, instr) == "fbgdceah":
      return trial


assert solveA("abcde", """swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d""") == "decab"
let inp = requestLevel(21)
echo solveA("abcdefgh", inp)
echo solveB(inp)
