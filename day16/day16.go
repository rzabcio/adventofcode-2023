package day16

import (
	"fmt"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day16_1(filename string) (result int) {
	c := NewContraption(filename)
	result = c.Run(-1, 0)
	return result
}

func Day16_2(filename string) (result int) {
	c := NewContraption(filename)
	for startX := 0; startX < len(c.grid[0]); startX++ {
		subResult := c.Run(startX, -1)
		if subResult > result {
			result = subResult
		}
		subResult = c.Run(startX, len(c.grid))
		if subResult > result {
			result = subResult
		}
	}
	for startY := 0; startY < len(c.grid); startY++ {
		subResult := c.Run(-1, startY)
		if subResult > result {
			result = subResult
		}
		subResult = c.Run(len(c.grid[0]), startY)
		if subResult > result {
			result = subResult
		}
	}
	return result
}

type Contraption struct {
	startX, startY int
	grid           []string
	visitVert      [][]bool
	visitHoriz     [][]bool
	beams          []*Beam
}

func NewContraption(filename string) (c *Contraption) {
	c = &Contraption{}
	for line := range utils.InputCh(filename) {
		c.grid = append(c.grid, line)
		c.visitVert = append(c.visitVert, make([]bool, len(line)))
		c.visitHoriz = append(c.visitHoriz, make([]bool, len(line)))
	}
	c.beams = append(c.beams, &Beam{x: -1, y: 0, dX: 1, dY: 0})
	return c
}

func (c *Contraption) Run(startX, startY int) (result int) {
	c.startX = startX
	c.startY = startY
	c.Reset()
	for !c.finished() {
		c.nextStep()
	}
	result = c.CountVisited()
	return result
}

func (c *Contraption) Reset() {
	c.beams = []*Beam{}
	var dX, dY int
	switch {
	case c.startX == -1:
		dX, dY = 1, 0
	case c.startX == len(c.grid[0]):
		dX, dY = -1, 0
	case c.startY == -1:
		dX, dY = 0, 1
	case c.startY == len(c.grid):
		dX, dY = 0, -1
	}
	c.beams = append(c.beams, &Beam{x: c.startX, y: c.startY, dX: dX, dY: dY})
	for y := range c.visitVert {
		for x := range c.visitVert[y] {
			c.visitVert[y][x] = false
			c.visitHoriz[y][x] = false
		}
	}
}

func (c *Contraption) String() (result string) {
	for y := range c.grid {
		result += c.grid[y] + "\n"
	}
	result += "\n"
	for y := range c.visitVert {
		for x := range c.visitVert[y] {
			if c.visitVert[y][x] || c.visitHoriz[y][x] {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	result += "\n"
	for _, beam := range c.beams {
		result += beam.String() + "\n"
	}

	return result
}

func (c *Contraption) nextStep() {
	for _, beam := range c.beams {
		c.nextStepBeam(beam)
	}
}

func (c *Contraption) nextStepBeam(beam *Beam) {
	if beam.finished {
		return
	}
	newx := beam.x + beam.dX
	newy := beam.y + beam.dY
	// check if we're out of bounds
	if newx < 0 || newx >= len(c.grid[0]) || newy < 0 || newy >= len(c.grid) {
		beam.finished = true
		return
	}
	// check if we've already visited this spot
	if beam.dX == 0 {
		if c.visitVert[newy][newx] && c.grid[newy][newx] != '/' && c.grid[newy][newx] != '\\' {
			beam.finished = true
			return
		}
		c.visitVert[newy][newx] = true
	} else {
		if c.visitHoriz[newy][newx] && c.grid[newy][newx] != '/' && c.grid[newy][newx] != '\\' {
			beam.finished = true
			return
		}
		c.visitHoriz[newy][newx] = true
	}
	// check what we hit
	switch c.grid[newy][newx] {
	case '|':
		if beam.dX != 0 {
			beam.dX, beam.dY = 0, -1
			c.beams = append(c.beams, &Beam{x: newx, y: newy, dX: 0, dY: 1})
		}
	case '/':
		beam.dX, beam.dY = -beam.dY, -beam.dX
	case '\\':
		beam.dX, beam.dY = beam.dY, beam.dX
	case '-':
		if beam.dY != 0 {
			beam.dX, beam.dY = -1, 0
			c.beams = append(c.beams, &Beam{x: newx, y: newy, dX: 1, dY: 0})
		}
	}
	// set new position
	beam.x = newx
	beam.y = newy
}

func (c *Contraption) finished() bool {
	for _, beam := range c.beams {
		if !beam.finished {
			return false
		}
	}
	return true
}

func (c *Contraption) CountVisited() (result int) {
	for y := range c.visitVert {
		for x := range c.visitVert[y] {
			if c.visitVert[y][x] || c.visitHoriz[y][x] {
				result++
			}
		}
	}
	return result
}

type Beam struct {
	x        int
	y        int
	dX       int
	dY       int
	finished bool
}

func (b *Beam) String() (result string) {
	result += fmt.Sprintf("(%d,%d)", b.x, b.y)
	if b.finished {
		result += "#"
		return result
	}
	switch {
	case b.dX == 1 && b.dY == 0:
		result += ">"
	case b.dX == -1 && b.dY == 0:
		result += "<"
	case b.dX == 0 && b.dY == 1:
		result += "v"
	case b.dX == 0 && b.dY == -1:
		result += "^"
	}
	return result
}
