package utilities

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Erreur dans l'ouverture")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Erreur dans la fermeture")
		}
	}(file)

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
