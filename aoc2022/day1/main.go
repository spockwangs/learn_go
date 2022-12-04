package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	food []int
}

func (e Elf) sum() int {
	total := 0
	for _, n := range e.food {
		total += n
	}
	return total
}

func (e *Elf) addFood(n int) {
	e.food = append(e.food, n)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v filename\n", os.Args[0])
		return
	}

	lines, err := readLines(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	var elves []Elf
	var curElf Elf
	for _, line := range lines {
		if len(line) == 0 {
			elves = append(elves, curElf)
			curElf = Elf{}
		} else {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			curElf.addFood(n)
		}
	}
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].sum() > elves[j].sum()
	})
	fmt.Printf("%v\n", elves[0].sum())
	fmt.Printf("%v\n", elves[0].sum() + elves[1].sum() + elves[2].sum())
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 10)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
	
