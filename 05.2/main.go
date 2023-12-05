package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var seedToSoil [][]int
var seedToSoilKeyBounds [2]int
var soilToFertilizer [][]int
var soilToFertilizerKeyBounds [2]int
var fertilizerToWater [][]int
var fertilizerToWaterKeyBounds [2]int
var waterToLight [][]int
var waterToLightKeyBounds [2]int
var lightToTemperature [][]int
var lightToTemperatureKeyBounds [2]int
var temperatureToHumidity [][]int
var temperatureToHumidityKeyBounds [2]int
var humidityToLocation [][]int
var humidityToLocationKeyBounds [2]int

func getTable(filePath string) ([][]int, int, int) {
	var table [][]int
	var lowKey, topKey int
	lowKeyIsSet := false

	input, _ := os.ReadFile(filePath)
	rows := strings.Split(string(input), "\n")
	for _, row := range rows {
		fields := strings.Fields(row)
		if len(fields) < 3 {
			continue
		}

		key1RangeStart, _ := strconv.Atoi(fields[0])
		key2RangeStart, _ := strconv.Atoi(fields[1])
		rangeLength, _ := strconv.Atoi(fields[2])
		key2RangeEnd := key2RangeStart + rangeLength - 1

		table = append(table, []int{key1RangeStart, key2RangeStart, rangeLength, key2RangeEnd})

		if !lowKeyIsSet {
			lowKey = key2RangeStart
			lowKeyIsSet = true
		} else {
			if key2RangeStart < lowKey {
				lowKey = key2RangeStart
			}
		}

		if key2RangeEnd > topKey {
			topKey = key2RangeEnd
		}
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i][1] < table[j][1]
	})

	return table, lowKey, topKey
}

func getCorrespondingKey(externalKey int, table [][]int, keyBounds [2]int) int {
	if externalKey < keyBounds[0] {
		// fmt.Println("Low from the start")
		return externalKey
	}
	if externalKey > keyBounds[1] {
		// fmt.Println("High from the start")
		return externalKey
	}

	for _, row := range table {
		externalKeyStartRange := row[1]
		if externalKey < externalKeyStartRange {
			// fmt.Println("Too high", externalKey, row)
			break
		}

		if externalKey > row[3] {
			// fmt.Println("Too low", externalKey, row)
			continue
		}

		offset := externalKey - externalKeyStartRange

		return row[0] + offset
	}

	return externalKey
}

func solve(seeds []string) int {
	var minLocation int
	hasLocation := false

	for i := 0; i < len(seeds); i += 2 {
		seedsStartRange, _ := strconv.Atoi(seeds[i])
		seedsRangeLength, _ := strconv.Atoi(seeds[i+1])

		for seedId := seedsStartRange; seedId < seedsStartRange+seedsRangeLength; seedId++ {
			soil := getCorrespondingKey(seedId, seedToSoil, seedToSoilKeyBounds)
			fertilizer := getCorrespondingKey(soil, soilToFertilizer, soilToFertilizerKeyBounds)
			water := getCorrespondingKey(fertilizer, fertilizerToWater, fertilizerToWaterKeyBounds)
			light := getCorrespondingKey(water, waterToLight, waterToLightKeyBounds)
			temperature := getCorrespondingKey(light, lightToTemperature, lightToTemperatureKeyBounds)
			humidity := getCorrespondingKey(temperature, temperatureToHumidity, temperatureToHumidityKeyBounds)
			location := getCorrespondingKey(humidity, humidityToLocation, humidityToLocationKeyBounds)

			// fmt.Println(seedId, soil, fertilizer, water, light, temperature, humidity, location)

			if !hasLocation {
				minLocation = location
				hasLocation = true
			} else {
				if location < minLocation {
					minLocation = location
				}
			}
		}
	}

	return minLocation
}

func main() {
	seedToSoil, seedToSoilKeyBounds[0], seedToSoilKeyBounds[1] = getTable("./05.2/seed_to_soil.txt")
	soilToFertilizer, soilToFertilizerKeyBounds[0], soilToFertilizerKeyBounds[1] = getTable("./05.2/soil_to_fertilizer.txt")
	fertilizerToWater, fertilizerToWaterKeyBounds[0], fertilizerToWaterKeyBounds[1] = getTable("./05.2/fertilizer_to_water.txt")
	waterToLight, waterToLightKeyBounds[0], waterToLightKeyBounds[1] = getTable("./05.2/water_to_light.txt")
	lightToTemperature, lightToTemperatureKeyBounds[0], lightToTemperatureKeyBounds[1] = getTable("./05.2/light_to_temperature.txt")
	temperatureToHumidity, temperatureToHumidityKeyBounds[0], temperatureToHumidityKeyBounds[1] = getTable("./05.2/temperature_to_humidity.txt")
	humidityToLocation, humidityToLocationKeyBounds[0], humidityToLocationKeyBounds[1] = getTable("./05.2/humidity_to_location.txt")

	seedsInput, _ := os.ReadFile("./05.2/seeds.txt")
	seeds := strings.Fields(string(seedsInput))

	answer := solve(seeds)

	fmt.Println(answer)
}
