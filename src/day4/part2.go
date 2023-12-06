package main

import (
	"advent_of_code_2023/src/utilities"
)

func Part2(lines []string) int {
	var listWinners [][]int
	var listNumbers [][]int
	listWinners, listNumbers = utilities.SeparateWinnersNumbers(lines)

	originalWinners := listWinners
	originalNumbers := listNumbers
	numCard := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		numCard[i] = i
	}
	i := 0
	for i < len(listNumbers) {
		nb := 0
		for j := range listNumbers[i] {
			if utilities.IsIn(listWinners[i], listNumbers[i][j]) {
				nb++
			}
		}
		for j := 1; j <= nb; j++ {
			numCard = append(numCard, numCard[i]+j)

			listWinners = append(listWinners, originalWinners[numCard[i]+j])

			listNumbers = append(listNumbers, originalNumbers[numCard[i]+j])
		}
		i++
	}
	return len(listNumbers)
}
