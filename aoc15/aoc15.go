package aoc15

import (
	"container/list"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func AtoiPanic(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type pos struct {
	x int
	y int
}

func level1a(data string) int {
	return strings.Count(data, "(") - strings.Count(data, ")")
}

func level1b(data string) int {
	floor := 0
	for i, c := range data {
		if c == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return -1
}

func level2a(data string) int {
	var total = 0
	// iterate lines with for loop
	for _, line := range strings.Split(data, "\n") {
		dim := strings.Split(line, "x")
		var l, w, h = AtoiPanic(dim[0]), AtoiPanic(dim[1]), AtoiPanic(dim[2])
		total += 2*l*w + 2*w*h + 2*h*l + MinOf(l*w, w*h, h*l)
	}
	return total
}

func level2b(data string) int {
	var total = 0
	// iterate lines with for loop
	for _, line := range strings.Split(data, "\n") {
		dim := strings.Split(line, "x")
		var l, w, h = AtoiPanic(dim[0]), AtoiPanic(dim[1]), AtoiPanic(dim[2])
		total += 2*MinOf(l+w, w+h, h+l) + l*w*h
	}
	return total
}

func setPos(c int32, p *pos, visited map[int]bool) {
	switch c {
	case '^':
		p.y++
	case 'v':
		p.y--
	case '<':
		p.x--
	case '>':
		p.x++
	}
	visited[p.x+p.y*10000] = true
}

func level3a(data string) int {
	var p = pos{0, 0}

	visited := make(map[int]bool)
	for _, c := range data {
		setPos(c, &p, visited)
	}
	return len(visited)
}

func level3b(data string) int {
	var pos0 = pos{0, 0}
	var pos1 = pos{0, 0}
	visited := make(map[int]bool)
	for i, c := range data {
		p := &pos0
		if i%2 == 0 {
			p = &pos1
		}
		setPos(c, p, visited)
	}
	return len(visited)
}

func level4(data string, prefix string) int {
	for i := 0; i < 100000000; i++ {
		var s = data + strconv.Itoa(i)
		var h = fmt.Sprintf("%x", md5.Sum([]byte(s)))
		if strings.HasPrefix(h, prefix) {
			return i
		}
	}
	return -1
}

func level4a(data string) int {
	return level4(data, "00000")
}

func level4b(data string) int {
	return level4(data, "000000")
}

func isNice(word string) int {
	// create set of letters
	var doubleLetter = false
	var nrVowels = 0
	for i, c := range word {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			nrVowels++
		}
		if i > 0 {
			var last = int32(word[i-1])
			if last == c {
				doubleLetter = true
			}
			switch fmt.Sprintf("%c%c", last, c) {
			case "ab", "cd", "pq", "xy":
				return 0
			}
		}
	}
	if nrVowels >= 3 && doubleLetter {
		return 1
	}
	return 0
}

func isNice2(word string) int {
	var pair = 0
	var dist1 = 0
	for i := 0; i < len(word)-2; i++ {
		if strings.Contains(word[i+2:], word[i:i+2]) {
			pair = 1
		}
		if word[i] == word[i+2] {
			dist1 = 1
		}
	}
	return pair * dist1
}

func level5a(data string) int {
	total := 0
	for _, word := range strings.Split(data, "\n") {
		total += isNice(word)
	}
	return total
}

func level5b(data string) int {
	total := 0
	for _, word := range strings.Split(data, "\n") {
		total += isNice2(word)
	}
	return total
}

func level6a(data string) int {
	var lights [1000][1000]int
	data = strings.ReplaceAll(data, "turn", "")
	for _, line := range strings.Split(data, "\n") {
		var start string
		var x1, y1, x2, y2 int
		_, _ = fmt.Sscanf(line, "%s %d,%d through %d,%d", &start, &x1, &y1, &x2, &y2)
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				if start == "on" || (lights[x][y] == 0 && start == "toggle") {
					lights[x][y] = 1
				} else if start == "off" || (lights[x][y] == 1 && start == "toggle") {
					lights[x][y] = 0
				}
			}
		}
	}
	return sumField(lights)
}

