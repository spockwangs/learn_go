package main

import (
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"os"
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

	grid := newGrid()
	for _, line := range lines {
		grid.addLine(line)
	}
	fmt.Printf("%v\n", grid.solve(2000000))
	fmt.Printf("%v\n", grid.solve2())
}

type Grid struct {
	measurements []Measurement
}

type Measurement struct {
	sensor, beacon Position
}

func newGrid() *Grid {
	return &Grid{}
}

func (this *Grid) addLine(s string) {
	sensorPos := Position{}
	beaconPos := Position{}
	next := 0
	util.ConsumeStr(s, "Sensor at x=", &next)
	sensorPos.x = util.ConsumeInt(s, &next)
	util.ConsumeStr(s, ", y=", &next)
	sensorPos.y = util.ConsumeInt(s, &next)
	util.ConsumeStr(s, ": closest beacon is at x=", &next)
	beaconPos.x = util.ConsumeInt(s, &next)
	util.ConsumeStr(s, ", y=", &next)
	beaconPos.y = util.ConsumeInt(s, &next)
	this.measurements = append(this.measurements, Measurement{
		sensor: sensorPos,
		beacon: beaconPos,
	})
}

func (this Grid) solve(y int) int {
	multiIntervals := MultiInterval{}
	exclusiveBeacons := make(map[Position]bool)
	for _, measurement := range this.measurements {
		interval := computeCoveredInterval(measurement.sensor, measurement.beacon, y)
		multiIntervals.merge(interval)
		if measurement.beacon.y == y {
			exclusiveBeacons[measurement.beacon] = true
		}
	}
	return multiIntervals.length() - len(exclusiveBeacons)
}

func (this Grid) solve2() int {
	for y := 0; y <= 4000000; y++ {
		multiIntervals := MultiInterval{}
		for _, measurement := range this.measurements {
			interval := computeCoveredInterval(measurement.sensor, measurement.beacon, y)
			multiIntervals.merge(interval)
		}
		if !multiIntervals.contains(newInterval(0, 4000000)) {
			for x := 0; x <= 4000000; x++ {
				found := false
				for _, v := range multiIntervals.intervals {
					if x >= v.first && x <= v.last {
						found = true
					}
				}
				if !found {
					return x*4000000 + y
				}
			}
		}
	}
	return 0
}

func computeCoveredInterval(sensor, beacon Position, y int) Interval {
	dist := computeManhattanDistance(sensor, beacon)
	absY := util.Abs(sensor.y - y)
	if dist < absY {
		return emptyInterval()
	} else {
		absX := dist - absY
		return newInterval(-absX+sensor.x, absX+sensor.x)
	}
}

type Position struct {
	x, y int
}

func newPosition(x, y int) Position {
	return Position{
		x: x,
		y: y,
	}
}

func computeManhattanDistance(p1, p2 Position) int {
	return util.Abs(p1.x - p2.x) + util.Abs(p1.y - p2.y)
}

type MultiInterval struct {
	intervals []Interval
}

func (this MultiInterval) length() int {
	ret := 0
	for _, v := range this.intervals {
		ret += v.length()
	}
	return ret
}

func (this *MultiInterval) merge(interval Interval) {
	result := []Interval{}
	for _, v := range this.intervals {
		ok := interval.merge(v)
		if !ok {
			result = append(result, v)
		}
	}
	result = append(result, interval)
	this.intervals = result
}

func (this MultiInterval) contains(interval Interval) bool {
	for _, v := range this.intervals {
		if v.contains(interval) {
			return true
		}
	}
	return false
}

type Interval struct {
	first, last int
}

func emptyInterval() Interval {
	return Interval{
		first: 1,
		last: 0,
	}
}

func newInterval(a, b int) Interval {
	return Interval{
		first: a,
		last: b,
	}
}
		
func (this Interval) isEmpty() bool {
	if this.first > this.last {
		return true
	}
	return false
}

func (this *Interval) merge(o Interval) bool {
	if o.isEmpty() {
		return true
	}
	if this.last < o.first-1 || this.first > o.last+1 {
		return false
	}
	this.first = util.MinInt(this.first, o.first)
	this.last = util.MaxInt(this.last, o.last)
	return true
}

func (this Interval) length() int {
	if this.isEmpty() {
		return 0
	}
	return this.last - this.first + 1
}

func (this Interval) contains(o Interval) bool {
	if o.isEmpty() {
		return true
	}
	if this.first <= o.first && this.last >= o.last {
		return true
	}
	return false
}
