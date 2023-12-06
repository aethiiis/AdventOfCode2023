package day6

import (
	"fmt"
	"math"
)

func Part2(filename string) int {
	tab := LireDay6Part2(filename)
	delta := int(math.Pow(float64(tab[0]), 2)) - 4*tab[1]
	if delta == 0 {
		return 1
	} else if delta > 0 {
		x1 := int((-float64(tab[0]) + math.Sqrt(float64(delta))) / 2)
		x2 := int((-float64(tab[0])-math.Sqrt(float64(delta)))/2) + 1
		fmt.Printf("Les deux entiers bornes sont %d et %d : il y a donc %d solutions possibles\n", x2, x1, x1-x2+1)
		return x1 - x2 + 1
	} else {
		return 0
	}
}
