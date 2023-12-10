package main

import (
	"advent_of_code_2023/src/utilities"
)

func Part2(lines []string) int {
	xPos := 0
	yPos := 0
	found := false

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'S' {
				xPos = i
				yPos = j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	tab := utilities.CheckAround(lines, xPos, yPos)
	connection := utilities.IsConnected(tab)

	firstPipe := make([]int, 2)

	secondPipe := make([]int, 2)

	firstPipeOK := false
	runeTab := utilities.ConvertirListeStringsTableauRunes(lines)

	copieRuneTab := make([][]rune, len(runeTab))

	for i := range copieRuneTab {
		copieRuneTab[i] = make([]rune, len(runeTab[0]))
		copy(copieRuneTab[i], runeTab[i])
	}

	// Doubler le tableau
	runeTab2 := make([][]rune, 2*len(runeTab))

	for i := range runeTab2 {
		runeTab2[i] = make([]rune, 2*len(runeTab[0]))
	}

	for i := range connection {
		for j := range connection[i] {
			if connection[i][j] == 1 && !firstPipeOK {
				firstPipe[0] = i + xPos - 1
				firstPipe[1] = j + yPos - 1
				firstPipeOK = true
			} else if connection[i][j] == 1 && firstPipeOK {
				secondPipe[0] = i + xPos - 1
				secondPipe[1] = j + yPos - 1
				runeTab[xPos][yPos] = '*'
			}
		}
	}

	for firstPipe[0] != secondPipe[0] || firstPipe[1] != secondPipe[1] {
		// First Way
		firstPipe = utilities.NextDirection(firstPipe, runeTab)

		// Second Way
		secondPipe = utilities.NextDirection(secondPipe, runeTab)
	}

	// Dernier étape du chemin
	runeTab[firstPipe[0]][firstPipe[1]] = '*'

	// Recopier l'ancien tableau
	for i := range runeTab {
		for j := range runeTab[i] {
			if runeTab[i][j] == '*' || runeTab[i][j] == '.' {
				runeTab2[2*i][2*j] = copieRuneTab[i][j]
			} else {
				runeTab2[2*i][2*j] = '.'
			}
		}
	}

	// Doubler en colonne
	for i := 0; i < len(runeTab2); i += 2 {
		for j := 0; j < len(runeTab2[i]); j++ {
			if j == len(runeTab2[i])-1 && runeTab2[i][j] == 0 {
				runeTab2[i][j] = ','
			} else if runeTab2[i][j] == 0 {
				if (runeTab2[i][j-1] == 'L' || runeTab2[i][j-1] == 'F' || runeTab2[i][j-1] == '-' || runeTab2[i][j-1] == 'S') && (runeTab2[i][j+1] == 'J' || runeTab2[i][j+1] == '7' || runeTab2[i][j+1] == '-' || runeTab2[i][j+1] == 'S') {
					runeTab2[i][j] = '-'
				} else {
					runeTab2[i][j] = ','
				}
			}
		}
	}

	// Doubler en ligne
	for i := 1; i < len(runeTab2); i += 2 {
		for j := 0; j < len(runeTab2[i]); j++ {
			if i == len(runeTab2)-1 && runeTab2[i][j] == 0 {
				runeTab2[i][j] = ','
			} else if runeTab2[i][j] == 0 {
				if (runeTab2[i-1][j] == 'F' || runeTab2[i-1][j] == '7' || runeTab2[i-1][j] == '|' || runeTab2[i-1][j] == 'S') && (runeTab2[i+1][j] == 'L' || runeTab2[i+1][j] == 'J' || runeTab2[i+1][j] == '|' || runeTab2[i+1][j] == 'S') {
					runeTab2[i][j] = '|'
				} else {
					runeTab2[i][j] = ','
				}
			}
		}
	}

	// Couche extérieure
	for j := 0; j < len(runeTab2[0]); j++ {
		if runeTab2[0][j] == '.' || runeTab2[0][j] == ',' {
			runeTab2[0][j] = '0'
		}
	}

	for j := 0; j < len(runeTab2[0]); j++ {
		if runeTab2[len(runeTab2)-1][j] == '.' || runeTab2[len(runeTab2)-1][j] == ',' {
			runeTab2[len(runeTab2)-1][j] = '0'
		}
	}

	for i := 0; i < len(runeTab2); i++ {
		if runeTab2[i][0] == '.' || runeTab2[i][0] == ',' {
			runeTab2[i][0] = '0'
		}
	}

	for i := 0; i < len(runeTab2); i++ {
		if runeTab2[i][len(runeTab2[i])-1] == '.' || runeTab2[i][len(runeTab2[i])-1] == ',' {
			runeTab2[i][len(runeTab2[i])-1] = '0'
		}
	}

	for n := 0; n < 100; n++ {
		for i := range runeTab2 {
			for j := range runeTab2[i] {
				if (runeTab2[i][j] == '.' || runeTab2[i][j] == ',') && utilities.IsNeighbourToZero(runeTab2, i, j) {
					runeTab2[i][j] = '0'
				}
			}
		}
	}

	count1 := 0
	for _, liste := range runeTab2 {
		count1 += utilities.Count('.', liste)
	}
	return count1
}
