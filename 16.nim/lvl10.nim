import std/strutils
import std/tables
import aocd


type
  Action* = object
    a, b: int
    abot, bbot: bool

type
  Game = ref object
    bots, outputs: Table[int, int]
    actionTable: Table[int, Action]
    at, bt: int

proc give(self: Game, botId: int, targetType: bool, value: int) =
  if targetType:
    if self.bots.contains(botId):
      var lower = min(self.bots[botId], value)
      var higher = max(self.bots[botId], value)
      # self.bots.del(botId)
      if lower == self.at and higher == self.bt:
        echo "Solution a: ", botId
      var action = self.actionTable[botId]
      self.give(action.a, action.abot, lower)
      self.give(action.b, action.bbot, higher)
    else:
      self.bots[botId] = value
  else:
    self.outputs[botId] = value

proc solve(self: Game, inp: string): int =
  var startBot, startValue: int;
  for line in inp.splitLines:
    var segments = line.split(" ")
    if segments[0] == "value":
      let value = parseInt(segments[1])
      let botId = parseInt(segments[^1])
      if self.bots.contains(botId):
        startValue = value
        startBot = botId
      else:
        self.bots[botId] = value
    if segments[0] == "bot":
      self.actionTable[parseInt(segments[1])] = Action(
        a: parseInt(segments[6]),
        b: parseInt(segments[11]),
        abot: segments[5] == "bot",
        bbot: segments[10] == "bot")
  self.give(startBot, true, startValue)
  return self.outputs[0] * self.outputs[1] * self.outputs[2]


assert Game(at: 2, bt: 5).solve("""value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2
""") == 30
let inp = requestLevel(10)
echo Game(at: 17, bt: 61).solve(inp)
