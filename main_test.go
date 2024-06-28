package main

import (
	"testing"

	"github.com/rzabcio/adventofcode-2023/day04"
	"github.com/rzabcio/adventofcode-2023/day06"
	"github.com/rzabcio/adventofcode-2023/day07"
	"github.com/rzabcio/adventofcode-2023/day08"
	"github.com/rzabcio/adventofcode-2023/day09"
	"github.com/rzabcio/adventofcode-2023/day11"
	"github.com/rzabcio/adventofcode-2023/day13"
	"github.com/rzabcio/adventofcode-2023/day14"
	"github.com/rzabcio/adventofcode-2023/day15"
	"github.com/rzabcio/adventofcode-2023/day16"
	"github.com/rzabcio/adventofcode-2023/day18"
)

func TestDay01(t *testing.T) {
	got, want := Day01_1("input-files/day01-test1.txt"), 142
	if got != want {
		t.Errorf("Day01_1(test1) = %d; want %d", got, want)
	}
	got, want = Day01_2("input-files/day01-test2.txt"), 281
	if got != want {
		t.Errorf("Day01_2(test1) = %d; want %d", got, want)
	}
}

func TestDay02(t *testing.T) {
	got, want := Day02_1("input-files/day02-test1.txt"), 8
	if got != want {
		t.Errorf("Day02_1(test1) = %d; want %d", got, want)
	}
	got, want = Day02_2("input-files/day02-test1.txt"), 2286
	if got != want {
		t.Errorf("Day02_2(test1) = %d; want %d", got, want)
	}
}

func TestDay03(t *testing.T) {
	got, want := Day03_1("input-files/day03-test1.txt"), 4361
	if got != want {
		t.Errorf("Day03_1(test1) = %d; want %d", got, want)
	}
	got, want = Day03_2("input-files/day03-test1.txt"), 467835
	if got != want {
		t.Errorf("Day03_2(test1) = %d; want %d", got, want)
	}
}

func TestDay04(t *testing.T) {
	got, want := day04.Day04_1("input-files/day04-test1.txt"), 13
	if got != want {
		t.Errorf("Day04_1(test1) = %d; want %d", got, want)
	}
	got, want = day04.Day04_2("input-files/day04-test1.txt"), 30
	if got != want {
		t.Errorf("Day04_2(test1) = %d; want %d", got, want)
	}
}

// -- disabled - too long
// func TestDay05(t *testing.T) {
// 	got, want := day05.Day05_1("input-files/day05-test1.txt"), 35
// 	if got != want {
// 		t.Errorf("Day05_1(test1) = %d; want %d", got, want)
// 	}
// 	got, want = day05.Day05_2("input-files/day05-test1.txt"), 46
// 	if got != want {
// 		t.Errorf("Day05_2(test1) = %d; want %d", got, want)
// 	}
// }

func TestDay06(t *testing.T) {
	got, want := day06.Day06_1("input-files/day06-test1.txt"), 288
	if got != want {
		t.Errorf("Day06_1(test1) = %d; want %d", got, want)
	}
	got, want = day06.Day06_2("input-files/day06p2-test1.txt"), 71503
	if got != want {
		t.Errorf("Day06_2(test1) = %d; want %d", got, want)
	}
}

func TestDay07(t *testing.T) {
	got, want := day07.Day07_1("input-files/day07-test1.txt"), 6440
	if got != want {
		t.Errorf("Day07_1(test1) = %d; want %d", got, want)
	}
	got, want = day07.Day07_2("input-files/day07-test1.txt"), 5905
	if got != want {
		t.Errorf("Day07_2(test1) = %d; want %d", got, want)
	}
}

func TestDay08(t *testing.T) {
	got, want := day08.Day08_1("input-files/day08-test2.txt"), 6
	if got != want {
		t.Errorf("Day08_1(test1) = %d; want %d", got, want)
	}
	got, want = day08.Day08_2("input-files/day08p2-test1.txt"), 6
	if got != want {
		t.Errorf("Day08_2(test1) = %d; want %d", got, want)
	}
}

func TestDay09(t *testing.T) {
	got, want := day09.Day09_1("input-files/day09-test1.txt"), 114
	if got != want {
		t.Errorf("Day09_1(test1) = %d; want %d", got, want)
	}
	got, want = day09.Day09_2("input-files/day09-test1.txt"), 2
	if got != want {
		t.Errorf("Day09_2(test1) = %d; want %d", got, want)
	}
}

func TestDay11(t *testing.T) {
	got, want := day11.Day11_1("input-files/day11-test1.txt"), 374
	if got != want {
		t.Errorf("Day11_1(test1) = %d; want %d", got, want)
	}
	got, want = day11.Day11_2("input-files/day11-test1.txt"), 82000210
	if got != want {
		t.Errorf("Day11_2(test1) = %d; want %d", got, want)
	}
}

// func TestDay12(t *testing.T) {
// 	got, want := day12.Day12_1("input-files/day12-test1.txt"), 374
// 	if got != want {
// 		t.Errorf("Day12_1(test1) = %d; want %d", got, want)
// 	}
// 	got, want = day12.Day12_2("input-files/day12-test1.txt"), 0
// 	if got != want {
// 		t.Errorf("Day12_2(test1) = %d; want %d", got, want)
// 	}
// }

func TestDay13(t *testing.T) {
	got, want := day13.Day13_1("input-files/day13-test1.txt"), 405
	if got != want {
		t.Errorf("Day13_1(test1) = %d; want %d", got, want)
	}
	got, want = day13.Day13_2("input-files/day13-test1.txt"), 400
	if got != want {
		t.Errorf("Day13_2(test1) = %d; want %d", got, want)
	}
}

func TestDay14(t *testing.T) {
	got, want := day14.Day14_1("input-files/day14-test1.txt"), 136
	if got != want {
		t.Errorf("Day14_1(test1) = %d; want %d", got, want)
	}
	// -- disabled - too long
	// got, want = day14.Day14_2("input-files/day14-test1.txt"), 0
	// if got != want {
	// 	t.Errorf("Day14_2(test1) = %d; want %d", got, want)
	// }
}

func TestDay15(t *testing.T) {
	got, want := day15.Day15_1("input-files/day15-test1.txt"), 136
	if got != want {
		t.Errorf("Day15_1(test1) = %d; want %d", got, want)
	}
	got, want = day15.Day15_2("input-files/day15-test1.txt"), 0
	if got != want {
		t.Errorf("Day15_2(test1) = %d; want %d", got, want)
	}
}

func TestDay16(t *testing.T) {
	got, want := day16.Day16_1("input-files/day16-test1.txt"), 136
	if got != want {
		t.Errorf("Day16_1(test1) = %d; want %d", got, want)
	}
	got, want = day16.Day16_2("input-files/day16-test1.txt"), 0
	if got != want {
		t.Errorf("Day16_2(test1) = %d; want %d", got, want)
	}
}

func TestDay18(t *testing.T) {
	got, want := day18.Day18_1("input-files/day18-test1.txt"), 136
	if got != want {
		t.Errorf("Day18_1(test1) = %d; want %d", got, want)
	}
	got, want = day18.Day18_2("input-files/day18-test1.txt"), 0
	if got != want {
		t.Errorf("Day18_2(test1) = %d; want %d", got, want)
	}
}
