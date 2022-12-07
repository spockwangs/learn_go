package main

import (
	"os"
	"fmt"
	"github.com/learn_go/aoc2022/util"
	"strings"
	"strconv"
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

	tree := newTree()
	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd ") {
			dir := strings.TrimPrefix(line, "$ cd ")
			if dir == "/" {
				tree.reset()
			} else if dir == ".." {
				if !tree.goToParent() {
					panic(fmt.Sprintf("can't go to parent"))
				}
			} else {
				if !tree.goToChild(dir) {
					panic(fmt.Sprintf("dir %s does not exist", dir))
				}
			}
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else {
			splits := strings.Split(line, " ")
			if len(splits) != 2 {
				panic("")
			}
			if splits[0] == "dir" {
				tree.addChildDir(splits[1])
			} else {
				n, err := strconv.Atoi(splits[0])
				if err != nil {
					panic(err)
				}
				tree.addChildFile(splits[1], n)
			}
		}
	}
	used_space := 0
	total_size := 0
	tree.dfs(func (entry Entry, size int) {
		if entry.isDir && size <= 100_000 {
			fmt.Printf("%v %v\n", entry.name, size)
			total_size += size
		}
		if entry.name == "/" {
			used_space = size
		}
	})
	fmt.Printf("total size %v\n", total_size)

	disk_space := 70000000
	available_space := disk_space - used_space
	required_space := 30000000
	min_free_space := required_space - available_space
	var smallest_size int = int(^uint(0) >> 1)
	tree.dfs(func (entry Entry, size int) {
		if entry.isDir && size >= min_free_space {
			if smallest_size > size {
				smallest_size = size
			}
		}
	})
	fmt.Printf("%v\n", smallest_size)
}

type Tree struct {
	root *Node
	cur *Node
}

type Node struct {
	entry Entry
	children []*Node
	parent *Node
}

type Entry struct {
	name string
	size int
	isDir bool
}

func newTree() *Tree {
	root := &Node{
		entry: Entry{
			name: "/",
			size: 0,
			isDir: true,
		},
		children: nil,
		parent: nil,
	}
	return &Tree{
		root: root,
		cur: root,
	}
}

func (t *Tree) reset() {
	t.cur = t.root
}

func (t *Tree) goToParent() bool {
	if t.cur != nil && t.cur.parent != nil {
		t.cur = t.cur.parent
		return true
	}
	return false
}

func (t *Tree) goToChild(name string) bool {
	for _, child := range t.cur.children {
		if child.entry.name == name {
			t.cur = child
			return true
		}
	}
	return false
}

func (t *Tree) addChildFile(name string, size int) {
	t.cur.children = append(t.cur.children, &Node{
		entry: Entry{
			name: name,
			size: size,
			isDir: false,
		},
		children: nil,
		parent: t.cur,
	})
}

func (t *Tree) addChildDir(name string) {
	t.cur.children = append(t.cur.children, &Node{
		entry: Entry{
			name: name,
			size: 0,
			isDir: true,
		},
		children: nil,
		parent: t.cur,
	})
}

func (t *Tree) dfs(visit func(Entry, int)) {
	dfsHelper(t.root, visit)
}

func dfsHelper(node *Node, visit func(Entry, int)) int {
	if !node.entry.isDir {
		visit(node.entry, node.entry.size)
		return node.entry.size
	}

	sum := 0
	for _, child := range node.children {
		sum += dfsHelper(child, visit)
	}
	visit(node.entry, sum)
	return sum
}
