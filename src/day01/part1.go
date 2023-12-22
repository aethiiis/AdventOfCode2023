package main

import (
	"unicode"
)

func Part1(lines []string) int {
	sum := 0
	str := ""
	for i := 0; i < len(lines); i++ {
		var chiffres []int
		str = lines[i]
		for j := 0; j < len(str); j++ {
			if unicode.IsDigit(rune(str[j])) {
				chiffres = append(chiffres, int(rune(str[j]))-48)
			}
		}
		sum += chiffres[0]*10 + chiffres[len(chiffres)-1]
	}
	return sum
}
