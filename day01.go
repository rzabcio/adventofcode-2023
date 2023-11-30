package main

import (
	"fmt"
	"strings"
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
	return result
}

var txtNums = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

var intNums = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

func Day01_2(filename string) (result int) {
	for line := range inputCh(filename) {
		nums := make([]int, 0, len(line))
		for i := 0; i < len(line); i++ {
			nums = append(nums, -1)
		}
		for num, txtNum := range txtNums {
			if index := strings.Index(line, txtNum); index != -1 {
				nums[index] = num
			}
			if index := strings.LastIndex(line, txtNum); index != -1 {
				nums[index] = num
			}
		}
		for num, intNum := range intNums {
			if index := strings.Index(line, string(intNum)); index != -1 {
				nums[index] = num
			}
			if index := strings.LastIndex(line, string(intNum)); index != -1 {
				nums[index] = num
			}
		}
		firstNum := 0
		secondNum := 0
		for _, num := range nums {
			if num > -1 {
				if firstNum == 0 {
					firstNum = num
				}
				secondNum = num
			}
		}
		subResult := firstNum*10 + secondNum
		fmt.Printf("%s -> %v -> %d\n", line, nums, subResult)
		result += subResult
	}
	return result
}
