package main

import (
	"os"
	"fmt"
	"strconv"
	"github.com/learn_go/aoc2022/util"
	"strings"
	"github.com/golang-collections/collections/stack"
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

	stacks := make([]*stack.Stack, 0)
	instructionMode := false
	instructions := make([]Instruction, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, "[") {
			for i, r := range line {
				if i < 1 || ((i-1) % 4) != 0 {
					continue
				}
				idx := (i-1)/4
				if len(stacks) <= idx {
					stacks = append(stacks, stack.New())
				}
				if r != ' ' {
					stacks[idx].Push(r)
				}
			}
		} else if len(strings.TrimSpace(line)) == 0 {
			instructionMode = true
		} else if instructionMode {
			splits := strings.Split(line, " ")
			n, err := strconv.Atoi(strings.TrimSpace(splits[1]))
			if err != nil {
				panic(err)
			}
			from, err := strconv.Atoi(strings.TrimSpace(splits[3]))
			if err != nil {
				panic(err)
			}
			to, err := strconv.Atoi(strings.TrimSpace(splits[5]))
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, Instruction{
				n: n,
				from: from,
				to: to,
			})
		}
	}
	// Reverse all stacks
	for _, s := range stacks {
		*s = *reverseStack(s)
	}

	// Part 1
	stacks2 := copyStacks(stacks)
	for _, instruction := range instructions {
		runInstruction(stacks, instruction)
	}
	for _, s := range stacks {
		fmt.Printf("%c", s.Peek())
	}
	fmt.Println("")

	// Part 2
	for _, instruction := range instructions {
		runInstruction2(stacks2, instruction)
	}
	for _, s := range stacks2 {
		fmt.Printf("%c", s.Peek())
	}
	fmt.Println("")
}

type Instruction struct {
	n int
	from int
	to int
}

func reverseStack(s *stack.Stack) *stack.Stack {
	ret := stack.New()
	for {
		if v := s.Pop(); v != nil {
			ret.Push(v)
		} else {
			break
		}
	}
	return ret
}

func runInstruction(stacks []*stack.Stack, ins Instruction) {
	for i := 0; i < ins.n; i += 1 {
		v := stacks[ins.from-1].Pop()
		stacks[ins.to-1].Push(v)
	}
}

func runInstruction2(stacks []*stack.Stack, ins Instruction) {
	s := stack.New()
	for i := 0; i < ins.n; i += 1 {
		v := stacks[ins.from-1].Pop()
		s.Push(v)
	}
	for v := s.Pop(); v != nil; v = s.Pop() {
		stacks[ins.to-1].Push(v)
	}
}

func copyStack(s *stack.Stack) *stack.Stack {
	tmp := stack.New()
	ret := stack.New()
	for v := s.Pop(); v != nil; v = s.Pop() {
		tmp.Push(v)
	}
	for v := tmp.Pop(); v != nil; v = tmp.Pop() {
		s.Push(v)
		ret.Push(v)
	}
	return ret
}

func copyStacks(stacks []*stack.Stack) []*stack.Stack {
	ret := make([]*stack.Stack, 0, len(stacks))
	for _, s := range stacks {
		ret = append(ret, copyStack(s))
	}
	return ret
}
	
