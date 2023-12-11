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
	pathTest1 := "src/day10/input_day10_test1"
	linesTest1 := utilities.ReadLines(pathTest1)

	pathTest2 := "src/day10/input_day10_test2"
	linesTest2 := utilities.ReadLines(pathTest2)

	pathTest3 := "src/day10/input_day10_test3"
	linesTest3 := utilities.ReadLines(pathTest3)

	pathTest4 := "src/day10/input_day10_test4"
	linesTest4 := utilities.ReadLines(pathTest4)

	pathTest5 := "src/day10/input_day10_test5"
	linesTest5 := utilities.ReadLines(pathTest5)

	pathTest6 := "src/day10/input_day10_test6"
	linesTest6 := utilities.ReadLines(pathTest6)

	path := "src/day10/input_day10"
	lines := utilities.ReadLines(path)

	// Part1
	loc1Test1 := Part1(linesTest1)
	fmt.Printf("La distance pour la partie 1 avec l'exemple 1 est %d.\n", loc1Test1)
	loc1Test2 := Part1(linesTest2)
	fmt.Printf("La distance pour la partie 1 avec l'exemple 2 est %d.\n", loc1Test2)
	loc1 := Part1(lines)
	fmt.Printf("La distance pour la partie 1 est %d.\n", loc1)

	// Part2
	loc2Test3 := Part2(linesTest3)
	fmt.Printf("Les gains pour la partie 2 avec l'exemple 3 est %d.\n", loc2Test3)
	loc2Test4 := Part2(linesTest4)
	fmt.Printf("Les gains pour la partie 2 avec l'exemple 4 est %d.\n", loc2Test4)
	loc2Test5 := Part2(linesTest5)
	fmt.Printf("Les gains pour la partie 2 avec l'exemple 5 est %d.\n", loc2Test5)
	loc2Test6 := Part2(linesTest6)
	fmt.Printf("Les gains pour la partie 2 avec l'exemple 6 est %d.\n", loc2Test6)
	loc2 := Part2(lines)
	fmt.Printf("Les gains pour la partie 2 est %d.\n", loc2)

	fin := time.Now()
	duration := fin.Sub(debut)
	fmt.Printf("La fonction a pris %s pour s'exécuter.\n", duration)
}
