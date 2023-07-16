import std/md5
import aocd

proc solveA(inp: string): string =
  for i in 0..1_000_000_000_000:
    var curr = inp & $i
    var hash = curr.getMD5
    if hash[0..4] == "00000":
      result &= hash[5]
      if len(result) == 8:
        break

proc solveB(inp: string): string =
  result = "________"
  var revealed = 0
  for i in 0..1_000_000_000_000:
    var curr = inp & $i
    var hash = curr.getMD5
    if hash[0..4] == "00000":
      if hash[5] in '0'..'7':
        var idx: int = hash[5].ord - '0'.ord
        if result[idx] == '_':
          result[idx] = hash[6]
          inc revealed
          if revealed == 8:
            break


assert solveA("abc") == "18f47a30"
assert solveB("abc") == "05ace8e3"
let inp = requestLevel(5)
echo solveA(inp)
echo solveB(inp)
