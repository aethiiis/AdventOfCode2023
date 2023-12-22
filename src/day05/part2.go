package main

import (
	"advent_of_code_2023/src/utilities"
	"math"
)

func Part2(lines []string) int {

	mapSeeds, mapSeed2Soil, mapSoil2Fertiliser, mapFertiliser2Water, mapWater2Light, mapLight2Temperature, mapTemperature2Humidity, mapHumidity2Location := utilities.ExtractMap(lines)
	maps := [][][]int{mapSeed2Soil, mapSoil2Fertiliser, mapFertiliser2Water, mapWater2Light, mapLight2Temperature, mapTemperature2Humidity, mapHumidity2Location}

	var seeds [][]int
	for i := 0; i < len(mapSeeds); i += 2 {
		seeds = append(seeds, []int{mapSeeds[i], mapSeeds[i] + mapSeeds[i+1]})
	}

	for _, Map := range maps {
		var newSeeds [][]int
		for len(seeds) > 0 {
			start := seeds[len(seeds)-1][0]
			end := seeds[len(seeds)-1][1]
			seeds = seeds[:len(seeds)-1]

			var empty bool
			for _, triplet := range Map {
				unionStart := int(math.Max(float64(start), float64(triplet[1])))
				unionEnd := int(math.Min(float64(end), float64(triplet[1]+triplet[2])))
				empty = true
				if unionStart < unionEnd {
					newSeeds = append(newSeeds, []int{unionStart - triplet[1] + triplet[0], unionEnd - triplet[1] + triplet[0]})
					if unionStart > start {
						seeds = append(seeds, []int{start, unionStart})
					}
					if end > unionEnd {
						seeds = append(seeds, []int{unionEnd, end})
					}
					empty = false
					break
				}
			}
			if empty {
				newSeeds = append(newSeeds, []int{start, end})
			}
		}
		seeds = newSeeds
	}
	var minLoc = math.MaxInt
	for _, couple := range seeds {
		if couple[0] < minLoc {
			minLoc = couple[0]
		}
	}
	return minLoc
}
