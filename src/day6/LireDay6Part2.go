package day6

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"os"
	"strings"
)

func LireDay6Part2(filename string) []int {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = path + "/src/input/" + filename
	lines := utilities.ReadLines(path)
	tab := make([]int, 2)

	for i := range lines {
		line := strings.SplitN(lines[i], ":", 2)

		for range line[i] {
			line[1] = strings.Replace(line[1], "  ", " ", -1)
		}

		line = strings.SplitN(line[1], " ", -1)
		line = line[1:]
		nombre := utilities.ConcatenerNombres(line)
		tab[i] = nombre
	}
	fmt.Printf("tab : %v\n", tab)
	return tab
}
