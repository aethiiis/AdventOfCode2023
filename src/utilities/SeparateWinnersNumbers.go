package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

func SeparateWinnersNumbers(lines []string) ([][]int, [][]int) {

	listWinners := make([][]int, len(lines))
	listNumbers := make([][]int, len(lines))

	for i := range lines {

		lines[i] = strings.Replace(lines[i], "  ", " ", -1)

		split := strings.SplitN(lines[i], ":", 2)

		winnersNumbers := strings.SplitN(split[1], "|", 2)

		//fmt.Printf("Voici winnersNumbers %v\n", winnersNumbers)

		lenWinners := CompterNombreNombreInString(winnersNumbers[0])
		lenNumbers := CompterNombreNombreInString(winnersNumbers[1])

		//fmt.Printf("lenWinners : %d\n", lenWinners)
		//fmt.Printf("lenNumbers : %d\n", lenNumbers)

		winners := strings.SplitN(winnersNumbers[0], " ", lenWinners+1)
		numbers := strings.SplitN(winnersNumbers[1], " ", lenNumbers+1)

		winners = winners[1:]
		numbers = numbers[1:]

		listWinners[i] = make([]int, lenWinners)
		listNumbers[i] = make([]int, lenNumbers)

		for j := range winners {
			chiffre, err := strconv.Atoi(strings.Trim(winners[j], " "))
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
				return nil, nil
			}
			listWinners[i][j] = chiffre
		}

		for j := range numbers {
			chiffre, err := strconv.Atoi(strings.Trim(numbers[j], " "))
			if err != nil {
				fmt.Println("Erreur de conversion :", err)
				return nil, nil
			}
			listNumbers[i][j] = chiffre
		}
	}
	return listWinners, listNumbers
}
