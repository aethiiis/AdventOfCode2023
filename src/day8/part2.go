package main

import (
	"fmt"
	"strings"
)

func Part2(lines []string) int {
	instructions, mapLoc := getMap(lines)

	var startingNodes []string
	for cle := range mapLoc {
		if strings.HasSuffix(cle, "A") {
			startingNodes = append(startingNodes, cle)
		}
	}

	fmt.Printf("instructions est : %v\n", instructions)
	fmt.Printf("mapLoc est : %v\n", mapLoc)

	count := 0
	index := 0

	for !isAllZ(startingNodes) {
		fmt.Printf("Starting node est : %v\n", startingNodes)
		instruction := instructions[index]
		if index == len(instructions)-1 {
			index = 0
		} else {
			index++
		}

		for j := range startingNodes {
			startingNodes[j] = mapLoc[startingNodes[j]][instruction]
		}

		count++
	}
	return count
}

func isAllZ(startingNodes []string) bool {
	for i := range startingNodes {
		if !strings.HasSuffix(startingNodes[i], "Z") {
			return false
		}
	}
	return true
}
