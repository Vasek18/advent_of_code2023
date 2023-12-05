package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var seedToSoil [][]int
var soilToFertilizer [][]int
var fertilizerToWater [][]int
var waterToLight [][]int
var lightToTemperature [][]int
var temperatureToHumidity [][]int
var humidityToLocation [][]int

func getTable(filePath string) [][]int {
	var table [][]int

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
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i][1] < table[j][1]
	})

	return table
}

func getCorrespondingKey(externalKey int, table [][]int) int {
	for _, row := range table {
		externalKeyStartRange := row[1]
		if externalKey < externalKeyStartRange {
			fmt.Println("Too high", externalKey, row)
			continue
		}

		if externalKey > row[3] {
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
			soil := getCorrespondingKey(seedId, seedToSoil)
			fertilizer := getCorrespondingKey(soil, soilToFertilizer)
			water := getCorrespondingKey(fertilizer, fertilizerToWater)
			light := getCorrespondingKey(water, waterToLight)
			temperature := getCorrespondingKey(light, lightToTemperature)
			humidity := getCorrespondingKey(temperature, temperatureToHumidity)
			location := getCorrespondingKey(humidity, humidityToLocation)

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
	seedToSoil = getTable("./05.2/seed_to_soil.txt")
	soilToFertilizer = getTable("./05.2/soil_to_fertilizer.txt")
	fertilizerToWater = getTable("./05.2/fertilizer_to_water.txt")
	waterToLight = getTable("./05.2/water_to_light.txt")
	lightToTemperature = getTable("./05.2/light_to_temperature.txt")
	temperatureToHumidity = getTable("./05.2/temperature_to_humidity.txt")
	humidityToLocation = getTable("./05.2/humidity_to_location.txt")

	seedsInput, _ := os.ReadFile("./05.2/seeds.txt")
	seeds := strings.Fields(string(seedsInput))

	answer := solve(seeds)

	fmt.Println(answer)
}
