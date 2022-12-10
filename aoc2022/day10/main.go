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

	x := 1
	cycle := 0
	strength := 0
	for _, line := range lines {
		splits := strings.Split(line, " ")
		if splits[0] == "noop" {
			cycle++
			strength += computeStrength(cycle, x)
			crtPrint(cycle, x)
		} else {
			cycle++
			strength += computeStrength(cycle, x)
			crtPrint(cycle, x)
			
			n, err := strconv.Atoi(splits[1])
			if err != nil {
				panic(err)
			}
			cycle++
			strength += computeStrength(cycle, x)
			crtPrint(cycle, x)
			x += n
		}
	}
	fmt.Printf("%v\n", strength)
}

func computeStrength(cycle, x int) int {
	switch cycle {
	case 20, 60, 100, 140, 180, 220:
		return cycle * x
	default:
		return 0
	}
	return 0
}
		
func crtPrint(cycle, x int) {
	pos := ((cycle-1) % 40) + 1
	if pos >= x && pos <= x + 2 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	if pos == 40 {
		fmt.Printf("\n")
	}
}
