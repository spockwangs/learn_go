package main

import (
	"github.com/learn_go/aoc2022/util"
	"os"
	"fmt"
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

	total_priority := 0
	for _, line := range lines {
		var m = make(map[byte]bool)
		for i := 0; i < len(line)/2; i += 1 {
			for j := len(line)/2; j < len(line); j += 1 {
				if line[i] == line[j] {
					m[line[i]] = true
				}
			}
		}
		for key, _ := range m {
			total_priority += getPriority(key)
		}
	}
	fmt.Printf("%v\n", total_priority)

	total_priority = 0
	for i := 0; i < len(lines); i += 3 {
		line1 := []byte(lines[i])
		sort.Slice(line1, func(i, j int) bool {
			return line1[i] < line1[j]
		})
		line2 := []byte(lines[i+1])
		sort.Slice(line2, func(i, j int) bool {
			return line2[i] < line2[j]
		})
		line3 := []byte(lines[i+2])
		sort.Slice(line3, func(i, j int) bool {
			return line3[i] < line3[j]
		})
		j := 0
		k := 0
		l := 0
		for {
			if line1[j] == line2[k] && line2[k] == line3[l] {
				total_priority += getPriority(line1[j])
				break
			} else {
				m := util.MinByte(line1[j], line2[k], line3[l])
				if line1[j] == m {
					j += 1
				}
				if line2[k] == m {
					k += 1
				}
				if line3[l] == m {
					l += 1
				}
				if j >= len(line1) || k >= len(line2) || l >= len(line3) {
					break
				}
			}
		}
	}
	fmt.Printf("%v\n", total_priority)
		
}

func getPriority(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(b - 'a' + 1)
	} else if b >= 'A' && b <= 'Z' {
		return int(b - 'A' + 27)
	} else {
		panic("")
	}
}
