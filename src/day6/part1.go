package main

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	tab := GetTabPart1(lines)
	mult := 1
	for _, race := range tab {
		count := 0
		for j := 1; j < race[0]; j++ {
			distance := j * (race[0] - j)
			if distance > race[1] {
				count++
			}
		}
		mult *= count
	}
	return mult
}

func GetTabPart1(lines []string) [][]int {
	nbRace := utilities.CompterNombreNombreInString(lines[0][5:])
	tab := make([][]int, nbRace)

	for i := range tab {
		tab[i] = make([]int, 2)
	}

	for i := range lines {
		line := strings.SplitN(lines[i], ":", 2)

		for range line[i] {
			line[1] = strings.Replace(line[1], "  ", " ", -1)
		}

		line = strings.SplitN(line[1], " ", -1)
		line = line[1:]

		for j := range line {
			chiffre, err := strconv.Atoi(strings.Trim(line[j], " "))
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
			}
			tab[j][i] = chiffre
		}
	}
	return tab
}
