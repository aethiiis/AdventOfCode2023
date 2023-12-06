package day6

import (
	"fmt"
)

func Part1(filename string) int {
	tab := LireDay6(filename)
	mult := 1
	for _, race := range tab {
		count := 0
		for j := 1; j < race[0]; j++ {
			distance := j * (race[0] - j)
			if distance > race[1] {
				count++
			}
		}
		mult *= count
	}

	fmt.Printf("Le mult est : %d\n", mult)
	return mult
}
