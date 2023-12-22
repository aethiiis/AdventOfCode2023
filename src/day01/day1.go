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
	pathTest1 := "src/day01/input_day1_test1"
	linesTest1 := utilities.ReadLines(pathTest1)

	pathTest2 := "src/day01/input_day1_test2"
	linesTest2 := utilities.ReadLines(pathTest2)

	path := "src/day01/input_day1"
	lines := utilities.ReadLines(path)

	// Part1
	loc1Test := Part1(linesTest1)
	fmt.Printf("La somme pour la partie 1 avec l'exemple est %d.\n", loc1Test)
	loc1 := Part1(lines)
	fmt.Printf("La somme pour la partie 1 est %d.\n", loc1)

	// Part2
	loc2Test := Part2(linesTest2)
	fmt.Printf("La somme pour la partie 2 avec l'exemple est %d.\n", loc2Test)
	loc2 := Part2(lines)
	fmt.Printf("La somme pour la partie 2 est %d.\n", loc2)

	fin := time.Now()
	duration := fin.Sub(debut)
	fmt.Printf("La fonction a pris %s pour s'exécuter.\n", duration)
}
