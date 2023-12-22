package main

import (
	"strings"
)

func Part1(lines []string) int {
	instructions, mapLoc := getMap(lines)

	local := "AAA"
	count := 0
	index := 0
	for local != "ZZZ" {
		instruction := instructions[index]
		if index == len(instructions)-1 {
			index = 0
		} else {
			index++
		}

		dests := mapLoc[local]
		local = dests[instruction]

		count++
	}
	return count
}

func getMap(lines []string) ([]int, map[string][]string) {
	instructions := make([]int, len(lines[0]))
	mapLoc := make(map[string][]string)

	for i := range lines[0] {
		if rune(lines[0][i]) == 'L' {
			instructions[i] = 0
		} else if rune(lines[0][i]) == 'R' {
			instructions[i] = 1
		} else {
			panic("Caractère non-reconnu\n")
		}
	}

	for i := 2; i < len(lines); i++ {
		sourceDest := strings.SplitN(lines[i], "=", 2)
		source := sourceDest[0]
		dest := sourceDest[1]

		source = strings.Trim(source, " ")
		dest = strings.Trim(dest, " ()")

		dests := strings.SplitN(dest, ",", 2)
		dests[1] = dests[1][1:]

		mapLoc[source] = dests
	}

	return instructions, mapLoc
}
