package main

import (
	"advent_of_code_2023/src/utilities"
	"math"
	"strings"
)

func Part2(lines []string) int {
	sum := 0
	newLines := make([]string, len(lines))
	copy(newLines, lines)

	var emptyLine []int
	for i := range newLines {
		diese := false
		for _, char := range newLines[i] {
			if char == '#' {
				diese = true
			}
		}
		if !diese {
			newLines[i] = strings.Repeat("*", len(lines[i]))
			emptyLine = append(emptyLine, i)
		}
	}

	newLines = utilities.Transpose(newLines)

	var emptyColumn []int
	for i := range newLines {
		diese := false
		for _, char := range newLines[i] {
			if char == '#' {
				diese = true
			}
		}
		if !diese {
			newLines[i] = strings.Repeat("*", len(lines[i]))
			emptyColumn = append(emptyColumn, i)
		}
	}

	newLines = utilities.Transpose(newLines)

	var galax [][]int
	tab := utilities.ConvertirListeStringsTableauRunes(newLines)
	for i := range tab {
		for j := range tab[i] {
			if tab[i][j] == '#' {
				galax = append(galax, []int{i, j})
			}
		}
	}

	for index, point := range galax {
		for otherIndex, otherPoint := range galax {
			if index < otherIndex {

				// Check ligne
				var nbLine int
				for i := int(math.Min(float64(point[0]), float64(otherPoint[0]))); i <= int(math.Max(float64(point[0]), float64(otherPoint[0]))); i++ {
					if utilities.IsIn(emptyLine, i) {
						nbLine++
					}
				}

				// Check colonne
				var nbCol int
				for i := int(math.Min(float64(point[1]), float64(otherPoint[1]))); i <= int(math.Max(float64(point[1]), float64(otherPoint[1]))); i++ {
					if utilities.IsIn(emptyColumn, i) {
						nbCol++
					}
				}

				plus := int(math.Abs(float64(otherPoint[0]-point[0]))) + int(math.Abs(float64(otherPoint[1]-point[1]))) + (1000000-1)*nbLine + (1000000-1)*nbCol
				sum += plus
			}
		}
	}
	return sum
}
