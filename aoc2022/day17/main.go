package main

import (
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"os"
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

	game := newGame()
	for _, line := range lines {
		fmt.Printf("%v\n", game.run(line, 1000000000000))
	}
}

type Game struct {
	grid    map[Position]bool
	minX    int
	maxX    int
	maxY    int
	profile []int
	rock    Rock
	rockPos Position
	patterns map[Input]Pattern
}

type Rock map[Position]bool

func newGame() *Game {
	return &Game{
		grid: make(map[Position]bool),
		minX: 0,
		maxX: 6,
		maxY: -1,
		profile: []int{ -1, -1, -1, -1, -1, -1, -1 },
		rock: nil,
		patterns: make(map[Input]Pattern),
	}
}

func (this *Game) run(jetPatterns string, numOfRocks int) int {
	rocks := []Rock{
		{
			newPosition(0, 0): true,
			newPosition(1, 0): true,
			newPosition(2, 0): true,
			newPosition(3, 0): true,
		},
		{
			newPosition(1, 0): true,
			newPosition(0, 1): true,
			newPosition(1, 1): true,
			newPosition(2, 1): true,
			newPosition(1, 2): true,
		},
		{
			newPosition(0, 0): true,
			newPosition(1, 0): true,
			newPosition(2, 0): true,
			newPosition(2, 1): true,
			newPosition(2, 2): true,
		},
		{
			newPosition(0, 0): true,
			newPosition(0, 1): true,
			newPosition(0, 2): true,
			newPosition(0, 3): true,
		},
		{
			newPosition(0, 0): true,
			newPosition(0, 1): true,
			newPosition(1, 0): true,
			newPosition(1, 1): true,
		},
	}
	for i, j := 0, 0; i < numOfRocks; i++ {
		rockIdx := i % len(rocks)
		rock := rocks[i%len(rocks)]

		pattern, ok := this.patterns[newInput(rockIdx, j % len(jetPatterns))]
		if ok {
			diff := i - pattern.seq
			if equalProfile(this.profile, pattern.profile) && ((numOfRocks - i) % diff) == 0 {
				return this.maxY + (numOfRocks - i)/diff * (this.maxY - pattern.maxY)
			}
		}
		this.patterns[newInput(rockIdx, j % len(jetPatterns))] = Pattern{
			seq: i,
			maxY: this.maxY,
			profile: this.profile,
		}
				
		this.addRock(rock)
		for this.runOneStep(jetPatterns[j%len(jetPatterns)]) {
			j++
		}
		j++
	}
	return this.maxY
}

func (this *Game) addRock(rock Rock) {
	if this.rock != nil {
		panic("")
	}
	this.rock = rock
	this.rockPos = newPosition(2, this.maxY+4)
}

func (this *Game) runOneStep(jetPattern byte) bool {
	if jetPattern == '<' {
		this.moveLeft()
	} else if jetPattern == '>' {
		this.moveRight()
	} else {
		panic(fmt.Sprintf("bad jet pattern: %v", jetPattern))
	}
	return this.moveDown()
}

func (this *Game) moveLeft() {
	newRockPos := newPosition(this.rockPos.x-1, this.rockPos.y)
	if !this.detectCollision(newRockPos) {
		this.rockPos = newRockPos
	}
}

func (this *Game) moveRight() {
	newRockPos := newPosition(this.rockPos.x+1, this.rockPos.y)
	if !this.detectCollision(newRockPos) {
		this.rockPos = newRockPos
	}
}

func (this *Game) moveDown() bool {
	newRockPos := newPosition(this.rockPos.x, this.rockPos.y-1)
	if !this.detectCollision(newRockPos) {
		this.rockPos = newRockPos
		return true
	}

	for p := range this.rock {
		absPos := newPosition(this.rockPos.x+p.x, this.rockPos.y+p.y)
		this.grid[absPos] = true
		if absPos.y > this.maxY {
			this.maxY = absPos.y
		}
		if absPos.y > this.profile[absPos.x] {
			this.profile[absPos.x] = absPos.y
		}
	}
	this.rock = nil
	return false
}

func (this *Game) detectCollision(newRockPos Position) bool {
	for p := range this.rock {
		absPos := newPosition(newRockPos.x+p.x, newRockPos.y+p.y)
		if absPos.x < this.minX || absPos.x > this.maxX || absPos.y < 0 {
			return true
		}
		_, ok := this.grid[absPos]
		if ok {
			return true
		}
	}
	return false
}

type Position struct {
	x, y int
}

func newPosition(x, y int) Position {
	return Position{
		x: x,
		y: y,
	}
}

type Input struct {
	rockIdx int
	jetIdx int
}

func newInput(a, b int) Input {
	return Input{
		rockIdx: a,
		jetIdx: b,
	}
}

func equalProfile(profile1, profile2 []int) bool {
	if len(profile1) != len(profile2) {
		return false
	}

	diff := 0
	for i := range profile1 {
		if i == 0 {
			diff = profile1[i] - profile2[i]
		} else if diff != profile1[i] - profile2[i] {
			return false
		}
	}
	return true
}
			
type Pattern struct {
	seq int
	profile []int
	maxY int
}
