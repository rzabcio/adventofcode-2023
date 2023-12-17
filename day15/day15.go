package day15

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day15_1(filename string) (result int) {
	for line := range utils.InputRows(filename) {
		for _, cmd := range strings.Split(line, ",") {
			result += decodeLine(cmd)
		}
	}
	return result
}

func Day15_2(filename string) (result int) {
	boxes := &Boxes{boxes: make(map[int]map[string]*Lens)}
	for line := range utils.InputRows(filename) {
		for _, cmd := range strings.Split(line, ",") {
			lens := NewLens(cmd)
			if lens.focal > 0 {
				boxes.add(lens)
			} else {
				boxes.remove(lens)
			}
		}
	}

	for _, lenses := range boxes.boxes {
		for _, lens := range lenses {
			result += lens.FocusingPower()
		}
	}

	return result
}

func decodeLine(line string) (result int) {
	for _, r := range line {
		result += int(r)
		// fmt.Printf("%c => %d", r, result)
		result *= 17
		// fmt.Printf(" => %d", result)
		result %= 256
		// fmt.Printf(" => %d\n", result)
	}
	// fmt.Println()
	return result
}

type Boxes struct {
	boxes map[int]map[string]*Lens
}

func (b *Boxes) String() (result string) {
	for box, lenses := range b.boxes {
		result += fmt.Sprintf("Box %d: ", box)
		var lensesSlice []*Lens
		for _, lens := range lenses {
			lensesSlice = append(lensesSlice, lens)
		}
		sort.Slice(lensesSlice, func(i, j int) bool {
			return lensesSlice[i].slot < lensesSlice[j].slot
		})
		for _, lens := range lensesSlice {
			result += fmt.Sprintf("%s, ", lens)
		}
		result += "\n"
	}
	return result
}

func (b *Boxes) add(lens *Lens) {
	if b.boxes[lens.box] == nil {
		b.boxes[lens.box] = make(map[string]*Lens)
	}
	if b.boxes[lens.box][lens.label] != nil {
		b.boxes[lens.box][lens.label].focal = lens.focal
		return
	}
	lens.slot = len(b.boxes[lens.box]) + 1
	if b.boxes[lens.box] == nil {
		b.boxes[lens.box] = make(map[string]*Lens)
	}
	b.boxes[lens.box][lens.label] = lens
}

func (b *Boxes) remove(lens *Lens) {
	if b.boxes[lens.box] == nil {
		return
	}
	if b.boxes[lens.box][lens.label] == nil {
		return
	}
	empiedSlot := b.boxes[lens.box][lens.label].slot
	delete(b.boxes[lens.box], lens.label)
	for _, lens := range b.boxes[lens.box] {
		if lens.slot > empiedSlot {
			lens.slot--
		}
	}
}

type Lens struct {
	box   int
	label string
	focal int
	slot  int
}

func NewLens(cmd string) (lens *Lens) {
	lens = &Lens{}
	if strings.Contains(cmd, "=") {
		parts := strings.Split(cmd, "=")
		lens.label = parts[0]
		lens.focal, _ = strconv.Atoi(parts[1])
		lens.box = decodeLine(lens.label)
	} else if strings.Contains(cmd, "-") {
		parts := strings.Split(cmd, "-")
		lens.label = parts[0]
		lens.box = decodeLine(lens.label)
	} else {
		fmt.Println("!!! unknown command: ", cmd)
	}
	return lens
}

func (l *Lens) String() (result string) {
	if l.focal == 0 {
		return fmt.Sprintf("%s-", l.label)
	}
	if l.slot == 0 {
		return fmt.Sprintf("%s=%d", l.label, l.focal)
	}
	return fmt.Sprintf("%s=%d[%d]", l.label, l.focal, l.slot)
}

func (l *Lens) FocusingPower() (result int) {
	return (l.box + 1) * l.slot * l.focal
}