func level6b(data string) int {
	var lights [1000][1000]int
	data = strings.ReplaceAll(data, "turn", "")
	for _, line := range strings.Split(data, "\n") {
		var start string
		var x1, y1, x2, y2 int
		_, _ = fmt.Sscanf(line, "%s %d,%d through %d,%d", &start, &x1, &y1, &x2, &y2)
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				if start == "on" {
					lights[x][y] += 1
				} else if start == "off" {
					lights[x][y] = MaxOf(0, lights[x][y]-1)
				} else if start == "toggle" {
					lights[x][y] += 2
				}
			}
		}
	}
	return sumField(lights)
}

func sumField(lights [1000][1000]int) int {
	total := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			total += lights[x][y]
		}
	}
	return total
}

var wires map[string]func(string) uint16
var solved map[string]uint16

func resolveWire(wire string) uint16 {
	intVal, err := strconv.Atoi(wire)
	if err == nil {
		return uint16(intVal)
	}
	if val, ok := solved[wire]; ok {
		return val
	}
	w, ok := wires[wire]
	if ok {
		val := w(wire)
		solved[wire] = val
		return val
	} else {
		panic("unknown wire: " + wire)
	}
}

func level7a(data string) int {
	wires = make(map[string]func(string) uint16)
	solved = make(map[string]uint16)
	for _, line := range strings.Split(data, "\n") {
		var parts = strings.Split(line, " -> ")
		var name = parts[1]
		var op = strings.Split(parts[0], " ")
		if len(op) == 1 {
			wireFunction1(op, name)
		} else if len(op) == 2 {
			wireFunction2(name, op)
		} else if len(op) == 3 {
			wireFunction3(op, name)
		}
	}
	return int(resolveWire("a"))
}

func wireFunction1(op []string, name string) {
	var value, err = strconv.Atoi(op[0])
	if err == nil {
		solved[name] = uint16(value)
	} else {
		wires[name] = func(wire string) uint16 {
			return resolveWire(op[0])
		}
	}
}

func wireFunction2(name string, op []string) {
	wires[name] = func(name string) uint16 {
		return ^resolveWire(op[1])
	}
}

func wireFunction3(op []string, name string) {
	switch op[1] {
	case "AND":
		wires[name] = func(name string) uint16 {
			return resolveWire(op[0]) & resolveWire(op[2])
		}
	case "OR":
		wires[name] = func(name string) uint16 {
			return resolveWire(op[0]) | resolveWire(op[2])
		}
	case "LSHIFT":
		wires[name] = func(name string) uint16 {
			return resolveWire(op[0]) << resolveWire(op[2])
		}
	case "RSHIFT":
		wires[name] = func(name string) uint16 {
			return resolveWire(op[0]) >> resolveWire(op[2])
		}
	default:
		panic("unknown operator: " + op[1])
	}
}

func level7b(data string) int {
	res := level7a(data)
	return level7a(data + "\n" + strconv.Itoa(res) + " -> b")
}

func unescape(s string) string {
	s = s[1 : len(s)-1]
	var hex = regexp.MustCompile("\\\\x[0-9a-f]{2}")
	s = strings.ReplaceAll(s, "\\\"", "\"")
	s = strings.ReplaceAll(s, "\\\\", "\\")
	s = hex.ReplaceAllString(s, "_")
	return s
}

func level8a(data string) int {
	var total = 0
	for _, line := range strings.Split(data, "\n") {
		total += len(line) - len(unescape(line))
	}
	return total
}

func escape(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return "\"" + s + "\""
}

func level8b(data string) int {
	var total = 0
	for _, line := range strings.Split(data, "\n") {
		total += len(escape(line)) - len(line)
	}
	return total
}

var paths map[string]map[string]int

func addPath(from string, to string, value int) {
	_, ok := paths[from]
	if !ok {
		paths[from] = make(map[string]int)
	}
	paths[from][to] += value
	// set back path
	_, ok = paths[to]
	if !ok {
		paths[to] = make(map[string]int)
	}
	paths[to][from] += value
}

