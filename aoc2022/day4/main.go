package main

import (
	"os"
	"fmt"
	"strconv"
	"github.com/learn_go/aoc2022/util"
	"strings"
)

type Range struct {
	first int
	last int
}

type Assignment struct {
	r1, r2 Range
}

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

	var assignments []Assignment
	for _, line := range lines {
		splits := strings.Split(line, ",")
		if len(splits) != 2 {
			panic("")
		}
		assignments = append(assignments, Assignment{
			r1: parseRange(splits[0]),
			r2: parseRange(splits[1]),
		})
	}

	num := 0
	for _, a := range assignments {
		if fullyContain(a.r1, a.r2) {
			num += 1
		}
	}
	fmt.Printf("%v\n", num)

	num = 0
	for _, a := range assignments {
		if overlap(a.r1, a.r2) {
			num += 1
		}
	}
	fmt.Printf("%v\n", num)
}

func parseRange(s string) Range {
	splits := strings.Split(s, "-")
	if len(splits) != 2 {
		panic("")
	}
	first, err := strconv.Atoi(splits[0])
	if err != nil {
		panic(err)
	}
	last, err := strconv.Atoi(splits[1])
	if err != nil {
		panic(err)
	}
	return Range{
		first: first,
		last: last,
	}
}

func fullyContain(r1, r2 Range) bool {
	if r1.first <= r2.first && r1.last >= r2.last {
		return true
	} else if r2.first <= r1.first && r2.last >= r1.last {
		return true
	} else {
		return false
	}
}

func overlap(r1, r2 Range) bool {
	if r1.last < r2.first || r1.first > r2.last {
		return false
	} else {
		return true
	}
}
