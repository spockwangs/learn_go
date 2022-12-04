package util

import (
	"os"
	"bufio"
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
