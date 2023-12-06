package day6

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LireDay6(filename string) [][]int {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = path + "/src/input/" + filename
	lines := utilities.ReadLines(path)

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
			fmt.Printf("Le chiffre considéré est : %d\n", chiffre)
			tab[j][i] = chiffre
		}
	}
	fmt.Printf("tab : %v\n", tab)
	return tab
}
