package util

import (
	"os"
	"bufio"
	"strings"
	"strconv"
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

func ConsumeStr(s, prefix string, next *int) bool {
	if strings.HasPrefix(s[*next:], prefix) {
		*next += len(prefix)
		return true
	}
	return false
}

func ConsumeStrUntil(s, stop string, next *int) string {
	end := *next
	for end < len(s) {
		if end + len(stop) < len(s) {
			if s[end:end+len(stop)] != stop {
				end++
			} else {
				break
			}
		} else {
			end = len(s)
			break
		}
	}
	start := *next
	*next = MinInt(end + len(stop), len(s))
	return s[start:end]
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
