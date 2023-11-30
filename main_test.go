package main

import (
	"testing"
)

func TestDay01(t *testing.T) {
	got, want := Day01_1("input-files/day01-test1.txt"), 142
	if got != want {
		t.Errorf("Day01_1(test1) = %d; want %d", got, want)
	}
	// got, want = Day01_2("input-files/day01-test1.txt"), 45000
	// if got != want {
	// 	t.Errorf("Day01_2(test1) = %d; want %d", got, want)
	// }
}
