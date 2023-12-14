package main

import (
	"advent_of_code_2023/src/utilities"
)

func Part1(lines []string) int {
	var tab [][]rune
	var poids int
	poids = 0
	tab = utilities.ConvertirListeStringsTableauRunes(lines)
	for i := range tab {
		for j := range tab[i] {
			if tab[i][j] == 'O' {
				poids += len(lines) - north(&tab, i, j)
			}
		}
	}
	/*fmt.Println("Voici tab : ")
	for i := range tab {
		fmt.Println(string(tab[i]))
	}*/
	return poids
}

func north(tab *[][]rune, i, j int) int {
	if (*tab)[i][j] != 'O' {
		panic("Appel non-valide")
	}
	finalLine := i
	for k := i - 1; k >= 0; k-- {
		if (*tab)[k][j] == '.' {
			finalLine--
		} else {
			break
		}
	}
	if finalLine != i {
		(*tab)[i][j] = '.'
		(*tab)[finalLine][j] = 'O'
	}

	return finalLine
}
