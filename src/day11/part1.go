package main

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"math"
)

func Part1(lines []string) int {
	var galax [][]int
	//dist := 0
	newLines := utilities.DoubleEmptyLineColumn(lines, ".")
	fmt.Println("Voici newLines : ")
	for i := range newLines {
		fmt.Println(newLines[i])
	}

	tab := utilities.ConvertirListeStringsTableauRunes(newLines)
	for i := range tab {
		for j := range tab[i] {
			if tab[i][j] == '#' {
				galax = append(galax, []int{i, j})
			}
		}
	}
	//fmt.Printf("Voici galax : %v\n", galax)

	sum := 0

	for index, point := range galax {
		for otherIndex, otherPoint := range galax {
			if index > otherIndex {
				sum += int(math.Abs(float64(otherPoint[0]-point[0]))) + int(math.Abs(float64(otherPoint[1]-point[1])))
				//fmt.Printf("La distance entre %d et %d est %d\n", index, otherIndex, int(math.Abs(float64(otherPoint[0]-point[0])))+int(math.Abs(float64(otherPoint[1]-point[1]))))
			}
		}
	}
	return sum
}
