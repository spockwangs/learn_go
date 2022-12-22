package main

import (
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"os"
	"strconv"
	"strings"
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

	cubes := newCubes()
	for _, line := range lines {
		cubes.addCube(line)
	}
	fmt.Printf("%v\n", cubes.countSides())
	fmt.Printf("%v\n", cubes.countSides2())
}

type Cubes struct {
	cubes map[Position]bool
	maxX, minX, maxY, minY, maxZ, minZ int
}

func newCubes() Cubes {
	return Cubes{
		cubes: make(map[Position]bool),
		minX: kMaxInt,
		maxX: kMinInt,
		minY: kMaxInt,
		maxY: kMinInt,
		minZ: kMaxInt,
		maxZ: kMinInt,
	}
}

func (this *Cubes) addCube(s string) {
	splits := strings.Split(s, ",")
	if len(splits) != 3 {
		panic("")
	}
	x, err := strconv.Atoi(splits[0])
	if err != nil {
		panic("")
	}
	y, err := strconv.Atoi(splits[1])
	if err != nil {
		panic("")
	}
	z, err := strconv.Atoi(splits[2])
	if err != nil {
		panic("")
	}
	this.cubes[newPosition(x, y, z)] = true
	this.minX = util.MinInt(this.minX, x)
	this.maxX = util.MaxInt(this.maxX, x)
	this.minY = util.MinInt(this.minY, y)
	this.maxY = util.MaxInt(this.maxY, y)
	this.minZ = util.MinInt(this.minZ, z)
	this.maxZ = util.MaxInt(this.maxZ, z)
}

func (this Cubes) countSides() int {
	cnt := 0
	for cube := range this.cubes {
		neighbours := []Position{
			newPosition(cube.x, cube.y, cube.z+1),
			newPosition(cube.x, cube.y, cube.z-1),
			newPosition(cube.x-1, cube.y, cube.z),
			newPosition(cube.x+1, cube.y, cube.z),
			newPosition(cube.x, cube.y-1, cube.z),
			newPosition(cube.x, cube.y+1, cube.z),
		}
		for _, n := range neighbours {
			if _, ok := this.cubes[n]; !ok {
				cnt++
			}
		}
	}
	return cnt
}

func (this Cubes) countSides2() int {
	cnt := 0
	queue := []Position{}
	queue = append(queue, newPosition(this.minX-1, this.minY-1, this.minZ-1))
	visited := map[Position]bool{}
	for len(queue) > 0 {
		cube := queue[0]
		queue = queue[1:]
		neighbours := []Position{
			newPosition(cube.x, cube.y, cube.z+1),
			newPosition(cube.x, cube.y, cube.z-1),
			newPosition(cube.x-1, cube.y, cube.z),
			newPosition(cube.x+1, cube.y, cube.z),
			newPosition(cube.x, cube.y-1, cube.z),
			newPosition(cube.x, cube.y+1, cube.z),
		}
		if _, ok := visited[cube]; ok {
			continue
		}
		visited[cube] = true
		for _, n := range neighbours {
			if _, ok := this.cubes[n]; ok {
				cnt++
			} else if n.x >= this.minX-1 && n.x <= this.maxX+1 && n.y >= this.minY-1 &&
				n.y <= this.maxY+1 && n.z >= this.minZ-1 && n.z <= this.maxZ+1 {
				if _, ok := visited[n]; !ok {
					queue = append(queue, n)
				}
			}
		}
	}
	return cnt
}

type Position struct {
	x, y, z int
}

func newPosition(x, y, z int) Position {
	return Position{
		x: x,
		y: y,
		z: z,
	}
}

const (
	kMaxInt = int(^uint(0)>>1)
	kMinInt = -kMaxInt-1
)
