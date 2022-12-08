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

	grid := make([][]Tree, 0)
	for _, line := range lines {
		grid = append(grid, make([]Tree, 0))
		lastRow := len(grid)-1
		for _, c := range line {
			if c < '0' || c > '9' {
				panic("")
			}
			grid[lastRow] = append(grid[lastRow], newTree(int(c - '0')))
		}
	}

	numOfRows := len(grid)
	numOfColumns := len(grid[0])

	// Look from left.
	for i := 0; i < numOfRows; i++ {
		maxHeightSeen := -1
		for j := 0; j < numOfColumns; j++ {
			if grid[i][j].height > maxHeightSeen {
				grid[i][j].visible = true
				maxHeightSeen = grid[i][j].height
			}
		}
	}

	// Look from right.
	for i := 0; i < numOfRows; i++ {
		maxHeightSeen := -1
		for j := numOfColumns - 1; j >= 0; j-- {
			if grid[i][j].height > maxHeightSeen {
				grid[i][j].visible = true
				maxHeightSeen = grid[i][j].height
			}
		}
	}

	// Look from top.
	for j := 0; j < numOfColumns; j++ {
		maxHeightSeen := -1
		for i := 0; i < numOfRows; i++ {
			if grid[i][j].height > maxHeightSeen {
				grid[i][j].visible = true
				maxHeightSeen = grid[i][j].height
			}
		}
	}

	// Look from down.
	for j := 0; j < numOfColumns; j++ {
		maxHeightSeen := -1
		for i := numOfRows-1; i >= 0; i-- {
			if grid[i][j].height > maxHeightSeen {
				grid[i][j].visible = true
				maxHeightSeen = grid[i][j].height
			}
		}
	}

	numOfVisible := 0
	for i := 0; i < numOfRows; i++ {
		for j := 0; j < numOfColumns; j++ {
			if grid[i][j].visible {
				numOfVisible++
			}
		}
	}
	fmt.Printf("%v\n", numOfVisible)

	maxScenicScore := 0
	for i := 1; i < numOfRows - 1; i++ {
		for j := 1; j < numOfColumns - 1; j++ {
			score := computeScenicScore(grid, i, j)
			maxScenicScore = util.MaxInt(maxScenicScore, score)
		}
	}
	fmt.Printf("%v\n", maxScenicScore)
}

type Tree struct {
	height int
	visible bool
}

func newTree(n int) Tree {
	return Tree{
		height: n,
		visible: false,
	}
}

func computeScenicScore(grid [][]Tree, i, j int) int {
	if i <= 0 || j <= 0 {
		return 0
	}

	numOfRows := len(grid)
	numOfColumns := len(grid[0])
	myHeight := grid[i][j].height
	leftViewDistance := 0
	rightViewDistance := 0
	topViewDistance := 0
	downViewDistance := 0
	for jj := j-1; jj >= 0; jj-- {
		if grid[i][jj].height < myHeight {
			leftViewDistance++
		} else if grid[i][jj].height == myHeight {
			leftViewDistance++
			break
		} else {
			break
		}
	}

	for jj := j+1; jj < numOfColumns; jj++ {
		thisHeight := grid[i][jj].height
		if thisHeight < myHeight {
			rightViewDistance++
		} else if thisHeight == myHeight {
			rightViewDistance++
			break
		} else {
			break
		}
	}

	for ii := i-1; ii >= 0; ii-- {
		thisHeight := grid[ii][j].height
		if thisHeight < myHeight {
			topViewDistance++
		} else if thisHeight == myHeight {
			topViewDistance++
			break
		} else {
			break
		}
	}
		
	for ii := i+1; ii < numOfRows; ii++ {
		thisHeight := grid[ii][j].height
		if thisHeight < myHeight {
			downViewDistance++
		} else if thisHeight == myHeight {
			downViewDistance++
			break
		} else {
			break
		}
	}
		
	return leftViewDistance*rightViewDistance*topViewDistance*downViewDistance
}
