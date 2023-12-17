package day14

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day14_1(filename string) (result int) {
	rd := NewReflectorDish(filename)
	rd.RotateClockwise()
	rd.TiltEast()
	rd.RotateCounterClockwise()
	return rd.TotalLoad()
}

func Day14_2(filename string) (result int) {
	rd := NewReflectorDish(filename)
	for i := 0; i < 1000000000; i++ {
		if i%10000000 == 0 {
			fmt.Printf("i: %d\n", i)
		}
		rd.Cycle()
	}
	return rd.TotalLoad()
}

type ReflectorDish struct {
	original []string
	stones   []string
}

func NewReflectorDish(filename string) *ReflectorDish {
	rd := ReflectorDish{original: []string{}}
	for line := range utils.InputRows(filename) {
		rd.original = append(rd.original, line)
		rd.stones = append(rd.stones, line)
	}
	return &rd
}

func (rd *ReflectorDish) String() string {
	result := "===\n"
	for _, line := range rd.stones {
		result += line + "\n"
	}
	result += fmt.Sprintf("=== total load: %d\n", rd.TotalLoad())
	return result
}

func (rd *ReflectorDish) TotalLoad() (totalLoad int) {
	rowLoad := len(rd.stones)
	loadChange := -1
	for _, line := range rd.stones {
		totalLoad += strings.Count(line, "O") * rowLoad
		rowLoad += loadChange
	}
	return totalLoad
}

func (rd *ReflectorDish) TiltEast() {
	newStones := []string{}
	for _, line := range rd.stones {
		parts := strings.Split(line, "#")
		newParts := []string{}
		for _, part := range parts {
			// fmt.Printf("   '%s'", part)
			partsSplit := strings.Split(part, "")
			sort.Strings(partsSplit)
			part = strings.Join(partsSplit, "")
			// fmt.Printf(" => '%s'\n", part)
			newParts = append(newParts, part)
		}
		newStones = append(newStones, strings.Join(newParts, "#"))
	}
	rd.stones = newStones
}

func (rd *ReflectorDish) RotateClockwise() {
	newStones := []string{}
	for i := 0; i < len(rd.stones[0]); i++ {
		newLine := ""
		for j := len(rd.stones) - 1; j >= 0; j-- {
			newLine += string(rd.stones[j][i])
		}
		newStones = append(newStones, newLine)
	}
	rd.stones = newStones
}

func (rd *ReflectorDish) RotateCounterClockwise() {
	newStones := []string{}
	for i := len(rd.stones[0]) - 1; i >= 0; i-- {
		newLine := ""
		for j := 0; j < len(rd.stones); j++ {
			newLine += string(rd.stones[j][i])
		}
		newStones = append(newStones, newLine)
	}
	rd.stones = newStones
}

func (rd *ReflectorDish) FlipVertically() {
	newStones := []string{}
	for i := len(rd.stones) - 1; i >= 0; i-- {
		newStones = append(newStones, rd.stones[i])
	}
	rd.stones = newStones
}

func (rd *ReflectorDish) Cycle() {
	rd.RotateClockwise() // north -> east
	rd.TiltEast()        // tilt north
	rd.RotateClockwise() // north -> south
	rd.TiltEast()        // tilt west
	rd.RotateClockwise() // north -> west
	rd.TiltEast()        // tilt south
	rd.RotateClockwise() // north -> north
	rd.TiltEast()        // tilt east
}
