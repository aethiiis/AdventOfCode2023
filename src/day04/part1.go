package main

import (
	"advent_of_code_2023/src/utilities"
)

func Part1(lines []string) int {
	var listWinners [][]int
	var listNumbers [][]int
	listWinners, listNumbers = utilities.SeparateWinnersNumbers(lines)

	points := 0
	for i := 0; i < len(listNumbers); i++ {
		sum := 0
		nb := 0
		for j := range listNumbers[i] {
			if utilities.IsIn(listWinners[i], listNumbers[i][j]) {
				if sum == 0 {
					sum = 1
				} else {
					sum *= 2
				}
				nb++
			}
		}
		points += sum
	}
	return points
}
