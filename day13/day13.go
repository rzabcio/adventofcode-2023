package day13

import (
	"fmt"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day13_1(filename string) (result int) {
	reflectionPatterns := readReflectionPatterns(filename)
	for _, rp := range reflectionPatterns {
		rp.findVertAxis()
		rp.findHorizAxis()
		result += rp.vertAxis + rp.horizAxis*100
	}
	return result
}

func Day13_2(filename string) (result int) {
	return result
}

func readReflectionPatterns(filename string) (result []ReflectionPattern) {
	rp := ReflectionPattern{}

	row := 0
	for line := range utils.InputCh(filename) {
		if strings.HasPrefix(line, "--- ") {
			continue
		}
		// if empty line - add and new pattern
		if line == "" {
			result = append(result, rp)
			rp = ReflectionPattern{}
			row = 0
			continue
		}
		if rp.colCount == nil {
			rp.colCount = make([]int, len(line))
		}
		rp.rowCount = append(rp.rowCount, 0)
		pattern := []int{}
		for col, c := range line {
			if c == '#' {
				pattern = append(pattern, 1)
				rp.colCount[col]++
				rp.rowCount[row]++
			} else {
				pattern = append(pattern, 0)
			}
		}
		rp.pattern = append(rp.pattern, pattern)
		row++
	}
	result = append(result, rp)
	return result
}

type ReflectionPattern struct {
	pattern   [][]int
	colCount  []int
	rowCount  []int
	vertAxis  int
	horizAxis int
}

func (rp ReflectionPattern) String() string {
	result := fmt.Sprintf("vertAxis: %d, horizAxis: %d\n", rp.vertAxis, rp.horizAxis)
	result += "    "
	for i, col := range rp.colCount {
		if i > 1 && i == rp.vertAxis-1 {
			result += fmt.Sprintf(" %2d ", col)
		} else if i > 1 && i == rp.vertAxis {
			result += fmt.Sprintf("|%2d ", col)
		} else {
			result += fmt.Sprintf(" %2d ", col)
		}
	}
	result += "\n"
	for row, pattern := range rp.pattern {
		if row == rp.horizAxis-1 {
			result += fmt.Sprintf(" %2d_", rp.rowCount[row])
		} else if row == rp.horizAxis {
			result += fmt.Sprintf(" %2d ", rp.rowCount[row])
		} else {
			result += fmt.Sprintf(" %2d ", rp.rowCount[row])
		}
		for _, col := range pattern {
			if col == 1 {
				result += "  # "
			} else {
				result += "  . "
			}
		}
		result += "\n"
	}
	return result
}

func (rp *ReflectionPattern) findVertAxis() {
	possibleVertAxis := []int{}
NEXT_AXIS:
	for col := 0; col < len(rp.colCount)-1; col++ {
		for i := 0; ; i++ {
			if col-i < 0 || col+i+1 >= len(rp.colCount) {
				break
			}
			if rp.colCount[col-i] != rp.colCount[col+i+1] {
				continue NEXT_AXIS
			}
			for row := 0; row < len(rp.rowCount); row++ {
				if rp.pattern[row][col-i] != rp.pattern[row][col+i+1] {
					continue NEXT_AXIS
				}
			}
		}
		possibleVertAxis = append(possibleVertAxis, col)
	}
	if len(possibleVertAxis) > 0 {
		rp.vertAxis = possibleVertAxis[0] + 1
	}
	if len(possibleVertAxis) > 1 {
		fmt.Printf("!!! error - there is more than one possible vertical axis: %d\n", possibleVertAxis)
		fmt.Println(rp)
	}
}

func (rp *ReflectionPattern) findHorizAxis() {
	possibleHorizAxis := []int{}
NEXT_AXIS:
	for row := 0; row < len(rp.rowCount)-1; row++ {
		for i := 0; ; i++ {
			if row-i < 0 || row+i+1 >= len(rp.rowCount) {
				break
			}
			if rp.rowCount[row-i] != rp.rowCount[row+i+1] {
				continue NEXT_AXIS
			}
			for col := 0; col < len(rp.colCount); col++ {
				if rp.pattern[row-i][col] != rp.pattern[row+i+1][col] {
					continue NEXT_AXIS
				}
			}
		}
		possibleHorizAxis = append(possibleHorizAxis, row)
	}
	if len(possibleHorizAxis) > 0 {
		rp.horizAxis = possibleHorizAxis[0] + 1
	}
	if len(possibleHorizAxis) > 1 {
		fmt.Printf("!!! error - there is more than one possible horizontal axis: %d\n", possibleHorizAxis)
		fmt.Println(rp)
	}
}
