package main

import (
	"os"
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"unicode"
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

	var list1 List
	var list2 List
	type TaggedList struct {
		l List
		tag bool
	}
	lists := []TaggedList{}
	sum := 0
	for i, line := range lines {
		if (i%3) == 0 {
			list1 = parseString(line)
			lists = append(lists, TaggedList{ l: list1 })
		} else if (i%3) == 1 {
			list2 = parseString(line)
			lists = append(lists, TaggedList{ l: list2 })
			if compareLists(list1, list2) < 0 {
				sum += i/3 + 1
			}
		}
	}
	fmt.Printf("%v\n", sum)

	lists = append(lists, TaggedList{
		l: List{
			Element{
				list: List{
					Element{
						a: 2,
					},
				},
			},
		},
		tag: true,
	}, TaggedList{
		l: List{
			Element{
				list: List{
					Element{
						a: 6,
					},
				},
			},
		},
		tag: true,
	})
	sort.Slice(lists, func(i, j int) bool {
		return compareLists(lists[i].l, lists[j].l) < 0
	})
	product := 1
	for i, l := range lists {
		if l.tag {
			product *= i+1
		}
	}
	fmt.Printf("%v\n", product)
}

type List []Element

type Element struct {
	list List
	a int
}

func compareLists(list1, list2 List) int {
	for i := 0; i < len(list1) && i < len(list2); i++ {
		e1 := list1[i]
		e2 := list2[i]
		switch compareElements(e1, e2) {
		case -1:
			return -1
		case 1:
			return 1
		}
	}
	diff := len(list1) - len(list2)
	if diff < 0 {
		return -1
	} else if diff == 0 {
		return 0
	} else {
		return 1
	}
}

func compareElements(e1, e2 Element) int {
	if e1.list != nil && e2.list != nil {
		return compareLists(e1.list, e2.list)
	} else if e1.list == nil && e2.list == nil {
		diff := e1.a - e2.a
		if diff < 0 {
			return -1
		} else if diff == 0 {
			return 0
		} else {
			return 1
		}
	} else if e1.list == nil {
		return compareLists([]Element{
			Element{
				a: e1.a,
			},
		}, e2.list)
	} else if e2.list == nil {
		return compareLists(e1.list, []Element{
			Element{
				a: e2.a,
			},
		})
	}
	return 0
}
	
func parseString(s string) List {
	list, _ := parseList(s, 0)
	return list
}

func parseElement(s string, next int) (Element, int) {
	ret := Element{}
	if s[next] == '[' {
		list, next := parseList(s, next)
		ret.list = list
		return ret, next
	} else {
		a, next := parseInt(s, next)
		ret.a = a
		return ret, next
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
		element, next2 := parseElement(s, next)
		next = next2
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
			
func parseInt(s string, next int) (int, int) {
	start := next
	for unicode.IsDigit(rune(s[next])) {
		next++
	}
	i, err := strconv.Atoi(s[start:next])
	if err != nil {
		panic(err)
	}
	return i, next
}
