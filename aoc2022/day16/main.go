package main

import (
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"os"
	"strings"
	"sort"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v filename\n", os.Args[0])
		return
	}

	lines, err := util.ReadLines(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	valves := newValves()
	for _, line := range lines {
		valves.parseValve(line)
	}
	fmt.Printf("%v\n", valves.startFrom("AA", 30))
	fmt.Printf("%v\n", valves.startFrom2("AA", 26))
}

type Valves struct {
	graph map[string]Valve
	dist map[Pair]int
}

type Pair struct {
	a, b string
}

func makePair(a, b string) Pair {
	if a < b {
		return Pair{
			a: a,
			b: b,
		}
	} else {
		return Pair {
			a: b,
			b: a,
		}
	}
}
			
type Valve struct {
	name string
	flowRate int
	neighbours []string
}

func newValves() Valves {
	return Valves{
		dist: make(map[Pair]int),
	}
}

func (this *Valves) parseValve(s string) {
	next := 0
	util.ConsumeStr(s, "Valve ", &next)

	valve := Valve{}
	valve.name = util.ConsumeStrUntil(s, " ", &next)
	if !util.ConsumeStr(s, "has flow rate=", &next) {
		panic("")
	}
	valve.flowRate = util.ConsumeInt(s, &next)
	if !util.ConsumeStr(s, "; tunnels lead to valves ", &next) &&
		!util.ConsumeStr(s, "; tunnel leads to valve ", &next) {
		panic("")
	}
	valve.neighbours = strings.Split(s[next:], ", ")
	if this.graph == nil {
		this.graph = make(map[string]Valve)
	}
	this.graph[valve.name] = valve
}

func (this *Valves) startFrom(start string, timeLimit int) int {
	leftValves := []string{}
	for k, v := range this.graph {
		if v.flowRate > 0 {
			leftValves = append(leftValves, k)
		}
	}
	return this.startFromHelper(start, timeLimit, leftValves, nil)
}

func (this *Valves) startFromHelper(
	start string, leftTime int, leftValves []string, ctx *Context) int {
	bestScore := 0
	for i, next := range leftValves {
		dist := this.computeDistance(start, next)
		if leftTime > dist + 1 {
			curScore := (leftTime - dist - 1) * this.graph[next].flowRate
			var newCtx *Context = nil
			if ctx != nil {
				newCtx = ctx.appendPathScore(next, curScore)
			}
			score := curScore +
				this.startFromHelper(next, leftTime-dist-1, remove(leftValves, i), newCtx)
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

func (this Valves) startFrom2(start string, timeLimit int) int {
	leftValves := []string{}
	for k, v := range this.graph {
		if v.flowRate > 0 {
			leftValves = append(leftValves, k)
		}
	}

	ctx := newContext()
	this.startFromHelper(start, timeLimit, leftValves, ctx)

	pathScores := []PathScore{}
	for _, v := range ctx.pathScoreMap {
		pathScores = append(pathScores, v)
	}
	bestScore := 0 
	for i := range pathScores {
		for j := i+1; j < len(pathScores); j++ {
			path1 := pathScores[i].path
			path2 := pathScores[j].path
			if !intersect(path1, path2) {
				score := pathScores[i].score + pathScores[j].score
				if bestScore < score {
					bestScore = score
				}
			}
		}
	}
	return bestScore
}

// Compute the shortest distance using BFS.
func (this *Valves) computeDistance(src, dest string) int {
	d, ok := this.dist[makePair(src, dest)]
	if ok {
		return d
	}
	
	visited := make(map[string]bool)
	visited[src] = true
	tentativeNodes := make(map[string]bool)
	tentativeNodes[src] = true
	depth := 0
	for len(tentativeNodes) > 0 {
		nextLevel := make(map[string]bool)
		for cur := range tentativeNodes {
			this.dist[makePair(src, cur)] = depth
			if cur == dest {
				return depth
			}

			neighbours := this.graph[cur].neighbours
			for _, n := range neighbours {
				_, ok := visited[n]
				if ok {
					continue
				}
				nextLevel[n] = true
				visited[n] = true
			}
		}
		tentativeNodes = nextLevel
		depth++
	}
	return -1
}

// Make a copy of a slice, excluding the specified index.
func remove(s []string, idx int) []string {
	ret := make([]string, len(s))
	copy(ret, s)
	ret[idx] = ret[len(ret)-1]
	return ret[:len(ret)-1]
}

type Context struct {
	pathScoreMap map[string]PathScore
	path []string
	score int
}

type PathScore struct {
	path []string
	score int
}

func newContext() *Context {
	return &Context{
		pathScoreMap: make(map[string]PathScore),
	}
}

func (this *Context) appendPathScore(node string, score int) *Context {
	ret := &Context{
		pathScoreMap: this.pathScoreMap,
		path: make([]string, len(this.path)),
		score: this.score,
	}
	copy(ret.path, this.path)
	ret.path = append(ret.path, node)
	ret.score += score
	
	path := make([]string, len(ret.path))
	copy(path, ret.path)
	sort.Slice(path, func(i, j int) bool {
		return path[i] < path[j]
	})
	pathStr := strings.Join(path, "-")
	pathScore, ok := this.pathScoreMap[pathStr]
	if !ok || pathScore.score < ret.score {
		this.pathScoreMap[pathStr] = PathScore{
			path: ret.path,
			score: ret.score,
		}
	}
	return ret
}

func intersect(p1, p2 []string) bool {
	set := make(map[string]bool)
	for _, n := range p1 {
		set[n] = true
	}
	for _, n := range p2 {
		_, ok := set[n]
		if ok {
			return true
		}
	}
	return false
}
