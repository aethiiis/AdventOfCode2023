package utilities

import (
	"fmt"
	"regexp"
	"strconv"
)

func DetectNumbersSymbols(lines []string) ([]int, []int) {
	expression, err := regexp.Compile(`\d+`)
	var list []int
	carteGear := make([][]int, len(lines))
	for i := range carteGear {
		carteGear[i] = make([]int, len(lines[0]))
	}
	if err != nil {
		fmt.Println("Erreur de compilation de l'expression régulière :", err)
		return nil, nil
	}
	for i := 0; i < len(lines); i++ {
		positions := expression.FindAllStringIndex(lines[i], -1)
		for _, position := range positions {
			if SymbolSensor(position, i, lines, carteGear) {
				temp, err := strconv.Atoi(lines[i][position[0]:position[1]])
				if err != nil {
					fmt.Println("Erreur de conversion string --> entier :", err)
					return nil, nil
				}
				list = append(list, temp)
			}
		}
	}
	var groupes [][][2]int

	// Fonction pour vérifier si un indice est à côté de 2
	aCoteDeux := func(i, j int) (bool, [2]int) {
		for x := i - 1; x <= i+1; x++ {
			for y := j - 1; y <= j+1; y++ {
				if x >= 0 && x < len(carteGear) && y >= 0 && y < len(carteGear[0]) && carteGear[x][y] == 2 {
					return true, [2]int{x, y}
				}
			}
		}
		return false, [2]int{-1, -1}
	}

	for i := 0; i < len(carteGear); i++ {
		for j := 0; j < len(carteGear[0]); j++ {
			test2, pos2 := aCoteDeux(i, j)
			if carteGear[i][j] == -1 && test2 {
				// Trouver le début d'un groupe de -1 à côté de 2...
				debut := [2]int{i, j}
				y := j
				for y >= 0 && carteGear[i][y] == -1 {
					carteGear[i][y] = 0
					debut[1] = y
					y--
				}

				//...chercher la fin
				fin := [2]int{i, j}
				y = j + 1
				compteur := 0
				for y < len(carteGear[i]) && carteGear[i][y] == -1 {
					carteGear[i][y] = 0
					fin[1] = y
					compteur++
					y++
				}

				// Ajouter le groupe à la liste
				groupes = append(groupes, [][2]int{debut, fin})
				j += compteur

				// Recherche du second nombre
				x2 := pos2[0]
				y2 := pos2[1]
				for i2 := x2 - 1; i2 <= x2+1; i2++ {
					for j2 := y2 - 1; j2 <= y2+1; j2++ {
						if i2 >= 0 && i2 < len(carteGear) && j2 >= 0 && j2 < len(carteGear[0]) && carteGear[i2][j2] == -1 {
							// Trouver le début d'un groupe de -1 à côté de 2...
							debut := [2]int{i2, j2}
							yy := j2
							for yy >= 0 && carteGear[i2][yy] == -1 {
								carteGear[i2][yy] = 0
								debut[1] = yy
								yy--
							}

							//...chercher la fin
							fin := [2]int{i2, j2}
							yy = j2 + 1
							for yy < len(carteGear[i2]) && carteGear[i2][yy] == -1 {
								carteGear[i2][yy] = 0
								fin[1] = yy
								yy++
							}

							// Ajouter le groupe à la liste
							groupes = append(groupes, [][2]int{debut, fin})
						}
					}
				}
			}
		}
	}
	var groupesInt []int
	for i := 0; i < len(groupes); i++ {
		ligne := groupes[i][0][0]
		debut := groupes[i][0][1]
		fin := groupes[i][1][1]
		nombre, err := strconv.Atoi(lines[ligne][debut : fin+1])
		if err != nil {
			fmt.Println("Erreur de conversion string --> entier :", err)
			return nil, nil
		}
		groupesInt = append(groupesInt, nombre)
	}
	return list, groupesInt
}
