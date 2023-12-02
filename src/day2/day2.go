package day2

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day2() {
	debut := time.Now()
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = path + "\\src\\utilities\\input_day2"
	lines := utilities.ReadLines(path)
	resultat := utilities.MapMaker(lines)
	sum := 0
	pow := 0

	for key, grandeList := range resultat {
		possible := true
		red := 0
		green := 0
		blue := 0
		for _, petiteList := range grandeList {
			for _, valeur := range petiteList {
				split := strings.SplitN(valeur, " ", 2)
				chiffre, err := strconv.Atoi(split[0])
				if err != nil {
					fmt.Println("Erreur de conversion :", err)
					return
				}
				couleur := split[1]
				if couleur == "red" && chiffre > red {
					red = chiffre
				} else if couleur == "green" && chiffre > green {
					green = chiffre
				} else if couleur == "blue" && chiffre > blue {
					blue = chiffre
				}
				if (chiffre > 12 && couleur == "red") || (chiffre > 13 && couleur == "green") || (chiffre > 14 && couleur == "blue") {
					possible = false
				}
			}
		}
		pow += red * green * blue
		if possible {
			parts := strings.SplitN(key, " ", 2)
			temp, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
				return
			}
			sum += temp
		}
	}
	fmt.Printf("La somme des games possibles est : %d.\n", sum)
	fmt.Printf("La puissance des games est : %d.\n", pow)
	fin := time.Now()
	duree := fin.Sub(debut)
	fmt.Printf("La fonction a pris %s pour s'exécuter.\n", duree)
}
