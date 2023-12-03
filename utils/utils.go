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
