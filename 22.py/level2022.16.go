package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type MapSI map[string]int
type MapSSI map[string]MapSI

// global, read only, to avoid passing around as parameters
var cost = make(MapSSI)
var reward = make(MapSI)

type State struct {
	start   string
	toVisit MapSI
	time    int
	score   int
}

type Result struct {
	score   int
	visited MapSI
}

func (r Result) string() string { // for hashing, to eliminate duplicates
	var keys []string
	keys = append(keys, strconv.Itoa(r.score))
	for k := range r.visited {
		keys = append(keys, k)
	}
	return strings.Join(keys, ",")
}

func setDifference[K comparable, V any](a map[K]int, b map[K]V) map[K]V { // set a without b
	result := make(map[K]V)
	for key := range a {
		if v, ok := b[key]; !ok {
			result[key] = v
		}
	}
	return result
}

func copyWithout[K comparable, V any](ss map[K]V, without K) map[K]V { // copy a set without a key
	var toVisitCopy = make(map[K]V)
	for k, v := range ss {
		if k != without {
			toVisitCopy[k] = v
		}
	}
	return toVisitCopy
}

func values[V any](mr map[string]V) []V { // values of a map
	var result []V
	for _, v := range mr {
		result = append(result, v)
	}
	return result
}

func isDisjoint[K comparable, V any](m0 map[K]V, m1 map[K]V) bool { // are two sets/maps disjoint?
	for k := range m0 {
		if _, ok := m1[k]; ok {
			return false
		}
	}
	for k := range m1 {
		if _, ok := m0[k]; ok {
			return false
		}
	}
	return true
}

func worker(s State, results chan<- Result, wg *sync.WaitGroup) { // runner for parallel goroutines
	for target := range s.toVisit {
		if reward[target] == 0 {
			continue
		}
		var timeNext = s.time - cost[s.start][target] - 1
		if timeNext > 0 {
			wg.Add(1)
			var toVisit = copyWithout(s.toVisit, target)
			var score = s.score + reward[target]*timeNext
			go worker(State{target, toVisit, timeNext, score}, results, wg)
		}
	}
	results <- Result{s.score, setDifference(reward, s.toVisit)}
	wg.Done()
}

func walksShorterThan(toVisit MapSI, size int) chan Result {
	results := make(chan Result, 1_000_000)
	// Run parallel goroutines for recursive solution finding, using waitgroups to wait for end
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(State{"AA", toVisit, size, 0}, results, &wg)
	wg.Wait()      // wait until WaitGroup is 0
	close(results) // close channel to signal end of results, to make it readable
	return results
}

func parse(text string) (MapSSI, MapSI) {
	lines := strings.Split(text, "\n")
	graph := make(MapSSI)
	rewardMap := make(MapSI)

	// parse the input file, line by line, into a directed graph and a reward map
	re := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ([\w, ]+)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) == 0 {
			panic("Invalid line: " + line)
		}
		fromValve := match[1]
		costValve, _ := strconv.Atoi(match[2])
		for _, toValve := range strings.Split(match[3], ", ") {
			if _, ok := graph[fromValve]; !ok { // initialize map if not yet present
				graph[fromValve] = make(MapSI)
			}
			graph[fromValve][toValve] = 1
		}
		if costValve > 0 || fromValve == "AA" { // only add to reward map if it has a reward or is starting state
			rewardMap[fromValve] = costValve
		}
	}
	return addAnyToAnyDistance(graph), rewardMap
}

func addAnyToAnyDistance(graph MapSSI) MapSSI {
	// repeat until the graph doesn't change, future optimization: only for reward nodes + AA
	var again = true
	for again {
		again = false
		for start, middles := range graph {
			for middle, startMiddleDist := range middles {
				for end, middleEndDist := range graph[middle] {
					if start != end {
						var startEndDist = startMiddleDist + middleEndDist
						if oldVal, ok := graph[start][end]; !ok || oldVal > startEndDist {
							graph[start][end] = startEndDist
							again = true
						}
					}
				}
			}
		}
	}
	return graph
}

func main() {
	for _, file := range []string{"22.py/level2022.16.example", "22.py/level2022.16"} {
		var data, _ = os.ReadFile(file)
		cost, reward = parse(string(data))
		println(levelA(reward))
		println(levelB(reward))
	}

}

func levelA(toVisit MapSI) int {
	results := walksShorterThan(toVisit, 30)
	var best = 0
	for r := range results {
		if r.score > best {
			best = r.score
		}
	}
	return best
}

func levelB(toVisit MapSI) int {
	results := walksShorterThan(toVisit, 26)

	// get unique by toStringing them
	var uniqueSolutions = map[string]Result{}
	for r := range results {
		uniqueSolutions[r.string()] = r
	}
	// get values via reflection
	var solutions = values(uniqueSolutions)

	// results channel for goroutines
	var best = 0
	var wg sync.WaitGroup
	for is0, s0 := range solutions {
		wg.Add(1)
		go inner(solutions, is0, s0, &best, &wg)
	}
	wg.Wait()
	return best
}

func inner(solutions []Result, is0 int, s0 Result, best *int, wg *sync.WaitGroup) { // run a thread for each line
	for _, s1 := range solutions[is0+1:] {
		if (s0.score+s1.score) >= *best && // this optimization gives a 10x speedup
			isDisjoint(s0.visited, s1.visited) {
			score := s0.score + s1.score
			if score > *best {
				*best = score
			}
		}
	}
	wg.Done()
}
