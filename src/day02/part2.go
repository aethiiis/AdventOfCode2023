package main

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"strconv"
	"strings"
)

func Part2(lines []string) int {
	resultat := utilities.MapMaker(lines)
	pow := 0

	for _, grandeList := range resultat {
		red := 0
		green := 0
		blue := 0
		for _, petiteList := range grandeList {
			for _, valeur := range petiteList {
				split := strings.SplitN(valeur, " ", 2)
				chiffre, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println("Erreur de conversion :", err)
				}
				couleur := split[1]
				if couleur == "red" && chiffre > red {
					red = chiffre
				} else if couleur == "green" && chiffre > green {
					green = chiffre
				} else if couleur == "blue" && chiffre > blue {
					blue = chiffre
				}
			}
		}
		pow += red * green * blue
	}
	return pow
}
