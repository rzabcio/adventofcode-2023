package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"

	"github.com/rzabcio/adventofcode-2023/day04"
	"github.com/rzabcio/adventofcode-2023/day05"
	"github.com/rzabcio/adventofcode-2023/day06"
	"github.com/rzabcio/adventofcode-2023/day07"
	"github.com/rzabcio/adventofcode-2023/day08"
	"github.com/rzabcio/adventofcode-2023/day09"
	"github.com/rzabcio/adventofcode-2023/day11"
	"github.com/rzabcio/adventofcode-2023/day12"
	"github.com/rzabcio/adventofcode-2023/day13"
	"github.com/rzabcio/adventofcode-2023/utils"
)

func main() {
	start := time.Now().UnixNano() / int64(time.Millisecond)

	m := map[string]func(string) int{
		"day1_1": Day01_1, "day1_2": Day01_2,
		"day2_1": Day02_1, "day2_2": Day02_2,
		"day3_1": Day03_1, "day3_2": Day03_2,
		"day4_1": day04.Day04_1, "day4_2": day04.Day04_2,
		"day5_1": day05.Day05_1, "day5_2": day05.Day05_2,
		"day6_1": day06.Day06_1, "day6_2": day06.Day06_2,
		"day7_1": day07.Day07_1, "day7_2": day07.Day07_2,
		"day8_1": day08.Day08_1, "day8_2": day08.Day08_2,
		"day9_1": day09.Day09_1, "day9_2": day09.Day09_2,
		"day11_1": day11.Day11_1, "day11_2": day11.Day11_2,
		"day12_1": day12.Day12_1, "day12_2": day12.Day12_2,
		"day13_1": day13.Day13_1, "day13_2": day13.Day13_2,
	}

	day := &cobra.Command{
		Use:  "day [day_no] [test_no] [filename]",
		Args: cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			f := m["day"+args[0]+"_"+args[1]]
			fmt.Println(f(args[2]))
		},
	}

	rootCmd := &cobra.Command{Use: "app"}
	rootCmd.AddCommand(day)
	rootCmd.Execute()
	fmt.Printf("[time %dms]\n", time.Now().UnixNano()/int64(time.Millisecond)-start)
}

// TOOLS //////////////////////////////////////////////////////////////////////
func inputSl(filename string) []string {
	sl := make([]string, 0)
	for s := range inputCh(filename) {
		sl = append(sl, s)
	}
	return sl
}

func inputSlInt(filename string) []int {
	sl := make([]int, 0)
	for s := range inputChInt(filename) {
		sl = append(sl, s)
	}
	return sl
}

func inputCh(filename string) (ch chan string) {
	return utils.InputCh(filename)
}

func inputChInt(filename string) (ch chan int) {
	ch = make(chan int)
	go func() {
		for str := range inputCh(filename) {
			i, _ := strconv.Atoi(str)
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// TOOLS - STRING
func reverseStr(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

// TOOLS - ARRAYS
func remove(s []string, e string) []string {
	i := funk.IndexOf(s, e)
	if i < 0 {
		return s
	}
	res := make([]string, 0)
	if i == 0 {
		res = s[i+1:]
	} else if i == len(s)-1 {
		res = s[:i]
	} else {
		res = append(s[:i], s[i+1:]...)
	}
	return res
}

func removeInt(s []int, e int) []int {
	i := funk.IndexOf(s, e)
	if i < 0 {
		return s
	}
	res := make([]int, 0)
	if i == 0 {
		res = s[i+1:]
	} else if i == len(s)-1 {
		res = s[:i]
	} else {
		res = append(s[:i], s[i+1:]...)
	}
	return res
}

func contains(s []string, e string) bool {
	return funk.IndexOf(s, e) >= 0
}

func containsInt(s []int, e int) bool {
	return indexOfInt(s, e) >= 0
}

func indexOfInt(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func indexOfRune(s string, e rune) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func lastIndexOfRune(s string, e rune) (pos int) {
	pos = -1
	for i, a := range s {
		if a == e {
			pos = i
		}
	}
	return pos
}

func reverseStrArr(ss []string) []string {
	for i := 0; i < len(ss)/2; i++ {
		j := len(ss) - i - 1
		ss[i], ss[j] = ss[j], ss[i]
	}
	return ss
}

// TOOLS - NUMERICAL
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minMax(array []int) (int, int) {
	return utils.MinMax(array)
}
