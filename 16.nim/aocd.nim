import std/httpclient
import std/strutils
import std/strformat
import std/os
import std/json

let year = 2016
let PREFIX = getHomeDir() / ".config/aocd"
let token = read_file(PREFIX / "token").strip
let jsonNode = parseJson(read_file(PREFIX / "token2id.json"))
var id: string
for a, b in jsonNode.pairs:
  id = b.str
let dataDir = PREFIX / id


proc aocdClient(): HttpClient =
  return newHttpClient(headers = newHttpHeaders({
        "Cookie": fmt"session={token}",
        "Content-Type": "application/x-www-form-urlencoded"}))

proc requestLevel(day: int): string =
  # 2015_01_input.txt
  let cacheFile = fmt"{dataDir}/{year}_{day:02}_input.txt"
  if fileExists(cacheFile):
    return readFile(cacheFile).strip
  let url = fmt"https://adventofcode.com/2016/day/{day}/input"
  result = aocdClient().request(url, httpMethod = HttpMethod.HttpGet).body.strip
  write_file(cacheFile, result)

# export symbols
export requestLevel
