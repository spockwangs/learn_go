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

	for _, line := range lines {
		pos, marker := detectMarker(line, 4)
		fmt.Printf("%v %s\n", pos, marker)
		pos, marker = detectMarker(line, 14)
		fmt.Printf("%v %s\n", pos, marker)
	}
}

func detectMarker(s string, n int) (int, string) {
	marker := make([]rune, 0, n)
	for i, c := range s {
		if len(marker) == n {
			marker = marker[1:]
		}
		marker = append(marker, c)
		if len(marker) == n && isDistinct(marker) {
			return i+1, string(marker)
		}
	}
	return 0, ""
}

func isDistinct(s []rune) bool {
	m := make(map[rune]bool)
	for _, c := range s {
		_, found := m[c]
		if found {
			return false
		}
		m[c] = true
	}
	return true
}