func bfsDistances(curr string, visited map[string]bool, total int, toVisit int, aggFunc func(vars ...int) int) int {
	visitedCopy := copyMap(visited)
	visitedCopy[curr] = true

	if toVisit == 0 {
		return total
	}

	var distances []int
	var targets, pathExists = paths[curr]
	if pathExists {
		for destination, distance := range targets {
			if !visitedCopy[destination] {
				distances = append(distances, bfsDistances(destination, visitedCopy, total+distance, toVisit-1, aggFunc))
			}
		}
	}
	return aggFunc(distances...)
}

func copyMap[V any](original map[string]V) map[string]V {
	cpy := make(map[string]V)
	for k, v := range original {
		cpy[k] = v
	}
	return cpy
}

func level9(data string, aggFunc func(vars ...int) int) int {
	paths = make(map[string]map[string]int)
	// create map of paths
	for _, line := range strings.Split(data, "\n") {
		var from, to string
		var distance int
		_, _ = fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance)
		addPath(from, to, distance)
	}

	var dist []int
	// create all possible routes
	for from := range paths {
		var visited = make(map[string]bool)
		dist = append(dist, bfsDistances(from, visited, 0, len(paths)-1, aggFunc))
	}
	return aggFunc(dist...)
}

func level9a(data string) int {
	return level9(data, MinOf)
}

func level9b(data string) int {
	return level9(data, MaxOf)
}

func level10(data string) string {
	data += "x"
	var last = data[0]
	var count = 1
	var result = make([]string, 10_000_000)
	var currPos uint = 0
	for i := 1; i < len(data); i++ {
		if data[i] == last {
			count += 1
			continue
		} else {
			result[currPos] = strconv.Itoa(count) + string(last)
			currPos += 1
			last = data[i]
			count = 1
		}
	}
	return strings.Join(result[:currPos], "")
}

func level10a(data string) int {
	for i := 0; i < 40; i++ {
		data = level10(data)
	}
	return len(data)
}

func level10b(data string) int {
	for i := 0; i < 50; i++ {
		data = level10(data)
	}
	return len(data)
}

func isValid(s string) bool {
	// check if it contains any of the forbidden letters
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' || s[i] == 'o' || s[i] == 'l' {
			return false
		}
	}
	// count non-overlapping pairs
	var pairs = 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			pairs += 1
			i += 1
		}
	}
	// check for abc, bcd, cde, etc
	var followUp = false
	for i := 0; i < len(s)-2; i++ {
		if s[i]+1 == s[i+1] && s[i+1]+1 == s[i+2] {
			followUp = true
		}
	}
	return followUp && pairs >= 2
}

func increment(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'z' {
			s = s[:i] + "a" + s[i+1:]
		} else {
			s = s[:i] + string(s[i]+1) + s[i+1:]
			break
		}
	}
	return s
}

func level11a(data string) int {
	var curr = data
	for !isValid(curr) {
		curr = increment(curr)
	}
	println(curr)
	return 0
}

func level11b(data string) int {
	var curr = data
	for !isValid(curr) {
		curr = increment(curr)
	}
	println(increment(curr))
	return 0
}

func level12a(data string) int {
	var total = 0
	re := regexp.MustCompile("-?\\d+")
	for _, line := range strings.Split(data, "\n") {
		for _, numericString := range re.FindAll([]byte(line), -1) {
			var num int
			_, _ = fmt.Sscanf(string(numericString), "%d", &num)
			total += num
		}
	}
	return total
}

func recursiveSum(obj interface{}) int {
	var total = 0
	switch obj := obj.(type) {
	case []interface{}:
		for _, v := range obj {
			total += recursiveSum(v)
		}
	case map[string]interface{}:
		for _, v := range obj {
			total += recursiveSum(v)
			if v == "red" {
				return 0
			}
		}
	case float64:
		total += int(obj)
	}
	return total
}

func level12b(data string) int {
	content := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(data), &content)
	if err != nil {
		panic(err)
	}
	return recursiveSum(content)
}

