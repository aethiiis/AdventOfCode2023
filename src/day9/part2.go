package main

import "advent_of_code_2023/src/utilities"

func Part2(lines []string) int {
	tab := GenerateLists(lines)

	// Génération du tableau
	for i := range tab {
		currentTab := tab[i][0]
		for !utilities.EstVide(currentTab) {
			newTab := make([]int, len(currentTab)-1)
			for j := range newTab {
				newTab[j] = currentTab[j+1] - currentTab[j]
			}
			tab[i] = append(tab[i], newTab)
			currentTab = newTab
		}
	}
	sum := 0
	// Extrapolation du tableau
	for i := range tab {
		for j := len(tab[i]) - 1; j > -1; j-- {
			if utilities.EstVide(tab[i][j]) {
				tab[i][j] = append(tab[i][j], 0)
			} else {
				droite := tab[i][j][0]
				bas := tab[i][j+1][0]
				tab[i][j] = append([]int{droite - bas}, tab[i][j]...)
				if j == 0 {
					sum += droite - bas
				}
			}
		}
	}
	return sum
}
