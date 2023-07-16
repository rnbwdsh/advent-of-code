import std/strutils
import aocd

proc solve(inp: string, rows: int): int =
  result = inp.count(".")
  var currRow = inp
  var nextRow = inp
  for i in 0..<rows-1:
    var currRowPadded = "." & currRow & "."
    for i in 0..<len(currRow):
      let nextIsTrap: bool = currRowPadded[i..i+2] in ["^^.", ".^^", "^..", "..^"]
      nextRow[i] = (if nextIsTrap: '^' else: '.')
    currRow = nextRow
    result += currRow.count(".")


assert solve("""..^^.""", 3) == 6
assert solve(".^^.^.^^^^", 10) == 38
let inp = requestLevel(18)
echo solve(inp, 40)
echo solve(inp, 400000)
