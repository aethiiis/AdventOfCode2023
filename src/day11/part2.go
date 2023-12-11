package main

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
)

func Part2(lines []string) int {
	var galax [][]int
	/*var emptyLine []int
	var emptyColumn []int*/
	newLines := utilities.DoubleEmptyLineColumn(lines, "*")
	fmt.Println("Voici newLines : ")
	for i := range newLines {
		fmt.Println(newLines[i])
	}

	tab := utilities.ConvertirListeStringsTableauRunes(newLines)
	for i := range tab {
		for j := range tab[i] {
			if tab[i][j] == '#' {
				galax = append(galax, []int{i, j})
			}
		}
	}
	//fmt.Printf("Voici galax : %v\n", galax)*

	sum := 0

	/*grid := utilities.ConvertirTableauRuneGrid(tab)
	start := utilities.Pos{X: galax[0][0], Y: galax[0][1]}
	goal := utilities.Pos{X: galax[1][0], Y: galax[1][1]}
	frontier := make(utilities.PriorityQueue, 0)
	heap.Push(&frontier, &utilities.Item{Value: start, Priority: 1})
	cameFrom := make(map[utilities.Pos]utilities.Pos)
	costSoFar := make(map[utilities.Pos]int)
	cameFrom[start] = utilities.Pos{}
	costSoFar[start] = 0

	for frontier.Len() != 0 {
		current := heap.Pop(&frontier).(*utilities.Item)

		listNeighbors, _ := grid.NeighboursRune(current.Value)
		for _, next := range listNeighbors {
			var nextCost int
			if grid[next] == '*' {
				nextCost = 10
			} else {
				nextCost = 1
			}

			newCost := costSoFar[current.Value] + nextCost
			if _, ok := costSoFar[next]; !ok || newCost < costSoFar[next] {
				costSoFar[next] = newCost
				heap.Push(&frontier, &utilities.Item{
					Value:    next,
					Priority: newCost,
				})
				cameFrom[next] = current.Value
			}
			if current.Value == goal {
				break
			}
		}
	}
	fmt.Printf("Voici le coup pour aller de 1 à 2 : %d\n", costSoFar[goal])*/
	return sum
}
