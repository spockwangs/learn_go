package main

import (
	"os"
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"strings"
	"strconv"
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

	footprints := make(map[Position]int)
	rope := make([]Position, 2)
	footprints[Position{}]++
	rope10 := make([]Position, 10)
	footprints10 := make(map[Position]int)
	footprints10[Position{}]++
	for _, line := range lines {
		splits := strings.Split(line, " ")
		if len(splits) != 2 {
			panic("")
		}
		n, err := strconv.Atoi(splits[1])
		if err != nil {
			panic(err)
		}
		for i := 0; i < n; i++ {
			moveRope(rope, splits[0])
			footprints[rope[len(rope)-1]]++

			moveRope(rope10, splits[0])
			footprints10[rope10[len(rope10)-1]]++
		}
	}
	fmt.Printf("%v\n", len(footprints))
	fmt.Printf("%v\n", len(footprints10))
}

type Position struct {
	x, y int
}

func moveRope(rope []Position, direction string) {
	for j := range rope {
		p := &rope[j]
		if j == 0 {
			moveKnot(direction, p)
		} else {
			prev := rope[j-1]
			followKnot(prev, p)
		}
	}
}

func moveKnot(direction string, cur *Position) {
	switch direction {
	case "R":
		cur.x++
	case "L":
		cur.x--
	case "U":
		cur.y++
	case "D":
		cur.y--
	}
}

func followKnot(prev Position, cur *Position) {
	if prev.x == cur.x {
		if cur.y < prev.y - 1 {
			cur.y++
		} else if cur.y > prev.y + 1 {
			cur.y--
		}
	} else if prev.y == cur.y {
		if cur.x < prev.x - 1 {
			cur.x++
		} else if cur.x > prev.x + 1 {
			cur.x--
		}
	} else if util.Abs(cur.x - prev.x) > 1 || util.Abs(cur.y - prev.y) > 1 {
		if cur.x < prev.x {
			cur.x++
		} else if cur.x > prev.x {
			cur.x--
		}
		if cur.y < prev.y {
			cur.y++
		} else if cur.y > prev.y {
			cur.y--
		}
	}
}
