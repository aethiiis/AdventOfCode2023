package day4

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"os"
	"time"
)

func Day4() {
	debut := time.Now()
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = path + "/src/utilities/input_day4"
	lines := utilities.ReadLines(path)
	var listWinners [][]int
	var listNumbers [][]int
	listWinners, listNumbers = utilities.SeparateWinnersNumbers(lines)

	originalWinners := listWinners
	originalNumbers := listNumbers
	numCard := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		numCard[i] = i
	}
	points := 0
	i := 0
	for i < len(listNumbers) {
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
		for j := 1; j <= nb; j++ {
			numCard = append(numCard, numCard[i]+j)

			listWinners = append(listWinners, originalWinners[numCard[i]+j])

			listNumbers = append(listNumbers, originalNumbers[numCard[i]+j])
		}
		i++
	}

	fmt.Printf("La somme est égale à : %d.\n", points)
	fmt.Printf("Au final, on a %d cartes.\n", len(listNumbers))
	fin := time.Now()
	duree := fin.Sub(debut)
	fmt.Printf("La fonction a pris %s pour s'exécuter.\n", duree)
}
