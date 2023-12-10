package main

import (
	"advent_of_code_2023/src/utilities"
)

func Part1(lines []string) int {
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

	firstWay := 0
	firstPipe := make([]int, 2)

	secondWay := 0
	secondPipe := make([]int, 2)

	firstPipeOK := false
	runeTab := utilities.ConvertirListeStringsTableauRunes(lines)

	for i := range connection {
		for j := range connection[i] {
			if connection[i][j] == 1 && !firstPipeOK {
				firstPipe[0] = i + xPos - 1
				firstPipe[1] = j + yPos - 1
				firstWay++
				firstPipeOK = true
			} else if connection[i][j] == 1 && firstPipeOK {
				secondPipe[0] = i + xPos - 1
				secondPipe[1] = j + yPos - 1
				secondWay++
				runeTab[xPos][yPos] = '*'
			}
		}
	}

	for firstPipe[0] != secondPipe[0] || firstPipe[1] != secondPipe[1] {
		// First Way
		firstPipe = utilities.NextDirection(firstPipe, runeTab)
		firstWay++

		// Second Way
		secondPipe = utilities.NextDirection(secondPipe, runeTab)
		secondWay++
	}

	if firstWay != secondWay {
		panic("Problème dans le comptage des étapes")
	}

	return firstWay
}