func readNamesToPaths(data string) []string {
	paths = make(map[string]map[string]int)
	data = strings.ReplaceAll(data, "would gain ", "+")
	data = strings.ReplaceAll(data, "would lose ", "-")
	data = strings.ReplaceAll(data, " happiness units by sitting next to ", " ")
	data = strings.ReplaceAll(data, ".", "")
	for _, line := range strings.Split(data, "\n") {
		var from, to string
		var happiness int
		_, _ = fmt.Sscanf(line, "%s %d %s", &from, &happiness, &to)
		addPath(from, to, happiness)
	}
	// create names list
	var names = make([]string, 0, len(paths))
	for name := range paths {
		names = append(names, name)
	}
	return names
}

func findBestSeatingArrangement(names []string) int {
	var totals []int
	for i := 0; i < 100_000; i++ {
		rand.Shuffle(len(names), func(i, j int) {
			names[i], names[j] = names[j], names[i]
		})
		var total = 0
		for i := 0; i < len(names); i++ {
			var a = names[i]
			var b = names[(i+1)%len(names)]
			total += paths[a][b]
		}
		totals = append(totals, total)
	}
	return MaxOf(totals...)
}

func level13a(data string) int {
	names := readNamesToPaths(data)
	return findBestSeatingArrangement(names)
}

func level13b(data string) int {
	names := readNamesToPaths(data)
	// add myself
	for name := range paths {
		addPath("me", name, 0)
	}
	names = append(names, "me")
	return findBestSeatingArrangement(names)
}

type Reindeer struct {
	name                                    string
	speed, duration, rest, position, points int
	running                                 bool
}

func createReindeer(line string) Reindeer {
	var parts = strings.Split(line, " ")
	speed, duration, rest := AtoiPanic(parts[3]), AtoiPanic(parts[6]), AtoiPanic(parts[13])
	reindeer := Reindeer{parts[0], speed, duration, rest, 0, 0, false}
	return reindeer
}

func level14a(data string) int {
	var times []int
	for _, line := range strings.Split(data, "\n") {
		r := createReindeer(line)

		totalDist := 0
		for i := 0; i < 2503; i += r.duration + r.rest {
			totalDist += r.speed * MinOf(r.duration, 2503-i)
		}
		times = append(times, totalDist)
	}
	return MaxOf(times...)
}

func (reindeer *Reindeer) run(i int) {
	var offset = i % (reindeer.duration + reindeer.rest)
	if offset == 0 {
		reindeer.running = true
	}
	if offset == reindeer.duration {
		reindeer.running = false
	}
	if reindeer.running {
		reindeer.position += reindeer.speed
	}
}

func reindeerPos(r Reindeer) int    { return r.position }
func reindeerPoints(r Reindeer) int { return r.points }

func level14b(data string) int {
	var reindeerArr = make([]Reindeer, 0)
	for _, line := range strings.Split(data, "\n") {
		reindeerArr = append(reindeerArr, createReindeer(line))
	}

	for i := 0; i <= 2503; i++ {
		for j := range reindeerArr {
			reindeerArr[j].run(i)
		}
		var leadingPosition = MaxOf(Map(reindeerArr, reindeerPos)...)
		for j := range reindeerArr {
			if reindeerArr[j].position == leadingPosition {
				reindeerArr[j].points++
			}
		}
	}
	return MaxOf(Map(reindeerArr, reindeerPoints)...)
}

// level 15

type Ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

func (i Ingredient) score() int {
	return MaxOf(0, i.capacity) * MaxOf(0, i.durability) * MaxOf(0, i.flavor) * MaxOf(0, i.texture)
}

func (i Ingredient) plus(b Ingredient, weight int) Ingredient {
	return Ingredient{"", i.capacity + b.capacity*weight, i.durability + b.durability*weight, i.flavor + b.flavor*weight, i.texture + b.texture*weight, i.calories + b.calories*weight}
}

func readIngredients(data string) []Ingredient {
	var ingredients = make([]Ingredient, 0)
	for _, line := range strings.Split(data, "\n") {
		var ing Ingredient
		_, _ = fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &ing.name,
			&ing.capacity, &ing.durability, &ing.flavor, &ing.texture, &ing.calories)
		ingredients = append(ingredients, ing)
	}
	return ingredients
}

