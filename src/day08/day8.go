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
	pathTest1 := "src/day08/input_day8_test1"
	pathTest2 := "src/day08/input_day8_test2"
	pathTest3 := "src/day08/input_day8_test3"
	linesTest1 := utilities.ReadLines(pathTest1)
	linesTest2 := utilities.ReadLines(pathTest2)
	linesTest3 := utilities.ReadLines(pathTest3)

	path := "src/day08/input_day8"
	lines := utilities.ReadLines(path)

	// Part1
	loc1Test1 := Part1(linesTest1)
	fmt.Printf("Le nombre d'étapes pour la partie 1 avec l'exemple 1 est %d.\n", loc1Test1)
	loc1Test2 := Part1(linesTest2)
	fmt.Printf("Le nombre d'étapes pour la partie 1 avec l'exemple 1 est %d.\n", loc1Test2)
	loc1 := Part1(lines)
	fmt.Printf("Les gains pour la partie 1 est %d.\n", loc1)

	// Part2
	loc2Test := Part2(linesTest3)
	fmt.Printf("Le nombre d'étapes pour la partie 2 avec l'exemple est %d.\n", loc2Test)
	loc2 := Part2(lines)
	fmt.Printf("Les gains pour la partie 2 est %d.\n", loc2)

	fin := time.Now()
	duration := fin.Sub(debut)
	fmt.Printf("La fonction a pris %s pour s'exécuter.\n", duration)
}
