package day09

import (
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day09_1(filename string) (result int) {
	reads := ReadScanner(filename)
	for _, read := range reads {
		result += predictNextRead(read)
	}
	return result
}

func Day09_2(filename string) (result int) {
	reads := ReadScanner(filename)
	for _, read := range reads {
		result += predictPrevRead(read)
	}
	return result
}

func ReadScanner(filename string) (result [][]int) {
	for line := range utils.InputCh(filename) {
		readLine := make([]int, 0)
		for _, field := range strings.Fields(line) {
			read, _ := strconv.Atoi(field)
			readLine = append(readLine, read)
		}
		result = append(result, readLine)
	}
	return result
}

func diffSequence(sequence []int) (result []int) {
	// shouldn happen, but just in case
	if len(sequence) < 2 {
		panic("diffSequence: sequence too short")
	}
	// base case - if all zeros return empty sequence
	allZeroes := true
	for i := 0; i < len(sequence); i++ {
		if sequence[i] != 0 {
			allZeroes = false
			break
		}
	}
	if allZeroes {
		return
	}
	// general case - build next sequence: a[2]-a[1], a[3]-a[2], ..., a[n]-a[n-1] (so 1 less element)
	for i := 0; i < len(sequence)-1; i++ {
		result = append(result, int(float64(sequence[i+1]-sequence[i])))
	}
	// fmt.Printf("    - %v\n", result)
	return result
}

func predictNextRead(sequence []int) (result int) {
	return predictRead(sequence, "next")
}

func predictPrevRead(sequence []int) (result int) {
	return predictRead(sequence, "prev")
}

func predictRead(sequence []int, t string) (result int) {
	sequences := make([][]int, 0)
	sequences = append(sequences, sequence)
	for {
		sequence = diffSequence(sequence)
		sequences = append(sequences, sequence)
		if len(sequence) == 1 {
			return 0
		}
		if len(sequence) == 0 {
			break
		}
	}
	for i := len(sequences) - 3; i >= 1; i-- {
		if t == "prev" {
			// part 1 subtract from first element
			sequences[i-1] = append(sequences[i-1], sequences[i-1][0]-sequences[i][len(sequences[i])-1])
		} else {
			// part 2 add to last element
			sequences[i-1] = append(sequences[i-1], sequences[i-1][len(sequences[i-1])-1]+sequences[i][len(sequences[i])-1])
		}
	}
	// for _, seq := range sequences {
	// 	fmt.Printf("    + %v\n", seq)
	// }
	result = sequences[0][len(sequences[0])-1]
	return result
}
