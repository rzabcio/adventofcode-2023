package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
)

func main() {
	start := time.Now().UnixNano() / int64(time.Millisecond)

	m := map[string]func(string) int{
		"day1_1": Day01_1, "day1_2": Day01_2,
		"day2_1": Day02_1, "day2_2": Day02_2,
		"day3_1": Day03_1, "day3_2": Day03_2,
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
	ch = make(chan string)
	go func() {
		file, err := os.Open(filename)
		if err != nil {
			close(ch)
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()
	return ch
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
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
