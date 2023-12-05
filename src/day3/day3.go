package day3

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"os"
	"time"
)

func Day3() {
	debut := time.Now()
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = path + "/src/input/input_day3"
	lines := utilities.ReadLines(path)
	listSum, listFactor := utilities.DetectNumbersSymbols(lines)
	sum := utilities.SumList(listSum)
	factor := utilities.Sum2ProductList(listFactor)
	fmt.Printf("La somme des nombres concernés est : %d.\n", sum)
	fmt.Printf("La somme des facteurs est égale à %d.\n", factor)
	fin := time.Now()
	duree := fin.Sub(debut)
	fmt.Printf("La fonction a pris %s pour s'exécuter.\n", duree)
}
