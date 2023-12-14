package main

import (
	"advent_of_code_2023/src/utilities"
	"strings"
)

func Part2(lines []string) int {
	var tab [][]rune
	tab = utilities.ConvertirListeStringsTableauRunes(lines)
	vus := make(map[string]int)
	st := ""
	for i := range tab {
		st += string(tab[i])
	}
	vus[st] = 0
	var liste [][][]rune
	temp := make([][]rune, len(tab))
	for i := range tab {
		temp[i] = make([]rune, len(tab[i]))
		copy(temp[i], tab[i])
	}
	liste = append(liste, temp)
	var iter int
	iter = 0
	var cycle bool
	cycle = true
	for cycle {
		iter++

		for i := 1; i < len(tab); i++ {
			for j := range tab[i] {
				if tab[i][j] == 'O' {
					north(&tab, i, j)
				}

			}
		}
		for j := 1; j < len(tab[0]); j++ {
			for i := range tab {
				if tab[i][j] == 'O' {
					west(&tab, i, j)
				}

			}
		}
		for i := len(tab) - 2; i >= 0; i-- {
			for j := range tab[i] {
				if tab[i][j] == 'O' {
					south(&tab, i, j)
				}

			}
		}
		for j := len(tab[0]) - 2; j >= 0; j-- {
			for i := range tab {
				if tab[i][j] == 'O' {
					east(&tab, i, j)
				}

			}
		}
		sum := 0
		for i := range tab {
			sum += strings.Count(string(tab[i]), "O") * (len(tab) - i)
		}
		str := ""
		for i := range tab {
			str += string(tab[i])
		}
		_, ok := vus[str]
		if ok {
			cycle = false
		} else {
			vus[str] = iter
			tmp := make([][]rune, len(tab))
			for i := range tab {
				tmp[i] = make([]rune, len(tab[i]))
				copy(tmp[i], tab[i])
			}
			liste = append(liste, tmp)
		}
	}
	str := ""
	for i := range tab {
		str += string(tab[i])
	}
	first := vus[str]
	finalTab := liste[(1000000000-first)%(iter-first)+first]
	sum := 0
	for i := range finalTab {
		sum += strings.Count(string(finalTab[i]), "O") * (len(finalTab) - i)
	}
	return sum
}

func south(tab *[][]rune, i, j int) int {
	if (*tab)[i][j] != 'O' {
		panic("Appel non-valide")
	}
	finalLine := i
	for k := i + 1; k < len(*tab); k++ {
		if (*tab)[k][j] == '.' {
			finalLine++
		} else {
			break
		}
	}
	if finalLine != i {
		(*tab)[i][j] = '.'
		(*tab)[finalLine][j] = 'O'
	}

	return finalLine
}

func east(tab *[][]rune, i, j int) int {
	if (*tab)[i][j] != 'O' {
		panic("Appel non-valide")
	}
	finalColumn := j
	for k := j + 1; k < len((*tab)[i]); k++ {
		if (*tab)[i][k] == '.' {
			finalColumn++
		} else {
			break
		}
	}
	if finalColumn != j {
		(*tab)[i][j] = '.'
		(*tab)[i][finalColumn] = 'O'
	}

	return finalColumn
}

func west(tab *[][]rune, i, j int) int {
	if (*tab)[i][j] != 'O' {
		panic("Appel non-valide")
	}
	finalColumn := j
	for k := j - 1; k >= 0; k-- {
		if (*tab)[i][k] == '.' {
			finalColumn--
		} else {
			break
		}
	}
	if finalColumn != j {
		(*tab)[i][j] = '.'
		(*tab)[i][finalColumn] = 'O'
	}

	return finalColumn
}
