package main

import (
	"advent_of_code_2023/src/utilities"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
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
				gauche := tab[i][j][len(tab[i][j])-1]
				bas := tab[i][j+1][len(tab[i][j+1])-1]
				tab[i][j] = append(tab[i][j], gauche+bas)
				if j == 0 {
					sum += gauche + bas
				}
			}
		}
	}
	return sum
}

func GenerateLists(lines []string) [][][]int {
	tab := make([][][]int, len(lines))
	for i := range tab {
		tab[i] = make([][]int, 1)
		for j := range tab[i] {
			tab[i][j] = make([]int, 1)
		}
	}

	for i := range lines {
		nombreNombre := utilities.CompterNombreNombreInString(lines[i])
		nombresString := strings.SplitN(lines[i], " ", nombreNombre)

		var nombreInt []int
		for j := range nombresString {
			chiffre, err := strconv.Atoi(nombresString[j])
			if err != nil {
				panic("Erreur de conversion string -> int\n")
			} else {
				nombreInt = append(nombreInt, chiffre)
			}
		}
		tab[i][0] = nombreInt
	}
	return tab
}
