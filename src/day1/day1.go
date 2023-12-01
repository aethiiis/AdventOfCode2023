package day1

import "unicode"

func main() {
	lines := readLines("day1")
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
