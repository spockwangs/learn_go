package main

import (
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"os"
	"strconv"
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

	grid := newGrid()
	for _, line := range lines {
		segments := parseSegments(line)
		grid.addSegmentsOfRocks(segments)
	}
	//fmt.Printf("%v\n", grid.simulate(newPosition(500, 0)))
	fmt.Printf("%v\n", grid.simulate2(newPosition(500, 0)))
}

func parseSegments(s string) []Position {
	ret := []Position{}
	splits := strings.Split(s, "->")
	for _, split := range splits {
		fields := strings.Split(split, ",")
		if len(fields) != 2 {
			panic("")
		}
		x, err := strconv.Atoi(strings.TrimSpace(fields[0]))
		if err != nil {
			panic("")
		}
		y, err := strconv.Atoi(strings.TrimSpace(fields[1]))
		if err != nil {
			panic("")
		}
		ret = append(ret, newPosition(x, y))
	}
	return ret
}

type Grid struct {
	blocks map[Position]Tile
	rockMaxY int
}

type Tile int

const (
	ROCK Tile = iota
	SAND
)

func newGrid() *Grid {
	return &Grid{
		blocks: make(map[Position]Tile),
		rockMaxY: 0,
	}
}

func (this *Grid) addSegmentsOfRocks(segments []Position) {
	for i := 0; i < len(segments)-1; i++ {
		this.addOneSegmentOfRocks(segments[i], segments[i+1])
	}
}

func (this *Grid) addOneSegmentOfRocks(a, b Position) {
	if a.x != b.x && a.y != b.y {
		panic("")
	}
	if a.x == b.x {
		for y := util.MinInt(a.y, b.y); y <= util.MaxInt(a.y, b.y); y++ {
			this.addRock(newPosition(a.x, y))
		}
	} else {
		for x := util.MinInt(a.x, b.x); x <= util.MaxInt(a.x, b.x); x++ {
			this.addRock(newPosition(x, a.y))
		}
	}
}

func (this *Grid) addRock(p Position) {
	this.blocks[p] = ROCK
	if p.y > this.rockMaxY {
		this.rockMaxY = p.y
	}
}

func (this *Grid) addSand(p Position) {
	this.blocks[p] = SAND
}

func (this *Grid) isEmpty(p Position) bool {
	_, ok := this.blocks[p]
	return !ok
}

func (this *Grid) isEmpty2(p Position) bool {
	_, ok := this.blocks[p]
	if ok {
		return false
	}
	if p.y == this.rockMaxY + 2 {
		return false
	}
	return true
}

func (this *Grid) simulate(start Position) int {
	sum := 0
out:
	for {
		cur := start
		for {
			if this.downForever(cur) {
				break out
			} else if n := cur.down(); this.isEmpty(n) {
				cur = n
			} else if n := cur.leftDown(); this.isEmpty(n) {
				cur = n
			} else if n := cur.rightDown(); this.isEmpty(n) {
				cur = n
			} else {
				this.addSand(cur)
				sum++
				break
			}
		}
	}
	return sum
}

func (this *Grid) simulate2(start Position) int {
	sum := 0
out:
	for {
		cur := start
		for {
			if n := cur.down(); this.isEmpty2(n) {
				cur = n
			} else if n := cur.leftDown(); this.isEmpty2(n) {
				cur = n
			} else if n := cur.rightDown(); this.isEmpty2(n) {
				cur = n
			} else {
				this.addSand(cur)
				sum++
				if cur == start {
					break out
				} else {
					break
				}
			}
		}
	}
	return sum
}

func (this *Grid) downForever(p Position) bool {
	for y := p.y; y <= this.rockMaxY; y++ {
		if !this.isEmpty(newPosition(p.x, y)) {
			return false
		}
	}
	return true
}

type Position struct {
	x int
	y int
}

func newPosition(x, y int) Position {
	return Position{
		x: x,
		y: y,
	}
}

func (this Position) down() Position {
	return newPosition(this.x, this.y+1)
}

func (this Position) leftDown() Position {
	return newPosition(this.x-1, this.y+1)
}

func (this Position) rightDown() Position {
	return newPosition(this.x+1, this.y+1)
}
