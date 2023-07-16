import std/strutils
import std/enumerate
import std/sequtils
import std/sets
import aocd

type Item = int8
type LevelState = set[Item]

type
  GameState* = object
    floor: array[4, LevelState]
    currLevel: int8

proc isValid(layer: LevelState): bool =
  var lc = layer
  var chipWithoutGen = 0
  for item in layer:
    if item > 0:
      if -item in lc:
        lc.excl(item)
        continue
      else:
        inc chipWithoutGen
  return chipWithoutGen == 0 or len(layer) == chipWithoutGen

proc isValid(g: GameState): bool = g.floor.map(isValid).allIt(it)
proc isDone(g: GameState): bool = g.floor[0..2].allIt(len(it) == 0)
proc currFloor(g: GameState): LevelState = g.floor[g.currLevel]
proc currFloorIncl(g: var GameState, i: Item) = g.floor[g.currLevel].incl(i)
proc currFloorExcl(g: var GameState, i: Item) = g.floor[g.currLevel].excl(i)

proc step(gs: HashSet[GameState]): HashSet[GameState] =
  for g in gs:
    if not g.isValid:
      continue
    if g.isDone:
      return initHashSet[GameState]()
    for a in g.currFloor:
      for b in g.currFloor:
        if a < b: # symmetry breaking
          continue
        # remove from current floor
        var ng = g
        ng.currFloorExcl(a)
        ng.currFloorExcl(b)

        # add to target floors: up + down
        if ng.currLevel != 0:
          var gc = ng
          gc.currLevel -= 1
          gc.currFloorIncl(a)
          gc.currFloorIncl(b)
          if gc.isValid:
            result.incl(gc)

        if ng.currLevel != 3:
          var gc = ng
          gc.currLevel += 1
          gc.currFloorIncl(a)
          gc.currFloorIncl(b)
          if gc.isValid:
            result.incl(gc)

proc bfs(g: GameState): int =
  var states = toHashSet([g])
  var seen = states
  for depth in 0..100:
    seen.incl(states)
    states = step(states)
    states.excl(seen)
    if len(states) == 0:
      return depth
    # echo depth, " ", len(states)

proc toItem(sym: string): Item = (sym[0].ord * (if "-compatible" in
    sym: 1 else: -1)).int8

proc solve(inp: string, partB: bool): int =
  var floor: array[4, LevelState]
  for i, line in enumerate(inp.splitLines):
    if "contains nothing relevant" in line:
      continue
    var lineClean = line.replace(",", "")
    .replace("a ", "")
    .replace("and ", "")
    .replace(".", "")
    .replace(" generator", "")
    .replace(" microchip", "")
    .replace("polonium", "q") # our char tick doesn't work otherwise, as there are 2 "p"-ingredients
    for item in lineClean.split(" ")[4..^1]:
      floor[i].incl(toItem(item))
  if partB: # just add some numbers that aren't already taken
    floor[0].incl(1)
    floor[0].incl(-1)
    floor[0].incl(2)
    floor[0].incl(-2)
  return GameState(floor: floor, currLevel: 0).bfs


assert solve("""The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.""", false) == 11
let inp = requestLevel(11).strip
echo solve(inp, false)
echo solve(inp, true)
