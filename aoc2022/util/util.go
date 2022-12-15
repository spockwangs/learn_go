package util

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"unicode"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
	
func MinByte(a ...byte) byte {
	min := byte(0xFF)
	for _, i := range a {
		if min > i {
			min = i
		}
	}
	return min
}

func MinInt(a ...int) int {
	min := int(^uint(0) >> 1)
	for _, i := range a {
		if min > i {
			min = i
		}
	}
	return min
}

func MaxInt(a ...int) int {
	var max int
	for i, e := range a {
		if i == 0 {
			max = e
		} else if max < e {
			max = e
		}
	}
	return max
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ConsumeStr(s, prefix string, next *int) {
	if !strings.HasPrefix(s[*next:], prefix) {
		panic(fmt.Sprintf("`%v` does not has prefix `%v`", s[*next:], prefix))
	}
	*next += len(prefix)
}

func ConsumeInt(s string, next *int) int {
	end := *next
	for end < len(s) && (unicode.IsDigit(rune(s[end])) || s[end] == '-') {
		end++
	}
	n, err := strconv.Atoi(s[*next:end])
	if err != nil {
		panic(err)
	}
	*next = end
	return n
}
