package main

import (
	"advent_of_code_2023/src/utilities"
	"math"
	"strings"
)

func Part2(lines []string) int {
	tab := GetTabPart2(lines)
	delta := int(math.Pow(float64(tab[0]), 2)) - 4*tab[1]
	if delta == 0 {
		return 1
	} else if delta > 0 {
		x1 := int((-float64(tab[0]) + math.Sqrt(float64(delta))) / 2)
		x2 := int((-float64(tab[0])-math.Sqrt(float64(delta)))/2) + 1
		return x1 - x2 + 1
	} else {
		return 0
	}
}

func GetTabPart2(lines []string) []int {
	tab := make([]int, 2)

	for i := range lines {
		line := strings.SplitN(lines[i], ":", 2)

		for range line[i] {
			line[1] = strings.Replace(line[1], "  ", " ", -1)
		}

		line = strings.SplitN(line[1], " ", -1)
		line = line[1:]
		nombre := utilities.ConcatenerNombres(line)
		tab[i] = nombre
	}
	return tab
}
