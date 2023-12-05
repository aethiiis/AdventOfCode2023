package day5

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"os"
	"time"
)

func Day5() {
	debut := time.Now()
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = path + "/src/input/input_day5"
	lines := utilities.ReadLines(path)

	fin := time.Now()
	duree := fin.Sub(debut)
	fmt.Printf("La fonction a pris %s pour s'exécuter.\n", duree)
}
