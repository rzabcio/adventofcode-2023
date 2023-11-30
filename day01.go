package main

import (
	"unicode"
)

func Day01_1(filename string) (result int) {
	for line := range inputCh(filename) {
		var num int
		for _, r := range line {
			if unicode.IsDigit(r) {
				if num == 0 {
					result += int(r-'0') * 10
				}
				num = int(r - '0')
			}
		}
		result += num
	}
	// result = 142
	return result
}

func Day01_2(filename string) (result int) {
	return result
}
