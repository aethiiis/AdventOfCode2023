package main

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"time"
)

func main() {
	// Démarrage du chrono
	debut := time.Now()

	// On charge le fichier et le divise ligne par ligne
	pathTest := "src/day02/input_day2_test"
	linesTest := utilities.ReadLines(pathTest)

	path := "src/day02/input_day2"
	lines := utilities.ReadLines(path)

	// Part1
	loc1Test := Part1(linesTest)
	fmt.Printf("La somme des IDs pour la partie 1 avec l'exemple est %d.\n", loc1Test)
	loc1 := Part1(lines)
	fmt.Printf("La somme des IDs pour la partie 1 est %d.\n", loc1)

	// Part2
	loc2Test := Part2(linesTest)
	fmt.Printf("La somme de la puissance pour la partie 2 avec l'exemple est %d.\n", loc2Test)
	loc2 := Part2(lines)
	fmt.Printf("La somme de la puissance pour la partie 2 est %d.\n", loc2)

	fin := time.Now()
	duration := fin.Sub(debut)
	fmt.Printf("La fonction a pris %s pour s'exécuter.\n", duration)
}
