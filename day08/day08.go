package day08

import (
	"fmt"
	"regexp"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day08_1(filename string) (result int) {
	m := ReadMap(filename)
	fmt.Println(m)
	result = m.GoToZZZ()
	return result
}

func Day08_2(filename string) (result int) {
	return result
}

type Map struct {
	Instruction string
	Elements    map[string]map[string]string
}

func ReadMap(filename string) (m *Map) {
	elementRegex := regexp.MustCompile(`^([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)$`)
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
