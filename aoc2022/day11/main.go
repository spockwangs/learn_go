package main

import (
	"os"
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"strings"
	"strconv"
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

	monkeys := make([]Monkey, 0)
	var monkey *Monkey
	modulo := 1
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey") {
			monkey = new(Monkey)
		} else if strings.HasPrefix(line, "  Starting items: ") {
			line = strings.TrimPrefix(line, "  Starting items: ")
			splits := strings.Split(line, ",")
			for _, s := range splits {
				n, err := strconv.Atoi(strings.TrimSpace(s))
				if err != nil {
					panic(err)
				}
				monkey.addItem(int(n))
			}
		} else if strings.HasPrefix(line, "  Operation: new = old ") {
			line = strings.TrimPrefix(line, "  Operation: new = old ")
			if strings.HasPrefix(line, "*") {
				line = strings.TrimPrefix(line, "* ")
				if line == "old" {
					monkey.operation = func(a int) int {
						return a*a
					}
				} else {
					n, err := strconv.Atoi(line)
					if err != nil {
						panic(err)
					}
					monkey.operation = func(a int) int {
						return a * int(n)
					}
				}
			} else if strings.HasPrefix(line, "+") {
				n, err := strconv.Atoi(strings.TrimPrefix(line, "+ "))
				if err != nil {
					panic(err)
				}
				monkey.operation = func(a int) int {
					return a + int(n)
				}
			}
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			line = strings.TrimPrefix(line, "  Test: divisible by ")
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			monkey.test = func(i int) bool {
				return (i % int(n)) == 0
			}
			modulo *= n
		} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
			line = strings.TrimPrefix(line, "    If true: throw to monkey ")
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			monkey.trueBranch = n
		} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
			line = strings.TrimPrefix(line, "    If false: throw to monkey ")
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			monkey.falseBranch = n
		} else if len(strings.TrimSpace(line)) == 0 {
			monkeys = append(monkeys, *monkey)
			monkey = nil
		}
	}
	if monkey != nil {
		monkeys = append(monkeys, *monkey)
	}

	for i := 0; i < 10000; i++ {
		for j := range monkeys {
			throws := monkeys[j].inspectItems(modulo)
			for _, t := range throws {
				monkeys[t.monkey].addItem(t.worryLevel)
			}
		}
	}
	for _, m := range monkeys {
		fmt.Printf("%v\n", m.inspectTimes)
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectTimes > monkeys[j].inspectTimes
	})
	fmt.Printf("%v\n", monkeys[0].inspectTimes * monkeys[1].inspectTimes)
}

type Monkey struct {
	items []int
	operation func(int) int
	test func(int) bool
	trueBranch int
	falseBranch int
	inspectTimes int
}

type Throw struct {
	monkey int
	worryLevel int
}

func (m *Monkey) inspectItems(modulo int) []Throw {
	ret := make([]Throw, 0)
	for _, item := range m.items {
		throw := m.inspectItem(item, modulo)
		ret = append(ret, throw)
	}
	m.items = make([]int, 0)
	return ret
}
	
func (m *Monkey) inspectItem(item int, modulo int) Throw {
	m.inspectTimes++
	worryLevel := m.operation(item)
	worryLevel = worryLevel % modulo
	if m.test(worryLevel) {
		return Throw{
			monkey: m.trueBranch,
			worryLevel: worryLevel,
		}
	} else {
		return Throw{
			monkey: m.falseBranch,
			worryLevel: worryLevel,
		}
	}
}

func (m *Monkey) addItem(item int) {
	m.items = append(m.items, item)
}
