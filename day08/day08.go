package day08

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day08_1(filename string) (result int) {
	m := ReadMap(filename)
	fmt.Println(m)
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
	currents := m.FindStarting()
	lenght := len(currents)
	fmt.Printf("Starting: %s\n", currents)
	for {
		instruction := string(m.Instruction[step%len(m.Instruction)])

		needBreak := true
		for i := 0; i < lenght; i++ {
			currents[i] = m.Elements[currents[i]][instruction]
			if currents[i][2] != 'Z' {
				needBreak = false
			}
		}
		fmt.Printf("- %d: %s => %s\n", step, instruction, currents)
		if needBreak {
			break
		}
		step++
	}
	return step + 1
}

func (m *Map) FindStarting() (starting []string) {
	for element := range m.Elements {
		if strings.HasSuffix(element, "A") {
			starting = append(starting, element)
		}
	}
	return starting
}