func mixIngredients(is []Ingredient) Ingredient {
	var a, b, c = rand.Uint32() % 50, rand.Uint32() % 50, rand.Uint32() % 50
	var d = 100 - a - b - c
	if d > 0 {
		return Ingredient{}.plus(is[0], int(a)).plus(is[1], int(b)).plus(is[2], int(c)).plus(is[3], int(d))
	}
	return Ingredient{}
}

func level15(data string, noLimit bool) int {
	ingredients := readIngredients(data)

	var bestScore = 0
	for i := 0; i < 1_000_000; i++ {
		mix := mixIngredients(ingredients)
		if mix.calories == 500 || noLimit {
			bestScore = MaxOf(bestScore, mix.score())
		}
	}
	return bestScore
}

func level15a(data string) int {
	return level15(data, true)
}

func level15b(data string) int {
	return level15(data, false)
}

// level 16
// noinspection SpellCheckingInspection
var aunt = map[string]int{"children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3, "akitas": 0, "vizslas": 0, "goldfish": 5, "trees": 3, "cars": 2, "perfumes": 1} // noinspection SpellCheckingInspection

func checkAunt(parts []string) bool {
	for i := 1; i < len(parts); i += 2 {
		name, value := parts[i], parts[i+1]
		for k, v := range aunt {
			if k == name && v != AtoiPanic(value) {
				return false
			}
		}
	}
	return true
}

func level16a(data string) int {
	data = strings.ReplaceAll(data, "Sue ", "")
	data = strings.ReplaceAll(data, ":", "")
	data = strings.ReplaceAll(data, ",", "")
	for _, line := range strings.Split(data, "\n") {
		var parts = strings.Split(line, " ")
		if checkAunt(parts) {
			return AtoiPanic(parts[0])
		}
	}
	return -1
}

func checkAunt2(parts []string) bool {
	for i := 1; i < len(parts); i += 2 {
		name, v2 := parts[i], AtoiPanic(parts[i+1])
		for k, v := range aunt {
			if k == name {
				if k == "cats" || k == "trees" {
					if v >= v2 {
						return false
					}
				} else // noinspection SpellCheckingInspection
				if k == "pomeranians" || k == "goldfish" {
					if v <= v2 {
						return false
					}
				} else if v != v2 {
					return false
				}
			}
		}
	}
	return true
}

func level16b(data string) int {
	data = strings.ReplaceAll(data, "Sue ", "")
	data = strings.ReplaceAll(data, ":", "")
	data = strings.ReplaceAll(data, ",", "")
	for _, line := range strings.Split(data, "\n") {
		var parts = strings.Split(line, " ")
		if checkAunt2(parts) {
			return AtoiPanic(parts[0])
		}
	}
	return -1
}

// level 17
var smallestContainerCombination = math.MaxInt
var containerCombinations = 0

func combineContainers(free []int, used, after, total, checksum int) {
	if total == 150 {
		smallestContainerCombination = MinOf(smallestContainerCombination, used)
		containerCombinations += 1
		return
	} else if total > 150 {
		return
	}
	// copy the free list
	for i := 0; i < len(free); i++ {
		if free[i] > 0 && i >= after {
			var free2 = make([]int, len(free))
			copy(free2, free)
			var fc = free2[i]
			free2[i] = 0
			combineContainers(free2, used+1, i, total+free[i], checksum*fc)
		}
	}
}

func level17a(data string) int {
	containerCombinations = 0
	var containers = make([]int, 0)
	for _, line := range strings.Split(data, "\n") {
		var v, _ = strconv.Atoi(line)
		containers = append(containers, v)
	}
	sort.Sort(sort.IntSlice(containers))
	combineContainers(containers, 0, 0, 0, 1)
	return containerCombinations
}

func level17b(data string) int {
	level17a(data)
	return smallestContainerCombination
}

// level18
func boolSum(b ...bool) int {
	var s = 0
	for _, v := range b {
		if v {
			s += 1
		}
	}
	return s
}

func readGameOfLife(data string) [102][102]bool {
	var l = [102][102]bool{}
	for x, line := range strings.Split(data, "\n") {
		for y, c := range line {
			if c == '#' {
				l[x+1][y+1] = c == '#'
			}
		}
	}
	return l
}

func playGameOfLife(l [102][102]bool, stuckPixel bool) [102][102]bool {
	for t := 0; t < 100; t++ {
		var n = [102][102]bool{}
		fixPixel(&l, stuckPixel)
		for x := 1; x <= 100; x++ {
			for y := 1; y <= 100; y++ {
				var nrNeigh = boolSum(l[x-1][y-1], l[x-1][y], l[x-1][y+1], l[x][y-1], l[x][y+1], l[x+1][y-1], l[x+1][y], l[x+1][y+1])
				n[x][y] = nrNeigh == 3 || (nrNeigh == 2 && l[x][y])
			}
		}
		fixPixel(&n, stuckPixel)
		l = n
	}
	return l
}

func fixPixel(l *[102][102]bool, stuckPixel bool) {
	if stuckPixel {
		l[1][1] = true
		l[1][100] = true
		l[100][1] = true
		l[100][100] = true
	}
}

func sumField102(l [102][102]bool) int {
	var sum int
	for x := 1; x <= 100; x++ {
		for y := 1; y <= 100; y++ {
			sum += boolSum(l[x][y])
		}
	}
	return sum
}

func level18a(data string) int {
	return sumField102(playGameOfLife(readGameOfLife(data), false))
}

func level18b(data string) int {
	return sumField102(playGameOfLife(readGameOfLife(data), true))
}

// level 19

type RulePair struct {
	orig, sub string
}

type RuleSet []RulePair

func createRules(data string) (string, RuleSet) {
	var rules = make(RuleSet, 0)
	var final = false
	var original string
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			final = true
			continue
		}
		if final {
			original = line
			break
		}
		var parts = strings.Split(line, " => ")
		rules = append(rules, RulePair{parts[0], parts[1]})
	}
	return original, rules
}

