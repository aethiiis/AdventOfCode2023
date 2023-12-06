package main

import (
	"advent_of_code_2023/src/utilities"
	"math"
)

func Part1(lines []string) int {
	minLocation := math.MaxInt
	mapSeeds, mapSeed2Soil, mapSoil2Fertiliser, mapFertiliser2Water, mapWater2Light, mapLight2Temperature, mapTemperature2Humidity, mapHumidity2Location := utilities.ExtractMap(lines)

	for _, seed := range mapSeeds {

		var soil int
		var fertiliser int
		var water int
		var light int
		var temperature int
		var humidity int
		var location int

		var isInSeed = false
		var isInSoil = false
		var isInFert = false
		var isInWater = false
		var isInLight = false
		var isInTemp = false
		var isInHumi = false

		for i := range mapSeed2Soil {
			if seed >= mapSeed2Soil[i][1] && seed < mapSeed2Soil[i][1]+mapSeed2Soil[i][2] {
				isInSeed = true
				soil = seed - mapSeed2Soil[i][1] + mapSeed2Soil[i][0]
			}
		}
		if !isInSeed {
			soil = seed
		}

		for i := range mapSoil2Fertiliser {
			if soil >= mapSoil2Fertiliser[i][1] && soil < mapSoil2Fertiliser[i][1]+mapSoil2Fertiliser[i][2] {
				isInSoil = true
				fertiliser = soil - mapSoil2Fertiliser[i][1] + mapSoil2Fertiliser[i][0]
			}
		}
		if !isInSoil {
			fertiliser = soil
		}

		for i := range mapFertiliser2Water {
			if fertiliser >= mapFertiliser2Water[i][1] && fertiliser < mapFertiliser2Water[i][1]+mapFertiliser2Water[i][2] {
				isInFert = true
				water = fertiliser - mapFertiliser2Water[i][1] + mapFertiliser2Water[i][0]
			}
		}
		if !isInFert {
			water = fertiliser
		}

		for i := range mapWater2Light {
			if water >= mapWater2Light[i][1] && water < mapWater2Light[i][1]+mapWater2Light[i][2] {
				isInWater = true
				light = water - mapWater2Light[i][1] + mapWater2Light[i][0]
			}
		}
		if !isInWater {
			light = water
		}

		for i := range mapLight2Temperature {
			if light >= mapLight2Temperature[i][1] && light < mapLight2Temperature[i][1]+mapLight2Temperature[i][2] {
				isInLight = true
				temperature = light - mapLight2Temperature[i][1] + mapLight2Temperature[i][0]
			}
		}
		if !isInLight {
			temperature = light
		}

		for i := range mapTemperature2Humidity {
			if temperature >= mapTemperature2Humidity[i][1] && temperature < mapTemperature2Humidity[i][1]+mapTemperature2Humidity[i][2] {
				isInTemp = true
				humidity = temperature - mapTemperature2Humidity[i][1] + mapTemperature2Humidity[i][0]
			}
		}
		if !isInTemp {
			humidity = temperature
		}

		for i := range mapHumidity2Location {
			if humidity >= mapHumidity2Location[i][1] && humidity < mapHumidity2Location[i][1]+mapHumidity2Location[i][2] {
				isInHumi = true
				location = humidity - mapHumidity2Location[i][1] + mapHumidity2Location[i][0]
			}
		}
		if !isInHumi {
			location = humidity
		}

		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}
