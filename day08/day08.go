package day08

import (
	"regexp"
	"slices"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day08_1(filename string) (result int) {
	m := ReadMap(filename)
	result = m.GoToZZZ()
	return result
}

func Day08_2(filename string) (result int) {
	m := ReadMap(filename)
	result = m.GoToAllZ()
	return result
}

type Map struct {
	Instruction string
	Elements    map[string]map[string]string
}

func ReadMap(filename string) (m *Map) {
	elementRegex := regexp.MustCompile(`^([0-9A-Z]{3}) = \(([0-9A-Z]{3}), ([0-9A-Z]{3})\)$`)
	m = new(Map)
	for line := range utils.InputCh(filename) {
		if m.Instruction == "" {
			m.Instruction = line
			continue
		}
		if len(line) == 0 {
			continue
		}
		match := elementRegex.FindStringSubmatch(line)
		if len(match) != 4 {
			panic("Invalid line: " + line)
		}
		if m.Elements == nil {
			m.Elements = make(map[string]map[string]string)
		}
		if m.Elements[match[1]] == nil {
			m.Elements[match[1]] = make(map[string]string)
		}
		m.Elements[match[1]]["L"] = match[2]
		m.Elements[match[1]]["R"] = match[3]
	}
	return m
}

// part 1
func (m *Map) GoToZZZ() (i int) {
	current := "AAA"
	for {
		instruction := m.Instruction[i%len(m.Instruction)]
		current = m.Elements[current][string(instruction)]
		if current == "ZZZ" {
			break
		}
		i++
	}
	return i + 1
}

// part 2
func (m *Map) GoToAllZ() (step int) {
	startings := m.FindStarting()
	// zeds := make(map[int][]int)
	var zeds []int
	for _, starting := range startings {
		zeds = append(zeds, m.FindStepsWithZ(starting)[0])
	}
	return LCM(zeds...)
}

func (m *Map) FindStarting() (starting []string) {
	for element := range m.Elements {
		if strings.HasSuffix(element, "A") {
			starting = append(starting, element)
		}
	}
	return starting
}

func (m *Map) FindStepsWithZ(starting string) (zeds []int) {
	step := 0
	visited := []string{starting}
	current := starting
	for {
		instruction := string(m.Instruction[step%len(m.Instruction)])
		current = m.Elements[current][instruction]
		if slices.Contains(visited, current) && step%len(m.Instruction) == 0 && len(zeds) > 0 {
			return zeds
		}
		visited = append(visited, current)
		if current[2] == 'Z' {
			zeds = append(zeds, step+1)
		}
		step++
	}
}

// code copied from: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)
	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}
