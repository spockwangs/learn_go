package main

import (
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"os"
	"strings"
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

func (this Valves) startFrom(start string, timeLimit int) int {
	leftValves := []string{}
	for k, v := range this.graph {
		if v.flowRate > 0 {
			leftValves = append(leftValves, k)
		}
	}
	return this.startFromHelper(start, timeLimit, leftValves)
}

func (this Valves) startFromHelper(start string, leftTime int, leftValves []string) int {
	bestScore := 0
	for i, next := range leftValves {
		dist := this.computeDistance(start, next)
		if leftTime > dist + 1 {
			score := (leftTime - dist - 1) * this.graph[next].flowRate +
				this.startFromHelper(next, leftTime-dist-1, remove(leftValves, i))
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
	return this.startFromHelper2(start, timeLimit, leftValves)
}

func (this Valves) startFromHelper2(start string, leftTime int, leftValves []string) int {
	bestScore := 0
	for i, next := range leftValves {
		dist := this.computeDistance(start, next)
		if leftTime > dist + 1 {
			score := (leftTime - dist - 1) * this.graph[next].flowRate +
				this.startFromHelper(next, leftTime-dist-1, remove(leftValves, i))
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

// Compute the shortest distance using BFS.
func (this Valves) computeDistance(src, dest string) int {
	d, ok := this.dist[makePair(src, dest)]
	if ok {
		return d
	}
	
	tentativeNodes := make(map[string]bool)
	tentativeNodes[src] = true
	depth := 0
	visited := make(map[string]bool)
	for len(tentativeNodes) > 0 {
		nextLevel := make(map[string]bool)
		for cur := range tentativeNodes {
			if src < cur {
				this.dist[makePair(src, cur)] = depth
			} else {
				this.dist[makePair(src, cur)] = depth
			}
			if cur == dest {
				return depth
			}

			visited[cur] = true
			neighbours := this.graph[cur].neighbours
			for _, n := range neighbours {
				_, ok := visited[n]
				if ok {
					continue
				}
				nextLevel[n] = true
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
