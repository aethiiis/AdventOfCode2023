package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	sum := 0
	for i := range lines {
		// Traitement de l'input
		springsNumber := strings.SplitN(lines[i], " ", 2)
		springs := springsNumber[0]
		numberString := springsNumber[1]

		numbersList := strings.SplitN(numberString, ",", (len(numberString)/2)+1)
		var numbers []int

		numHashtagInTotal := 0
		for _, value := range numbersList {
			num, err := strconv.Atoi(value)
			if err != nil {
				panic("Erreur de conversion string --> int")
			} else {
				numbers = append(numbers, num)
				numHashtagInTotal += num
			}
		}
		numPointInTotal := len(springs) - numHashtagInTotal
		listPermut := generateStrings(numPointInTotal, numHashtagInTotal, "")
		var listPermutAutorisees []string

		for i := range listPermut {
			if checkConformity(listPermut[i], springs, numbers) {
				sum++
				listPermutAutorisees = append(listPermutAutorisees, listPermut[i])
			}
		}

		/*fmt.Printf("Voici la liste de toutes les permutations possibles pour %s : \n", springs)
		for i := range listPermutAutorisees {
			fmt.Println(listPermutAutorisees[i])
		}*/
	}
	return sum
}

func generateStrings(n, m int, current string) []string {
	var result []string

	if n == 0 && m == 0 {
		return []string{current}
	}

	if n > 0 {
		result = append(result, generateStrings(n-1, m, current+".")...)
	}

	if m > 0 {
		result = append(result, generateStrings(n, m-1, current+"#")...)
	}

	return result
}

func checkConformity(possible string, control string, numbers []int) bool {
	for i := range control {
		if control[i] != possible[i] && control[i] != '?' {
			return false
		}
	}
	currentIndex := 0
	for i := range numbers {
		var index int
		//fmt.Printf("Pour la string %s, on vérifie si %s match à l'indice %d\n", control, possible[currentIndex:], currentIndex)
		if i == len(numbers)-1 {
			index = strings.Index(possible[currentIndex:], strings.Repeat("#", numbers[i]))
		} else {
			index = strings.Index(possible[currentIndex:], strings.Repeat("#", numbers[i])+".")
		}

		if index == -1 {
			return false
		}

		currentIndex += index + numbers[i]
	}
	return true
}

func generate(n int, A []int) {
	c := make([]int, n)

	output := func(arr []int) {
		fmt.Println(arr)
	}

	output(A)

	i := 1
	for i < n {
		if c[i] < i {
			if i%2 == 0 {
				A[0], A[i] = A[i], A[0]
			} else {
				A[c[i]], A[i] = A[i], A[c[i]]
			}
			output(A)
			c[i]++
			i = 1
		} else {
			c[i] = 0
			i++
		}
	}
}
