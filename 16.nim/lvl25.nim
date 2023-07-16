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
  var lines = inp.splitLines.mapIt(it.split(" "))
  var pc = 0
  var lastOut = 1
  var outCnt = 0
  var reg: Table[char, int] = {'a': initA, 'c': 0}.toTable
  const toggle: Table[string, string] = {"jnz": "cpy", "cpy": "jnz",
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
      of "out":
        let outv = resolve(instr[1], reg)
        if outv != lastOut:
          lastOut = outv
          inc outCnt
        else:
          return 0
    if outCnt == 100:
      return initA
    inc pc
  return 0

proc solveA(inp: string): int =
  for i in 1..1_000_00_000:
    if solve(inp, i) != 0:
      return i

let inp = requestLevel(25)
echo solveA(inp)
