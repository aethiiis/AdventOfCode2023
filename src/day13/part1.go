package main

import (
	"advent_of_code_2023/src/utilities"
)

func Part1(lines []string) int {
	// Découpage en blocs
	var blocs [][]string
	var emptyLines []int
	for i := range lines {
		if len(lines[i]) == 0 {
			emptyLines = append(emptyLines, i)
		}
	}

	blocs = append(blocs, lines[0:emptyLines[0]])
	for i := 0; i < len(emptyLines); i++ {
		if i == len(emptyLines)-1 {
			blocs = append(blocs, lines[emptyLines[i]+1:])
		} else {
			blocs = append(blocs, lines[emptyLines[i]+1:emptyLines[i+1]])
		}
	}

	// Symétrie
	sum := 0
	for _, bloc := range blocs {
		blockSum := 0
		symmetry := false
		for j := range bloc {
			blockSum += checkSymmetry(bloc, j)
			if blockSum != 0 {
				symmetry = true
				blockSum *= 100
				break
			}
		}

		if !symmetry {
			blocTrans := utilities.Transpose(bloc)
			for j := range blocTrans {
				blockSum += checkSymmetry(blocTrans, j)
				if blockSum != 0 {
					symmetry = true
					break
				}
			}
		}
		sum += blockSum
	}

	return sum
}

func checkSymmetry(bloc []string, j int) int {
	sum := 0
	if j < len(bloc)-1 && bloc[j] == bloc[j+1] {
		check := true
		count := 1

		for j-count >= 0 && j+count+1 < len(bloc) && check {
			firstPart := bloc[j-count : j+1]
			secondPart := utilities.FlipLine(bloc[j+1 : j+2+count])
			if len(firstPart) != len(secondPart) {
				panic("Problème de dimension")
			} else {
				for k := range firstPart {
					if firstPart[k] != secondPart[k] {
						check = false
						break
					}
				}
			}
			count++
		}
		if check {
			sum += j + 1
		}
	}
	return sum
}
