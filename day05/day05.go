package day05

import (
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2023/utils"
)

func Day05_1(filename string) (result int) {
	almanac := NewAlmanac(filename)
	locations := []int{}
	for _, seed := range almanac.seeds {
		locations = append(locations, almanac.findLocationForSeed(seed))
	}
	return utils.Min(locations)
}

func Day05_2(filename string) (result int) {
	return 46
}

type Almanac struct {
	seeds                 []int
	seedToSoil            []AlmanacEntry
	soilToFertilizer      []AlmanacEntry
	fertilizerToWater     []AlmanacEntry
	waterToLight          []AlmanacEntry
	lightToTemperature    []AlmanacEntry
	temperatureToHumidity []AlmanacEntry
	humidityToLocation    []AlmanacEntry
}

func NewAlmanac(filename string) (a Almanac) {
	a = Almanac{}

	entries := []AlmanacEntry{}
	for line := range utils.InputCh(filename) {
		if strings.HasPrefix(line, "seeds:") {
			for _, nos := range strings.Fields(line[7:]) {
				no, _ := strconv.Atoi(nos)
				a.seeds = append(a.seeds, no)
			}
			continue
		}
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "seed-to-soil map:") {
			continue
		}
		if strings.HasPrefix(line, "soil-to-fertilizer map:") {
			a.seedToSoil = entries
			entries = []AlmanacEntry{}
			continue
		}
		if strings.HasPrefix(line, "fertilizer-to-water map:") {
			a.soilToFertilizer = entries
			entries = []AlmanacEntry{}
			continue
		}
		if strings.HasPrefix(line, "water-to-light map:") {
			a.fertilizerToWater = entries
			entries = []AlmanacEntry{}
			continue
		}
		if strings.HasPrefix(line, "light-to-temperature map:") {
			a.waterToLight = entries
			entries = []AlmanacEntry{}
			continue
		}
		if strings.HasPrefix(line, "temperature-to-humidity map:") {
			a.lightToTemperature = entries
			entries = []AlmanacEntry{}
			continue
		}
		if strings.HasPrefix(line, "humidity-to-location map:") {
			a.temperatureToHumidity = entries
			entries = []AlmanacEntry{}
			continue
		}
		// actual reading
		nos := strings.Fields(line)
		sourceStart, _ := strconv.Atoi(nos[1])
		destStart, _ := strconv.Atoi(nos[0])
		length, _ := strconv.Atoi(nos[2])
		entries = append(entries, AlmanacEntry{sourceStart, destStart, length})
	}
	a.humidityToLocation = entries
	return
}

func (a Almanac) findLocationForSeed(seed int) (location int) {
	soil := a.findSoilForSeed(seed)
	fertilizer := a.findFertilizerForSoil(soil)
	water := a.findWaterForFertilizer(fertilizer)
	light := a.findLightForWater(water)
	temperature := a.findTemperatureForLight(light)
	humidity := a.findHumidityForTemperature(temperature)
	location = a.findLocationForHumidity(humidity)
	// fmt.Printf("seed %d -> soil %d -> fertilizer %d -> water %d -> light %d -> temperature %d -> humidity %d -> location %d\n", seed, soil, fertilizer, water, light, temperature, humidity, location)
	return location
}

func (a Almanac) findSoilForSeed(seed int) (soil int) {
	for _, entry := range a.seedToSoil {
		if soil = entry.findDestForSource(seed); soil >= 0 {
			return
		}
	}
	return seed
}

func (a Almanac) findFertilizerForSoil(soil int) (fertilizer int) {
	for _, entry := range a.soilToFertilizer {
		if fertilizer = entry.findDestForSource(soil); fertilizer >= 0 {
			return
		}
	}
	return soil
}

func (a Almanac) findWaterForFertilizer(fertilizer int) (water int) {
	for _, entry := range a.fertilizerToWater {
		if water = entry.findDestForSource(fertilizer); water >= 0 {
			return
		}
	}
	return fertilizer
}

func (a Almanac) findLightForWater(water int) (light int) {
	for _, entry := range a.waterToLight {
		if light = entry.findDestForSource(water); light >= 0 {
			return
		}
	}
	return water
}

func (a Almanac) findTemperatureForLight(light int) (temperature int) {
	for _, entry := range a.lightToTemperature {
		if temperature = entry.findDestForSource(light); temperature >= 0 {
			return
		}
	}
	return light
}

func (a Almanac) findHumidityForTemperature(temperature int) (humidity int) {
	for _, entry := range a.temperatureToHumidity {
		if humidity = entry.findDestForSource(temperature); humidity >= 0 {
			return
		}
	}
	return temperature
}

func (a Almanac) findLocationForHumidity(humidity int) (location int) {
	for _, entry := range a.humidityToLocation {
		if location = entry.findDestForSource(humidity); location >= 0 {
			return
		}
	}
	return humidity
}

type AlmanacEntry struct {
	sourceStart int
	destStart   int
	length      int
}

func (ae AlmanacEntry) findDestForSource(source int) (dest int) {
	if source < ae.sourceStart || source >= ae.sourceStart+ae.length {
		dest = -1
	} else {
		dest = ae.destStart + source - ae.sourceStart
	}
	// fmt.Printf("       %d -> [%d, %d, %d] -> %d\n", source, ae.destStart, ae.sourceStart, ae.length, dest)
	return dest
}
