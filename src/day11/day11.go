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
	pathTest1 := "src/day11/input_day11_test"
	linesTest1 := utilities.ReadLines(pathTest1)

	path := "src/day11/input_day11"
	lines := utilities.ReadLines(path)

	// Part1
	loc1Test1 := Part1(linesTest1)
	fmt.Printf("La distance pour la partie 1 avec l'exemple 1 est %d.\n", loc1Test1)
	loc1 := Part1(lines)
	fmt.Printf("La distance pour la partie 1 est %d.\n", loc1)

	// Part2
	loc2Test1 := Part2(linesTest1)
	fmt.Printf("La distance pour la partie 2 avec l'exemple 1 est %d.\n", loc2Test1)
	loc2 := Part2(lines)
	fmt.Printf("Les gains pour la partie 2 est %d.\n", loc2)

	fin := time.Now()
	duration := fin.Sub(debut)
	fmt.Printf("La fonction a pris %s pour s'exécuter.\n", duration)
}
