package main

import (
	"math/big"
	"strings"
)

func Part2(lines []string) *big.Int {
	instructions, mapLoc := getMap(lines)

	var startingNodes []string
	for cle := range mapLoc {
		if strings.HasSuffix(cle, "A") {
			startingNodes = append(startingNodes, cle)
		}
	}

	index := 0

	countNodes := make([]int64, len(startingNodes))

	for i := range startingNodes {
		for !strings.HasSuffix(startingNodes[i], "Z") {
			instruction := instructions[index]
			if index == len(instructions)-1 {
				index = 0
			} else {
				index++
			}

			startingNodes[i] = mapLoc[startingNodes[i]][instruction]

			countNodes[i]++
		}
	}

	return ppcmMultiples(countNodes)
}

func pgcdMultiples(nombres []int64) *big.Int {
	if len(nombres) < 2 {
		panic("La fonction pgcdMultiples nécessite au moins deux nombres.")
	}

	resultat := big.NewInt(nombres[0])

	for _, nombre := range nombres[1:] {
		resultat.GCD(nil, nil, resultat, big.NewInt(nombre))
	}

	return resultat
}

func ppcmMultiples(nombres []int64) *big.Int {
	if len(nombres) < 2 {
		panic("La fonction ppcmMultiples nécessite au moins deux nombres.")
	}

	resultat := new(big.Int).Set(big.NewInt(nombres[0]))

	for _, nombre := range nombres[1:] {
		resultat.Mul(resultat, big.NewInt(nombre))
		resultat.Div(resultat, pgcdMultiples(nombres))
	}

	return resultat
}