func (rules RuleSet) applyRules(original string) map[string]bool {
	var nextVariants = make(map[string]bool)
	for i := 0; i < len(original); i++ {
		var pre = original[:i]
		var post = original[i:]
		var remaining = len(original) - i
		for _, rule := range rules {
			if remaining >= len(rule.orig) && rule.orig == post[:len(rule.orig)] {
				var replaced = pre + rule.sub + post[len(rule.orig):]
				nextVariants[replaced] = true
			}
		}
	}
	return nextVariants
}

func (rules RuleSet) applyRule(original string) string {
	for i := len(original); i >= 0; i-- {
		var pre = original[:i]
		var post = original[i:]
		var remaining = len(original) - i
		for _, rule := range rules {
			if remaining >= len(rule.orig) && rule.orig == post[:len(rule.orig)] {
				var replaced = pre + rule.sub + post[len(rule.orig):]
				return replaced
			}
		}
	}
	return ""
}

func level19a(data string) int {
	var original, rules = createRules(data)
	nextVariants := rules.applyRules(original)
	return len(nextVariants)
}

func level19b(data string) int {
	var original, rules = createRules(data)
	var invertedRules RuleSet
	for _, rule := range rules {
		invertedRules = append(invertedRules, RulePair{rule.sub, rule.orig})
	}
	for time := 1; time < 10_000; time++ {
		original = invertedRules.applyRule(original)
		if original == "e" {
			return time
		}
	}
	return -1
}

// level 20

const MaxHouses = 1_000_000 // upper bound to speed up

func deliverPresents(housesPerElf int, nrElfs int, presentsPerElf int, houses *[MaxHouses]int) {
	for elf := 1; elf < nrElfs; elf++ {
		for house := elf; house < MinOf(housesPerElf*elf, MaxHouses); house += elf {
			houses[house] += elf * presentsPerElf
		}
	}
}

func level20a(data string) int {
	var targetNumber, _ = strconv.Atoi(data)
	var houses = [MaxHouses]int{}
	deliverPresents(1_000_000, 1_000_000, 10, &houses)
	for i := 1; i < MaxHouses; i++ {
		if houses[i] >= targetNumber {
			return i
		}
	}
	return -1
}

