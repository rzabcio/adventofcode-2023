package day12

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day12_1(filename string) (result int) {
	sls := readSpringLines(filename)
	for i, sl := range sls {
		fmt.Printf("%d/%d = ", i+1, len(sls))
		sl.Unfold(5)
		count := sl.CountPossibilities()
		fmt.Printf("%d\n", count)
		result += count
	}
	return result
}

func Day12_2(filename string) (result int) {
	sls := readSpringLines(filename)
	for i, sl := range sls {
		fmt.Printf("%d/%d = ", i+1, len(sls))
		sl.Unfold(5)
		count := sl.CountPossibilities2()
		fmt.Printf("%d\n", count)
		result += count
	}
	return result
}

func readSpringLines(filename string) (sls []*SpringLine) {
	for line := range utils.InputCh(filename) {
		if strings.HasPrefix(line, "---") { // disable reading for debugging purposes
			continue
		}
		sl := NewSpringLine(line)
		sls = append(sls, sl)
	}
	return sls
}

type SpringLine struct {
	springs    string
	pattern    []int
	patternSum int
	r          *regexp.Regexp
}

func NewSpringLine(line string) (sl *SpringLine) {
	sl = &SpringLine{}
	fields := strings.Fields(line)
	sl.springs = fields[0]
	for _, s := range strings.Split(fields[1], ",") {
		i, _ := strconv.Atoi(s)
		sl.pattern = append(sl.pattern, i)
		sl.patternSum += i
	}
	sl.BuildR()
	return sl
}

// --- build regexp from pattern (number changes to count of # or ?, separated by . or ?)
// 1,1,3 => [#?]{1}[.?]+[#?]{1}[.?]+[#?]{3}
func (sl *SpringLine) BuildR() {
	var regs []string
	for _, p := range sl.pattern {
		regs = append(regs, fmt.Sprintf("[#?]{%d}", p))
	}
	rs := strings.Join(regs, "[\\.?]+")    // join elements
	rs = fmt.Sprintf("^[^#]*%s[^#]*$", rs) // add start and end
	sl.r = regexp.MustCompile(rs)
}

func (sl *SpringLine) String() string {
	return fmt.Sprintf("%s => %v => %v", sl.springs, sl.pattern, sl.r)
}

func (sl *SpringLine) CountPossibilities() (count int) {
	return sl.CheckNextVersion2(sl.springs, 0)
}

func (sl *SpringLine) CheckNextVersion(prev string, count int) int {
	if sl.r.MatchString(prev) {
		if !strings.Contains(prev, "?") {
			return count + 1 // base case - match and no more ?
		}
	} else {
		return count // stop recursion - no match
	}
	// generate two next branches - change first ? to # and .
	next1 := strings.Replace(prev, "?", "#", 1)
	next2 := strings.Replace(prev, "?", ".", 1)
	return sl.CheckNextVersion(next1, count) + sl.CheckNextVersion(next2, count)
}

// part 2
func (sl *SpringLine) Unfold(times int) {
	// multiply springs
	var newsprings []string
	var newpattern []int
	for i := 0; i < times; i++ {
		newsprings = append(newsprings, sl.springs)
		newpattern = append(newpattern, sl.pattern...)
	}
	sl.springs = strings.Join(newsprings, "?")
	sl.pattern = newpattern
	sl.patternSum *= times
	sl.BuildR()
}

func (sl *SpringLine) CountPossibilities2() (count int) {
	return sl.CheckNextVersion2(sl.springs, 0)
}

func (sl *SpringLine) CheckNextVersion2(prev string, count int) (result int) {
	if sl.r.MatchString(prev) {
		if !strings.Contains(prev, "?") {
			return count + 1 // base case - match and no more ?
		}
	} else {
		return count // stop recursion - no match
	}
	// generate two next branches - change first ? to # and .
	next1 := strings.Replace(prev, "?", "#", 1)
	next2 := strings.Replace(prev, "?", ".", 1)
	return sl.CheckNextVersion2(next1, count) + sl.CheckNextVersion2(next2, count)
}
