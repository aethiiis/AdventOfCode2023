package main

import (
	"advent_of_code_2023/src/utilities"
)

func Part2(lines []string) int {
	_, listFactor := utilities.DetectNumbersSymbols(lines)
	factor := utilities.Sum2ProductList(listFactor)
	return factor
}