func level20b(data string) int {
	var targetNumber, _ = strconv.Atoi(data)
	var houses = [MaxHouses]int{}
	deliverPresents(50, 1_000_000, 11, &houses)
	for i := 1; i < MaxHouses; i++ {
		if houses[i] >= targetNumber {
			return i
		}
	}
	return -1
}

func level23(data string, reg map[string]int) int {
	data = strings.ReplaceAll(data, ",", "")
	var lines = strings.Split(data, "\n")
	for pc := 0; pc < len(lines); pc++ {
		var parts = strings.Split(lines[pc], " ")
		var id = parts[1]
		var value = 0
		if len(parts) > 2 {
			value, _ = strconv.Atoi(parts[2])
		}
		switch parts[0] {
		case "hlf":
			reg[id] /= 2
		case "tpl":
			reg[id] *= 3
		case "inc":
			reg[id] += 1
		case "jmp":
			var iid, _ = strconv.Atoi(id)
			pc += iid - 1
		case "jie":
			if reg[id]%2 == 0 {
				pc += value - 1
			}
		case "jio":
			if reg[id] == 1 {
				pc += value - 1
			}
		}
	}
	return reg["b"]
}

type Item struct {
	name                string
	cost, damage, armor int
}

func (a Item) plus(b Item) Item {
	return Item{a.name + " " + b.name, a.cost + b.cost, a.damage + b.damage, a.armor + b.armor}
}

// noinspection SpellCheckingInspection
var weapons = []Item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

// noinspection SpellCheckingInspection
var armors = []Item{
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
	{"NoArmor", 0, 0, 0},
}

var rings = []Item{
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
	{"NoRing1", 0, 0, 0},
	{"NoRing2", 0, 0, 0},
}

func simulateFight(player Item, bossHitPoints int, damage int, armor int) bool {
	var playerHealth = 100
	var playerDamage = MaxOf(1, player.damage-armor)
	var bossDamage = MaxOf(1, damage-player.armor)
	for {
		bossHitPoints -= playerDamage
		if bossHitPoints <= 0 {
			return true
		}
		playerHealth -= bossDamage
		if playerHealth <= 0 {
			return false
		}
	}
}

func combineRecursive(items ...[]Item) []Item {
	if len(items) == 0 {
		return []Item{{}}
	}
	var result = make([]Item, 0)
	for _, item := range items[0] {
		for _, rest := range combineRecursive(items[1:]...) {
			if strings.Contains(rest.name, item.name) {
				continue
			}
			result = append(result, item.plus(rest))
		}
	}
	return result
}

func level21(data string, partB bool) int {
	var bossHitPoints, bossDamage, bossArmor int
	var cheapest, mostExpensive = math.MaxInt, 0
	_, _ = fmt.Sscanf(data, "Hit Points: %d\nDamage: %d\nArmor: %d", &bossHitPoints, &bossDamage, &bossArmor)

	// create list of all possible combinations of items
	items := combineRecursive(weapons, armors, rings, rings)

	for _, total := range items {
		if simulateFight(total, bossHitPoints, bossDamage, bossArmor) {
			cheapest = MinOf(cheapest, total.cost)
		} else {
			mostExpensive = MaxOf(mostExpensive, total.cost)
		}
	}
	if partB {
		return mostExpensive
	} else {
		return cheapest
	}
}

func level21a(data string) int { return level21(data, false) }
func level21b(data string) int { return level21(data, true) }

// level 22

type GameState struct {
	playerHitPoints, bossHitPoints          int
	bossDamage, mana                        int
	myTurn                                  bool
	manaSpent, shield                       int
	shieldTurns, poisonTurns, rechargeTurns int
}

func pushNextState(queue *list.List, originalState GameState, manaCost int, effectFunction func(state *GameState)) {
	if originalState.mana >= manaCost {
		var state = originalState
		state.mana -= manaCost
		state.manaSpent += manaCost
		effectFunction(&state)
		queue.PushBack(state)
	}
}

