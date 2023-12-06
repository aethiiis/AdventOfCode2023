package main

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	resultat := utilities.MapMaker(lines)
	sum := 0

	for key, grandeList := range resultat {
		possible := true

		for _, petiteList := range grandeList {
			for _, valeur := range petiteList {
				split := strings.SplitN(valeur, " ", 2)
				chiffre, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println("Erreur de conversion :", err)
				}
				couleur := split[1]
				if (chiffre > 12 && couleur == "red") || (chiffre > 13 && couleur == "green") || (chiffre > 14 && couleur == "blue") {
					possible = false
				}
			}
		}
		if possible {
			parts := strings.SplitN(key, " ", 2)
			temp, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
			}
			sum += temp
		}
	}
	return sum
}
