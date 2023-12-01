package day1

import (
	"advent_of_code_2023/src/utilities"
	"unicode"
)

func Day1() {
	lines := utilities.ReadLines("input_day1")
	sum := 0
	str := ""
	for i := 0; i < len(lines); i++ {
		var chiffres []int
		str = lines[i]
		for j := 0; j < len(lines); j++ {
			if unicode.IsDigit(rune(str[j])) {
				chiffres = append(chiffres, int(rune(str[j])))
			}
		}
		sum += chiffres[0] + chiffres[len(chiffres)-1]
	}
	print(sum)
}