func simulateMagicFight(state GameState, queue *list.List) {
	// tick effects
	if state.rechargeTurns > 0 {
		state.mana += 101
	}
	if state.shieldTurns > 0 {
		state.shield = 7
	} else {
		state.shield = 0
	}
	if state.poisonTurns > 0 {
		state.bossHitPoints -= 3
	}

	state.rechargeTurns = MaxOf(0, state.rechargeTurns-1)
	state.shieldTurns = MaxOf(0, state.shieldTurns-1)
	state.poisonTurns = MaxOf(0, state.poisonTurns-1)

	if state.myTurn {
		state.myTurn = false
		pushNextState(queue, state, 53, func(state *GameState) { state.bossHitPoints -= 4 })
		pushNextState(queue, state, 73, func(state *GameState) { state.bossHitPoints -= 2; state.playerHitPoints += 2 })
		if state.poisonTurns <= 0 {
			pushNextState(queue, state, 113, func(state *GameState) { state.poisonTurns = 6 })
		}
		if state.shieldTurns <= 0 {
			pushNextState(queue, state, 173, func(state *GameState) { state.shieldTurns = 6 })
		}
		if state.rechargeTurns <= 0 {
			pushNextState(queue, state, 229, func(state *GameState) { state.rechargeTurns = 5 })
		}
	} else {
		state.myTurn = true
		state.playerHitPoints -= state.bossDamage - state.shield
		queue.PushBack(state)
	}
}

func level22(data string, hard bool) int {
	var bossHitPoints, bossDamage int
	_, _ = fmt.Sscanf(data, "Hit Points: %d\nDamage: %d", &bossHitPoints, &bossDamage)
	var minManaSpent = math.MaxInt

	var startState = GameState{50, bossHitPoints, bossDamage, 500, true,
		0, 0, 0, 0, 0}
	var queue = list.List{}
	queue.PushBack(startState)

	for queue.Len() > 0 {
		// pop the first element
		var state = queue.Remove(queue.Front()).(GameState)
		if hard && state.myTurn {
			state.playerHitPoints--
		}
		if state.playerHitPoints <= 0 {
			continue
		}
		if state.bossHitPoints <= 0 {
			if state.manaSpent < minManaSpent {
				minManaSpent = state.manaSpent
			}
			continue
		} else if state.manaSpent >= minManaSpent {
			continue // this trims the search space a lot
		} else {
			simulateMagicFight(state, &queue)
		}
	}
	return minManaSpent
}

func level22a(data string) int { return level22(data, false) }
func level22b(data string) int { return level22(data, true) } // doesn't work properly for some reason
func level23a(data string) int { return level23(data, make(map[string]int)) }
func level23b(data string) int { return level23(data, map[string]int{"a": 1}) }

func level24(data string, groupings int) int {
	var packages = make([]int, 0)
	for _, line := range strings.Split(data, "\n") {
		var v, _ = strconv.Atoi(line)
		packages = append(packages, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(packages)))
	var targetWeight = Sum(packages...) / groupings
	var best = combineUntilTarget(packages, targetWeight, 0, 0, 1)

	return best
}

func combineUntilTarget(packages []int, targetWeight int, offset int, sum, prod int) int {
	var best = math.MaxInt
	if sum == targetWeight && prod > 0 {
		best = MinOf(best, prod)
	}
	if offset >= len(packages) || sum >= targetWeight {
		return best
	}
	for i := offset; i < len(packages); i++ {
		best = MinOf(best, combineUntilTarget(packages, targetWeight, i+1, sum+packages[i], prod*packages[i]))
	}
	return best
}

func level24a(data string) int { return level24(data, 3) }
func level24b(data string) int { return level24(data, 4) }

func level25a(data string) int {
	var row, col int
	_, _ = fmt.Sscanf(data, "To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.", &row, &col)
	var currNum = 20151125
	var targetIndex = (row+col-1)*(row+col-2)/2 + col
	for i := 1; i < targetIndex; i++ {
		currNum = (currNum * 252533) % 33554393
	}
	return currNum
}
func level25b(data string) int {
	data = data + data
	return 0
} // noqa
