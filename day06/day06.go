package day06

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day06_1(filename string) (result int) {
	races := readRaces(filename)
	result = 1
	for _, race := range races {
		minIndex, _ := race.calculateMinPushTime()
		count := race.countOptions(minIndex)
		result *= count
	}
	return result
}

func Day06_2(filename string) (result int) {
	return 0
}

func readRaces(filename string) (races []Race) {
	for line := range utils.InputCh(filename) {

		fields := strings.Fields(line[9:])
		if strings.HasPrefix(line, "Time:") {
			for _, field := range fields {
				time, _ := strconv.Atoi(field)
				halfTime := int(math.Ceil(float64(time) / 2))
				races = append(races, Race{Time: time, HalfTime: halfTime})

			}
		}
		if strings.HasPrefix(line, "Distance:") {
			for i, field := range fields {
				distance, _ := strconv.Atoi(field)
				races[i].Distance = distance
			}
		}
	}
	return races
}

type Race struct {
	Time     int
	HalfTime int
	Distance int
}

func (r Race) calculateDistances() {
	for i := 1; i <= r.HalfTime; i++ {
		fmt.Printf(" - %d -> (%d-%d)*%d = %d\n", i, r.Time, i, i, (r.Time-i)*i)
	}
}

func (r Race) calculateMinPushTime() (i, minPushTime int) {
	for i = 1; i <= r.HalfTime; i++ {
		if (r.Time-i)*i > r.Distance {
			minPushTime = i
			break
		}
	}
	return i, minPushTime
}

func (r Race) countOptions(minIndex int) (count int) {
	return r.Time + 1 - 2*(minIndex)
}
