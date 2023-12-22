package main

import (
	"unicode"
)

func Part2(lines []string) int {
	sum := 0
	str := ""
	for i := 0; i < len(lines); i++ {
		var chiffres []int
		str = lines[i]
		for j := 0; j < len(str); j++ {
			if unicode.IsDigit(rune(str[j])) {
				chiffres = append(chiffres, int(rune(str[j]))-48)
			} else if j+3 <= len(str) && str[j:j+3] == "one" {
				chiffres = append(chiffres, 1)
			} else if j+3 <= len(str) && str[j:j+3] == "two" {
				chiffres = append(chiffres, 2)
			} else if j+5 <= len(str) && str[j:j+5] == "three" {
				chiffres = append(chiffres, 3)
			} else if j+4 <= len(str) && str[j:j+4] == "four" {
				chiffres = append(chiffres, 4)
			} else if j+4 <= len(str) && str[j:j+4] == "five" {
				chiffres = append(chiffres, 5)
			} else if j+3 <= len(str) && str[j:j+3] == "six" {
				chiffres = append(chiffres, 6)
			} else if j+5 <= len(str) && str[j:j+5] == "seven" {
				chiffres = append(chiffres, 7)
			} else if j+5 <= len(str) && str[j:j+5] == "eight" {
				chiffres = append(chiffres, 8)
			} else if j+4 <= len(str) && str[j:j+4] == "nine" {
				chiffres = append(chiffres, 9)
			} else if j+4 <= len(str) && str[j:j+4] == "zero" {
				chiffres = append(chiffres, 0)
			}
		}
		sum += chiffres[0]*10 + chiffres[len(chiffres)-1]
	}
	return sum
}
