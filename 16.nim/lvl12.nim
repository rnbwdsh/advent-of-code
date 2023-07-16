import std/strutils
import std/tables
import aocd

proc solve(inp: string, partB: bool): int =
  let prog = inp.splitLines
  var pc = 0
  var variables: Table[char, int] = {'c': if partB: 1 else: 0, '1': 1}.toTable
  while pc < prog.len:
    let chunks = prog[pc].split(" ")
    case chunks[0]:
      of "cpy":
        try:
          variables[chunks[2][0]] = chunks[1].parseInt
        except ValueError:
          variables[chunks[2][0]] = variables[chunks[1][0]]
      of "inc":
        inc variables[chunks[1][0]]
      of "dec":
        dec variables[chunks[1][0]]
      of "jnz":
        if variables[chunks[1][0]] != 0:
          pc += chunks[2].parseInt
          continue
    inc pc
  return variables['a']


assert solve("""cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a""", true) == 42
let inp = requestLevel(12)
echo solve(inp, false)
echo solve(inp, true)
