package main

import (
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"os"
	"sort"
	"strconv"
	"unicode"
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

	var list1 List
	var list2 List
	type TaggedList struct {
		l   List
		tag bool
	}
	lists := []TaggedList{}
	sum := 0
	for i, line := range lines {
		if (i % 3) == 0 {
			list1 = parseString(line)
			lists = append(lists, TaggedList{l: list1})
		} else if (i % 3) == 1 {
			list2 = parseString(line)
			lists = append(lists, TaggedList{l: list2})
			if list1.compare(list2) < 0 {
				sum += i/3 + 1
			}
		}
	}
	fmt.Printf("%v\n", sum)

	lists = append(lists, TaggedList{
		l: List{
			List{
				Int(2),
			},
		},
		tag: true,
	}, TaggedList{
		l: List{
			List{
				Int(6),
			},
		},
		tag: true,
	})
	sort.Slice(lists, func(i, j int) bool {
		return lists[i].l.compare(lists[j].l) < 0
	})
	product := 1
	for i, l := range lists {
		if l.tag {
			product *= i + 1
		}
	}
	fmt.Printf("%v\n", product)
}

type List []Element

type Element interface {
	compare(other Element) int
}

type Int int

func (this List) compare(other Element) int {
	switch v := other.(type) {
	case List:
		for i := 0; i < len(this) && i < len(v); i++ {
			diff := this[i].compare(v[i])
			if diff != 0 {
				return diff
			}
		}
		diff := len(this) - len(v)
		if diff < 0 {
			return -1
		} else if diff == 0 {
			return 0
		} else {
			return 1
		}
	case Int:
		return this.compare(List{v})
	default:
		panic("")
	}
}

func (this Int) compare(other Element) int {
	switch v := other.(type) {
	case List:
		return List{this}.compare(v)
	case Int:
		diff := this - v
		if diff < 0 {
			return -1
		} else if diff == 0 {
			return 0
		} else {
			return 1
		}
	default:
		panic("")
	}
}

func parseString(s string) List {
	list, _ := parseList(s, 0)
	return list
}

func parseElement(s string, next int) (Element, int) {
	if s[next] == '[' {
		return parseList(s, next)
	} else {
		return parseInt(s, next)
	}
}

func parseList(s string, next int) (List, int) {
	if s[next] != '[' {
		panic("")
	}
	ret := List{}
	next++
	for {
		if s[next] == ']' {
			next++
			break
		}
		var element Element
		element, next = parseElement(s, next)
		ret = append(ret, element)
		if s[next] == ',' {
			next++
		} else if s[next] == ']' {
			next++
			break
		} else {
			panic("")
		}
	}
	return ret, next
}

func parseInt(s string, next int) (Int, int) {
	start := next
	for unicode.IsDigit(rune(s[next])) {
		next++
	}
	i, err := strconv.Atoi(s[start:next])
	if err != nil {
		panic(err)
	}
	return Int(i), next
}
