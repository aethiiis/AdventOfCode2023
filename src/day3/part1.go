package main

import (
	"advent_of_code_2023/src/utilities"
)

func Part1(lines []string) int {
	listSum, _ := utilities.DetectNumbersSymbols(lines)
	sum := utilities.SumList(listSum)
	return sum
}
