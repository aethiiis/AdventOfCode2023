package day1

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
	"os"
	"time"
	"unicode"
)

func Day1() {
	debut := time.Now()
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = path + "\\src\\utilities\\input_day1"
	lines := utilities.ReadLines(path)
	sum := 0
	str := ""
	for i := 0; i < len(lines); i++ {
		var chiffres []int
		str = lines[i]
		j := 0
		for j < len(str) {
			if unicode.IsDigit(rune(str[j])) {
				chiffres = append(chiffres, int(rune(str[j]))-48)
				//j++
			} else if j+3 <= len(str) && str[j:j+3] == "one" {
				//j += 3
				chiffres = append(chiffres, 1)
			} else if j+3 <= len(str) && str[j:j+3] == "two" {
				//j += 3
				chiffres = append(chiffres, 2)
			} else if j+5 <= len(str) && str[j:j+5] == "three" {
				//j += 5
				chiffres = append(chiffres, 3)
			} else if j+4 <= len(str) && str[j:j+4] == "four" {
				//j += 4
				chiffres = append(chiffres, 4)
			} else if j+4 <= len(str) && str[j:j+4] == "five" {
				//j += 4
				chiffres = append(chiffres, 5)
			} else if j+3 <= len(str) && str[j:j+3] == "six" {
				//j += 3
				chiffres = append(chiffres, 6)
			} else if j+5 <= len(str) && str[j:j+5] == "seven" {
				//j += 5
				chiffres = append(chiffres, 7)
			} else if j+5 <= len(str) && str[j:j+5] == "eight" {
				//j += 5
				chiffres = append(chiffres, 8)
			} else if j+4 <= len(str) && str[j:j+4] == "nine" {
				//j += 4
				chiffres = append(chiffres, 9)
			} else if j+4 <= len(str) && str[j:j+4] == "zero" {
				//j += 4
				chiffres = append(chiffres, 0)
			}
			j++
		}
		sum += chiffres[0]*10 + chiffres[len(chiffres)-1]
	}
	fmt.Printf("La somme est égal à %d.\n", sum)
	fin := time.Now()
	duree := fin.Sub(debut)
	fmt.Printf("La fonction a pris %s pour s'exécuter.\n", duree)
}
