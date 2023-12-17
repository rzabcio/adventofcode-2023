package day13

import (
	"fmt"
	"math"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day13_1(filename string) (result int) {
	reflectionPatterns := readReflectionPatterns(filename)
	for _, rp := range reflectionPatterns {
		rp.findHorizAxis()
		rp.findVertAxis()
		// fmt.Println(rp)
		result += rp.vertAxis + rp.horizAxis*100
	}
	return result
}

func Day13_2(filename string) (result int) {
	reflectionPatterns := readReflectionPatterns(filename)
	for _, rp := range reflectionPatterns {
		rp.findHorizAxis()
		rp.findVertAxis()
		// fmt.Printf("===== START %d =====\n%v\n", i, rp)
		oldVertAxis, oldHorizAxis := rp.vertAxis, rp.horizAxis
		rp.findAndFixSmudge()
		rp.count()
		// fmt.Printf(" oldVertAxis: %d, oldHorizAxis: %d => newVertAxis: %d, newHorizAxis: %d\n", oldVertAxis, oldHorizAxis, rp.vertAxis, rp.horizAxis)
		if oldVertAxis != rp.vertAxis {
			rp.horizAxis = 0
		} else if oldHorizAxis != rp.horizAxis {
			rp.vertAxis = 0
		}
		// fmt.Println(rp)
		result += rp.vertAxis + rp.horizAxis*100
	}
	return result
}

func readReflectionPatterns(filename string) (result []ReflectionPattern) {
	rp := ReflectionPattern{}

	row := 0
	for line := range utils.InputCh(filename) {
		if strings.HasPrefix(line, "- ") {
			continue
		}
		// if empty line - add and new pattern
		if line == "" {
			result = append(result, rp)
			rp = ReflectionPattern{}
			row = 0
			continue
		}
		pattern := []int{}
		for _, c := range line {
			if c == '#' {
				pattern = append(pattern, 1)
			} else {
				pattern = append(pattern, 0)
			}
		}
		rp.pattern = append(rp.pattern, pattern)
		rp.count()
		row++
	}
	result = append(result, rp)
	return result
}

func (rp *ReflectionPattern) count() {
	rp.colCount = make([]int, len(rp.pattern[0]))
	rp.rowCount = make([]int, len(rp.pattern))
	for row, pattern := range rp.pattern {
		for col, c := range pattern {
			if c == 1 {
				rp.colCount[col]++
				rp.rowCount[row]++
			}
		}
	}
}

type ReflectionPattern struct {
	pattern   [][]int
	colCount  []int
	rowCount  []int
	vertAxis  int
	horizAxis int
	rowDiffs  [][]int
	colDiffs  [][]int
}

func (rp ReflectionPattern) String() string {
	result := fmt.Sprintf("--- vertAxis: %d, horizAxis: %d\n", rp.vertAxis, rp.horizAxis)
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

	if rp.rowDiffs != nil {
		result += "\n"
		for row, diffs := range rp.rowDiffs {
			result += fmt.Sprintf("row %d: %v\n", row, diffs)
		}
	}
	if rp.colDiffs != nil {
		for col, diffs := range rp.colDiffs {
			result += fmt.Sprintf("col %d: %v\n", col, diffs)
		}
	}

	return result
}

func (rp *ReflectionPattern) findVertAxis() bool {
	possibleVertAxis := []int{}
	for col := 0; col < len(rp.colCount)-1; col++ {
		if rp.isColAxis(col) {
			possibleVertAxis = append(possibleVertAxis, col)
		}
	}
	if len(possibleVertAxis) > 0 {
		rp.vertAxis = possibleVertAxis[0] + 1
		return true
	}
	if len(possibleVertAxis) > 1 {
		fmt.Printf("!!! error - there is more than one possible vertical axis: %d\n", possibleVertAxis)
		fmt.Println(rp)
	}
	return false
}

func (rp *ReflectionPattern) findHorizAxis() bool {
	possibleHorizAxis := []int{}
	for row := 0; row < len(rp.rowCount)-1; row++ {
		if rp.isRowAxis(row) {
			possibleHorizAxis = append(possibleHorizAxis, row)
		}
	}
	if len(possibleHorizAxis) > 0 {
		rp.horizAxis = possibleHorizAxis[0] + 1
		return true
	}
	if len(possibleHorizAxis) > 1 {
		fmt.Printf("!!! error - there is more than one possible horizontal axis: %d\n", possibleHorizAxis)
		fmt.Println(rp)
	}
	return false
}

func (rp *ReflectionPattern) isColAxis(col int) bool {
	return rp.isColAxisDiffNotGreaterThan(col, 0) == 0
}

func (rp *ReflectionPattern) isColAxisDiffNotGreaterThan(col, maxDiff int) (colDiff int) {
	for i := 0; ; i++ {
		if col-i < 0 || col+i+1 >= len(rp.colCount) {
			break
		}
		colDiff += rp.findColsDiff(col-i, col+i+1)
		if colDiff > maxDiff {
			return colDiff
		}
	}
	return colDiff
}

func (rp *ReflectionPattern) findColsDiffs(col1, col2 int) (result []int) {
	for row := 0; row < len(rp.rowCount); row++ {
		result = append(result, rp.pattern[row][col1]-rp.pattern[row][col2])
	}
	return result
}

func (rp *ReflectionPattern) findColsDiff(col1, col2 int) (result int) {
	for _, row := range rp.findColsDiffs(col1, col2) {
		result += int(math.Abs(float64(row)))
	}
	return result
}

func (rp *ReflectionPattern) isRowAxis(row int) bool {
	return rp.isRowAxisDiffNotGreaterThan(row, 0) == 0
}

func (rp *ReflectionPattern) isRowAxisDiffNotGreaterThan(row, maxDiff int) (rowDiff int) {
	for i := 0; ; i++ {
		if row-i < 0 || row+i+1 >= len(rp.rowCount) {
			break
		}
		rowDiff += rp.findRowsDiff(row-i, row+i+1)
		if rowDiff > maxDiff {
			return rowDiff
		}
	}
	return rowDiff
}

func (rp *ReflectionPattern) findRowsDiffs(row1, row2 int) (result []int) {
	for col := 0; col < len(rp.colCount); col++ {
		result = append(result, rp.pattern[row1][col]-rp.pattern[row2][col])
	}
	return result
}

func (rp *ReflectionPattern) findRowsDiff(row1, row2 int) (result int) {
	for _, col := range rp.findRowsDiffs(row1, row2) {
		result += int(math.Abs(float64(col)))
	}
	return result
}

func (rp *ReflectionPattern) findAndFixSmudge() {
	for row := 0; row < len(rp.rowCount)-1; row++ {
		rowDiff := 0
		ii := 0
		for i := 0; ; i++ {
			if row-i < 0 || row+i+1 >= len(rp.rowCount) {
				break
			}
			d := rp.findRowsDiff(row-i, row+i+1)
			rowDiff += d
			if d == 1 {
				ii = i
			}
		}
		if rowDiff == 1 {
			rp.fixSmudgeInRow(row-ii, row+ii+1)
			rp.horizAxis = row + 1
			return
		}
	}
	for col := 0; col < len(rp.colCount)-1; col++ {
		colDiff := 0
		ii := 0
		for i := 0; ; i++ {
			if col-i < 0 || col+i+1 >= len(rp.colCount) {
				break
			}
			d := rp.findColsDiff(col-i, col+i+1)
			colDiff += d
			if d == 1 {
				ii = i
			}
		}
		if colDiff == 1 {
			rp.fixSmudgeInCol(col-ii, col+ii+1)
			rp.vertAxis = col + 1
			return
		}
	}
}

func (rp *ReflectionPattern) fixSmudgeInRow(row1, row2 int) {
	if rp.findRowsDiff(row1, row2) != 1 {
		fmt.Printf("!!! error - fixing smudge between rows %d, %d, but diff is not 1\n", row1, row2)
		return
	}
	for i, diff := range rp.findRowsDiffs(row1, row2) {
		if diff == 1 {
			rp.pattern[row2][i] = 1
		}
		if diff == -1 {
			rp.pattern[row1][i] = 1
		}
	}
}

func (rp *ReflectionPattern) fixSmudgeInCol(col1, col2 int) {
	if rp.findColsDiff(col1, col2) != 1 {
		fmt.Printf("!!! error - fixing smudge between cols %d, %d, but diff is not 1\n", col1, col2)
		return
	}
	for i, diff := range rp.findColsDiffs(col1, col2) {
		if diff == 1 {
			rp.pattern[i][col2] = 1
		}
		if diff == -1 {
			rp.pattern[i][col1] = 1
		}
	}
}
