package utilities

import (
	"strings"
)

func DoubleEmptyLineColumn(lines []string, filler string) []string {
	newLines := make([]string, len(lines))
	copy(newLines, lines)
	count := 0
	for i := range lines {
		diese := false
		for _, char := range lines[i] {
			if char == '#' {
				diese = true
			}
		}
		if !diese {
			nouvelleLigne := strings.Repeat(filler, len(lines[i]))
			newLines = append(newLines[:count], append([]string{nouvelleLigne}, newLines[count:]...)...)
			count++
		}
		count++
	}

	transLines := Transpose(newLines)

	newTransLines := make([]string, len(transLines))
	copy(newTransLines, transLines)
	count = 0
	for i := range transLines {
		diese := false
		for _, char := range transLines[i] {
			if char == '#' {
				diese = true
			}
		}
		if !diese {
			nouvelleLigne := strings.Repeat(filler, len(transLines[i]))
			newTransLines = append(newTransLines[:count], append([]string{nouvelleLigne}, newTransLines[count:]...)...)
			count++
		}
		count++
	}
	/*	fmt.Println("Voici newLines : ")
		for i := range newLines {
			fmt.Println(newLines[i])
		}

		fmt.Println("Voici transLines : ")
		for i := range newTransLines {
			fmt.Println(newTransLines[i])
		}*/
	return Transpose(newTransLines)
}
