package main

import (
	"os"
	"fmt"
	"github.com/learn_go/aoc2022/util"
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

	grid := newGrid()
	for i, line := range lines {
		row := []int{}
		for j, r := range line {
			if r == 'S' {
				grid.start = Position{ x: j, y: i }
				r = 'a'
			} else if r == 'E' {
				grid.destination = Position{ x: j, y: i }
				r = 'z'
			}
			row = append(row, int(r - 'a'))
		}
		grid.heightmap = append(grid.heightmap, row)
	}
	grid.numOfRows = len(grid.heightmap)
	grid.numOfColumns = len(grid.heightmap[0])
	fmt.Printf("%v\n", grid.computeFewestSteps())
	fmt.Printf("%v\n", grid.computeFewestSteps2())
}

type Grid struct {
	heightmap [][]int
	numOfRows int
	numOfColumns int
	start Position
	destination Position
}

type Position struct {
	x int
	y int
}

func newGrid() *Grid {
	return &Grid{}
}

func (g *Grid) computeFewestSteps() int {
	tentativeNodes := make(map[Position]bool)
	tentativeNodes[g.start] = true
	depth := 0
	visited := make(map[Position]bool)
	for len(tentativeNodes) > 0 {
		nextLevel := make(map[Position]bool)
		for cur := range tentativeNodes {
			if cur == g.destination {
				return depth
			}

			visited[cur] = true
			neighbors := g.getNeighbors(cur)
			for _, n := range neighbors {
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

func (g *Grid) computeFewestSteps2() int {
	tentativeNodes := make(map[Position]bool)
	tentativeNodes[g.destination] = true
	depth := 0
	visited := make(map[Position]bool)
	for len(tentativeNodes) > 0 {
		nextLevel := make(map[Position]bool)
		for cur := range tentativeNodes {
			if g.height(cur) == 0 {
				return depth
			}

			visited[cur] = true
			neighbors := g.getNeighbors2(cur)
			for _, n := range neighbors {
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

func (g *Grid) getNeighbors(p Position) []Position {
	myHeight := g.height(p)
	ret := []Position{}
	if n, e := g.top(p); e == nil && g.height(n) <= myHeight+1 {
		ret = append(ret, n)
	}
	if n, e := g.bottom(p); e == nil && g.height(n) <= myHeight+1 {
		ret = append(ret, n)
	}
	if n, e := g.left(p); e == nil && g.height(n) <= myHeight+1 {
		ret = append(ret, n)
	}
	if n, e := g.right(p); e == nil && g.height(n) <= myHeight+1 {
		ret = append(ret, n)
	}
	return ret
}

func (g *Grid) getNeighbors2(p Position) []Position {
	myHeight := g.height(p)
	ret := []Position{}
	if n, e := g.top(p); e == nil && g.height(n) >= myHeight-1 {
		ret = append(ret, n)
	}
	if n, e := g.bottom(p); e == nil && g.height(n) >= myHeight-1 {
		ret = append(ret, n)
	}
	if n, e := g.left(p); e == nil && g.height(n) >= myHeight-1 {
		ret = append(ret, n)
	}
	if n, e := g.right(p); e == nil && g.height(n) >= myHeight-1 {
		ret = append(ret, n)
	}
	return ret
}

func (g *Grid) height(p Position) int {
	return g.heightmap[p.y][p.x]
}

func (g *Grid) top(p Position) (Position, error) {
	if p.y <= 0 {
		return Position{}, fmt.Errorf("no top")
	}
	return Position{ x: p.x, y: p.y - 1 }, nil
}

func (g *Grid) bottom(p Position) (Position, error) {
	if p.y >= g.numOfRows - 1 {
		return Position{}, fmt.Errorf("no bottom")
	}
	return Position{ x: p.x, y: p.y + 1 }, nil
}

func (g *Grid) left(p Position) (Position, error) {
	if p.x <= 0 {
		return Position{}, fmt.Errorf("no left")
	}
	return Position{ x: p.x - 1, y: p.y }, nil
}

func (g *Grid) right(p Position) (Position, error) {
	if p.x >= g.numOfColumns - 1 {
		return Position{}, fmt.Errorf("no right")
	}
	return Position{ x: p.x + 1, y: p.y }, nil
}
