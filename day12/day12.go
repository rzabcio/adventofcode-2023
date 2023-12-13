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
	for _, sl := range sls {
		result += sl.CheckNextVersion(sl.springs, 0)
	}
	return result
}

func Day12_2(filename string) (result int) {
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
	springs string
	pattern []int
	r       *regexp.Regexp
}

func NewSpringLine(line string) (sl *SpringLine) {
	sl = &SpringLine{}
	fields := strings.Fields(line)
	sl.springs = fields[0]
	for _, s := range strings.Split(fields[1], ",") {
		i, _ := strconv.Atoi(s)
		sl.pattern = append(sl.pattern, i)
	}
	// --- build regexp from pattern (number changes to count of # or ?, separated by . or ?)
	// 1,1,3 => [#?]{1}[.?]+[#?]{1}[.?]+[#?]{3}
	var regs []string
	for _, p := range sl.pattern {
		regs = append(regs, fmt.Sprintf("[#?]{%d}", p))
	}
	rs := strings.Join(regs, "[\\.?]+")    // join elements
	rs = fmt.Sprintf("^[^#]*%s[^#]*$", rs) // add start and end
	sl.r = regexp.MustCompile(rs)
	return sl
}

func (sl *SpringLine) String() string {
	return fmt.Sprintf("%s => %v => %v", sl.springs, sl.pattern, sl.r)
}

func (sl *SpringLine) CheckNextVersion(prev string, count int) int {
	// -- base case - no more ? in prev
	if !strings.Contains(prev, "?") {
		if sl.r.MatchString(prev) {
			// fmt.Printf("*  %s\n", prev)
			return count + 1
		}
		return count
	}
	// if does not match - return (end this graph branch)
	if !sl.r.MatchString(prev) {
		return count
	}
	// generate two next branches - change first ? to # and .
	next1 := strings.Replace(prev, "?", "#", 1)
	next2 := strings.Replace(prev, "?", ".", 1)
	return sl.CheckNextVersion(next1, count) + sl.CheckNextVersion(next2, count)
}
