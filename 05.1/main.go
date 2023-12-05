package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var seedToSoil [][]string
var soilToFertilizer [][]string
var fertilizerToWater [][]string
var waterToLight [][]string
var lightToTemperature [][]string
var temperatureToHumidity [][]string
var humidityToLocation [][]string

func getTable(filePath string) [][]string {
	var table [][]string

	input, _ := os.ReadFile(filePath)
	rows := strings.Split(string(input), "\n")
	for _, row := range rows {
		table = append(table, strings.Fields(row))
	}

	return table
}

func getCorrespondingKey(externalKey int, table [][]string) int {
	for _, row := range table {
		externalKeyStartRange, _ := strconv.Atoi(row[1])
		if externalKey < externalKeyStartRange {
			continue
		}

		rangeLength, _ := strconv.Atoi(row[2])
		offset := externalKey - externalKeyStartRange
		if offset > rangeLength {
			continue
		}

		internalKey, _ := strconv.Atoi(row[0])

		return internalKey + offset
	}

	return externalKey
}

func solve(seeds []string) int {
	var minLocation int
	hasLocation := false

	for _, seed := range seeds {
		seedId, _ := strconv.Atoi(seed)
		soil := getCorrespondingKey(seedId, seedToSoil)
		fertilizer := getCorrespondingKey(soil, soilToFertilizer)
		water := getCorrespondingKey(fertilizer, fertilizerToWater)
		light := getCorrespondingKey(water, waterToLight)
		temperature := getCorrespondingKey(light, lightToTemperature)
		humidity := getCorrespondingKey(temperature, temperatureToHumidity)
		location := getCorrespondingKey(humidity, humidityToLocation)

		fmt.Println(location)
		if !hasLocation {
			minLocation = location
			hasLocation = true
		} else {
			if location < minLocation {
				minLocation = location
			}
		}
	}

	return minLocation
}

func main() {
	seedToSoil = getTable("./05.1/seed_to_soil.txt")
	soilToFertilizer = getTable("./05.1/soil_to_fertilizer.txt")
	fertilizerToWater = getTable("./05.1/fertilizer_to_water.txt")
	waterToLight = getTable("./05.1/seed_to_soil.txt")
	lightToTemperature = getTable("./05.1/light_to_temperature.txt")
	temperatureToHumidity = getTable("./05.1/temperature_to_humidity.txt")
	humidityToLocation = getTable("./05.1/water_to_light.txt")

	seedsInput, _ := os.ReadFile("./05.1/seeds.txt")
	seeds := strings.Fields(string(seedsInput))

	answer := solve(seeds)

	fmt.Println(answer)
}
