package main

import (
	"github.com/learn_go/aoc2022/util"
	"os"
	"fmt"
	"strings"
)

type Hand int

const (
	Rock Hand = iota
	Paper
	Scissors
)

type Pair struct {
	p1 Hand
	p2 Hand
}

var SITUATIONS = []Pair{
	Pair{ p1: Rock, p2: Scissors },
	Pair{ p1: Scissors, p2: Paper },
	Pair{ p1: Paper, p2: Rock },
}

func (h Hand) score() int {
	switch h {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	panic("bad hand")
}

func playAgainst(h1, h2 Hand) int {
	if h1 == h2 {
		return 0
	}

	for _, situation := range SITUATIONS {
		if h1 == situation.p1 && h2 == situation.p2 {
			return 1
		} else if h1 == situation.p2 && h2 == situation.p1 {
			return -1
		}
	}
	panic("")
}

func shouldLose(h Hand) Hand {
	for _, situation := range SITUATIONS {
		if h == situation.p1 {
			return situation.p2
		}
	}
	panic("")
}

func shouldWin(h Hand) Hand {
	for _, situation := range SITUATIONS {
		if h == situation.p2 {
			return situation.p1
		}
	}
	panic("")
}

func shouldDraw(h Hand) Hand {
	return h
}

func strToHand(s string) Hand {
	if s == "A" {
		return Rock
	} else if s == "B" {
		return Paper
	} else if s == "C" {
		return Scissors
	} else {
		panic(fmt.Sprintf("bad char: %v", s))
	}
}

type Game struct {
	player1 Hand
	player2 string
}

func (g Game) player2Hand() Hand {
	if g.player2 == "X" {
		return Rock
	} else if g.player2 == "Y" {
		return Paper
	} else if g.player2 == "Z" {
		return Scissors
	} else {
		panic("bad player2")
	}
}

func (g Game) play(s string) int {
	var hand2 Hand
	if s == "part1" {
		hand2 = g.player2Hand()
	} else {
		hand2 = g.player2Hand2()
	}
	score := hand2.score()

	switch playAgainst(hand2, g.player1) {
	case 1:
		score += 6
	case 0:
		score += 3
	}
	return score
}

func (g Game) player2Hand2() Hand {
	if g.player2 == "X" {
		return shouldLose(g.player1)
	} else if g.player2 == "Y" {
		return shouldDraw(g.player1)
	} else if g.player2 == "Z" {
		return shouldWin(g.player1)
	} else {
		panic("")
	}
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

	var games []Game
	for i, line := range lines {
		splits := strings.Split(line, " ")
		if len(splits) != 2 {
			panic(fmt.Sprintf("bad line: %v %v", i+1, line))
		}
		games = append(games, Game{
			player1: strToHand(splits[0]),
			player2: splits[1],
		})
	}

	score := 0
	for _, game := range games {
		score += game.play("part1")
	}
	fmt.Printf("%v\n", score)
	score = 0
	for _, game := range games {
		score += game.play("part2")
	}
	fmt.Printf("%v\n", score)
}
