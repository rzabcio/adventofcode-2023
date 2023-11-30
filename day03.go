package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Day03_1(filename string) (result int) {
	plan := NewPlan(filename)
	return plan.NeededSum()
}

func Day03_2(filename string) (result int) {
	plan := NewPlan(filename)
	return plan.GearRation
}

type Plan struct {
	Parts        []*Part
	PartsDiagram [][]*Part
	GearRation   int
}

func NewPlan(filename string) *Plan {
	partRegexp := regexp.MustCompile("(\\d+)")

	// first pass: read all parts
	plan := &Plan{}
	plan.PartsDiagram = make([][]*Part, 0)
	for line := range inputCh(filename) {
		planLine := make([]*Part, len(line))
		matches := partRegexp.FindAllStringSubmatch(line, -1)
		if matches == nil {
			plan.PartsDiagram = append(plan.PartsDiagram, planLine)
			continue
		}
		lastIndex := 0 // needed for my stupid idea to find part position by its index in the line (cas 495 vs 5), so we need to look for in in substring
		for _, match := range matches {
			part := &Part{}
			part.NoStr = match[1]
			part.No, _ = strconv.Atoi(match[1])
			part.Index = strings.Index(line[lastIndex:], match[1]) + lastIndex
			part.Length = len(match[1])
			lastIndex = part.Index + part.Length
			part.Needed = false
			for i := part.Index; i < part.Index+part.Length; i++ {
				planLine[i] = part
			}
			plan.Parts = append(plan.Parts, part)
		}
		plan.PartsDiagram = append(plan.PartsDiagram, planLine)
	}

	// second pass: find non-needed parts
	row := 0
	for line := range inputCh(filename) {
		for col, r := range line {
			connectedParts := make([]*Part, 0)
			if r == '.' || unicode.IsDigit(r) {
				continue
			}
			for i := col - 1; i <= col+1; i++ {
				if i < 0 || i >= len(plan.PartsDiagram[row]) {
					continue
				}
			LOOP:
				for j := row - 1; j <= row+1; j++ {
					if j < 0 || j >= len(plan.PartsDiagram) {
						continue
					}
					part := plan.PartsDiagram[j][i]
					if part == nil {
						continue
					}
					part.Needed = true
					// part 2 - finding all part connected to any *
					if r != '*' {
						continue
					}
					for _, connectedPart := range connectedParts {
						if connectedPart == part {
							continue LOOP
						}
					}
					connectedParts = append(connectedParts, part)
				}
			}
			// part 2
			if len(connectedParts) != 2 { // getting only gears connected to exactly 2 other parts
				continue
			}
			ratio := 1
			for _, part := range connectedParts {
				ratio *= part.No
			}
			plan.GearRation += ratio
			// fmt.Printf("%s: %v -> %d -> %d\n", string(r), connectedParts, ratio, plan.GearRation)
		}
		row++
	}
	return plan
}

func (p *Plan) String() string {
	result := "--- FLAT:\n"
	for _, part := range p.Parts {
		result += fmt.Sprintf("%s %d %d %d %v\n", part.NoStr, part.No, part.Index, part.Length, part.Needed)
	}
	result += "--- DIAGRAM:\n"
	for _, line := range p.PartsDiagram {
		for _, part := range line {
			if part == nil {
				result += " . "
			} else {
				result += part.NoStr
				if part.Needed {
					result += " "
				} else {
					result += "-"
				}
			}
		}
		result += "\n"
	}
	return result
}

func (p Plan) NeededParts() []*Part {
	result := make([]*Part, 0)
	for _, part := range p.Parts {
		if part.Needed {
			result = append(result, part)
		}
	}
	return result
}

func (p Plan) NeededSum() int {
	result := 0
	for _, part := range p.NeededParts() {
		result += part.No
	}
	return result
}

type Part struct {
	NoStr  string
	No     int
	Index  int
	Length int
	Needed bool
}
