package day11

import (
	"fmt"
	"math"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day11_1(filename string) (result int) {
	u := NewUniverse(filename)
	result = u.DistanceSum()
	return result
}

func Day11_2(filename string) (result int) {
	return result
}

type Universe struct {
	emptyCols map[int]bool
	emptyRows map[int]bool
	galaxies  []*Galaxy
}

func NewUniverse(filename string) (u *Universe) {
	u = &Universe{}
	u.emptyCols = make(map[int]bool)
	u.emptyRows = make(map[int]bool)

	count := 1
	row := 0
	for line := range utils.InputCh(filename) {
		for col, char := range line {
			if char == '#' {
				u.galaxies = append(u.galaxies, &Galaxy{count, col, row})
				count++
				u.emptyCols[col] = false
				u.emptyRows[row] = false
			} else {
				if _, ok := u.emptyCols[col]; !ok {
					u.emptyCols[col] = true
				}
				if _, ok := u.emptyRows[row]; !ok {
					u.emptyRows[row] = true
				}
			}
		}
		row++
	}

	return u
}

func (u *Universe) Distance(g1, g2 *Galaxy) (dist int) {
	dist = int(math.Abs(float64(g1.x-g2.x)) + math.Abs(float64(g1.y-g2.y)))
	for i := utils.Min([]int{g1.x, g2.x}) + 1; i < utils.Max([]int{g1.x, g2.x}); i++ {
		if u.emptyCols[i] {
			dist++
		}
	}
	for i := utils.Min([]int{g1.y, g2.y}) + 1; i < utils.Max([]int{g1.y, g2.y}); i++ {
		if u.emptyRows[i] {
			dist++
		}
	}
	return dist
}

func (u *Universe) DistanceSum() (sum int) {
	for i := 0; i < len(u.galaxies)-1; i++ {
		for j := i + 1; j < len(u.galaxies); j++ {
			dist := u.Distance(u.galaxies[i], u.galaxies[j])
			sum += dist
		}
	}
	return sum
}

func (u *Universe) String() string {
	result := ""
	result += fmt.Sprintf("emptyCols: %+v\n", u.emptyCols)
	result += fmt.Sprintf("emptyRows: %+v\n", u.emptyRows)
	for _, galaxy := range u.galaxies {
		result += fmt.Sprintf("galaxy: %+v\n", galaxy)
	}
	return result
}

type Galaxy struct {
	no   int
	x, y int
}
