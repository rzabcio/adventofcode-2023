package utils

import (
	"bufio"
	"os"
)

func InputCh(filename string) (ch chan string) {
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

func MinMax(array []int) (int, int) {
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

func Min(array []int) int {
	min, _ := MinMax(array)
	return min
}

func Max(array []int) int {
	_, max := MinMax(array)
	return max
}
