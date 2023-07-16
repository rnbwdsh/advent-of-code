import std/strutils
import std/sequtils
import std/tables
import aocd

proc resolve(s: string, reg: Table[char, int]): int =
  try:
    return s.parseInt
  except:
    return reg[s[0]]

proc solve(inp: string, initA: int): int =
  var 
    lines = inp.splitLines.mapIt(it.split(" "))
    pc = 0
    reg: Table[char, int] = {'a': initA, 'c': 0}.toTable
    toggle: Table[string, string] = {"jnz": "cpy", "cpy": "jnz",
      "inc": "dec", "dec": "inc", "tgl": "inc"}.toTable
  while pc < lines.len:
    let instr = lines[pc]
    case instr[0]:
      of "cpy":
        reg[instr[2][0]] = resolve(instr[1], reg)
      of "inc":
        inc reg[instr[1][0]]
      of "dec":
        dec reg[instr[1][0]]
      of "jnz":
        if resolve(instr[1], reg) != 0:
          pc += resolve(instr[2], reg) - 1
      of "tgl":
        let target = pc + resolve(instr[1], reg)
        if target < len(lines):
          lines[target][0] = toggle[lines[target][0]]
    inc pc
  return reg['a']


assert solve("""cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a""", 0) == 3
let inp = requestLevel(23)
echo solve(inp, 7)
echo solve(inp, 12)
