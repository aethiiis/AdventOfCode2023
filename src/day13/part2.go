package main

import (
	"advent_of_code_2023/src/utilities"
)

func Part2(lines []string) int {
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
	// Partie 1
	var lignePart1 []int
	var colonnePart1 []int
	lignePart1 = make([]int, len(blocs))
	colonnePart1 = make([]int, len(blocs))
	for i := range lignePart1 {
		lignePart1[i] = -1
		colonnePart1[i] = -1
	}

	for i, bloc := range blocs {
		blockSum := 0
		symmetry := false
		for j := range bloc {
			blockSum += checkSymmetry(bloc, j)
			if blockSum != 0 {
				lignePart1[i] = j
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
					colonnePart1[i] = j
					symmetry = true
					break
				}
			}
		}
	}

	// Partie 2
	sum := 0
	for i, bloc := range blocs {
		blockSum := 0
		blockSum += possibleModif(bloc, lignePart1[i])
		if blockSum != 0 {
			sum += blockSum * 100
		} else {
			blocTrans := utilities.Transpose(bloc)
			blockSum += possibleModif(blocTrans, colonnePart1[i])
			if blockSum != 0 {
				sum += blockSum
			}
		}
	}
	return sum
}

func possibleModif(bloc []string, exclusion int) int {
	var sum int
	var breakFor bool
	breakFor = false

	for i := range bloc {
		for j := range bloc {
			if ok, index := utilities.StringsIdentiquesAUnCaracterePres(bloc[i], bloc[j]); ok && i < j && (i+j)%2 == 1 {
				bloc[i] = inverseChar(bloc[i], index)
				for k := range bloc {
					if k != exclusion {
						sum = checkSymmetry(bloc, k)
						if sum != 0 {
							break
						}
					}
				}
				if sum == 0 {
					bloc[i] = inverseChar(bloc[i], index)
				} else {
					breakFor = true
					break
				}
			}
		}
		if breakFor {
			break
		}
	}
	return sum
}

func inverseChar(line string, index int) string {
	var newLine string
	if line[index] == '#' {
		newLine = line[0:index] + "." + line[index+1:]
	} else if line[index] == '.' {
		newLine = line[0:index] + "#" + line[index+1:]
	} else {
		panic("Caractère non-reconnu")
	}
	return newLine
}
