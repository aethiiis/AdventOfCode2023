package main

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
)

func Part2(lines []string) int {
	// Découpage en blocs
	var blocs [][]string
	var emptyLines []int
	for i := range lines {
		if len(lines[i]) == 0 {
			emptyLines = append(emptyLines, i)
		}
	}

	/*fmt.Println("Les emptyLines sont : ")
	for i := range emptyLines {
		fmt.Println(emptyLines[i])
		fmt.Println("et")
	}*/

	blocs = append(blocs, lines[0:emptyLines[0]])
	for i := 0; i < len(emptyLines); i++ {
		if i == len(emptyLines)-1 {
			blocs = append(blocs, lines[emptyLines[i]+1:])
		} else {
			blocs = append(blocs, lines[emptyLines[i]+1:emptyLines[i+1]])
		}
	}

	/*fmt.Println("Les blocs sont : ")
	for i := range blocs {
		for j := range blocs[i] {
			fmt.Println(blocs[i][j])
		}
		fmt.Println("et")
	}*/

	// Symétrie
	/*fmt.Printf("Voici blocs : %v\n", blocs)*/
	sum := 0
	for _, bloc := range blocs {
		modified := false
		symmetry := false
		for j := range bloc {
			if j < len(bloc)-1 && bloc[j] == bloc[j+1] {
				check := true
				count := 1

				for j-count >= 0 && j+count+1 < len(bloc) && check {
					firstPart := bloc[j-count : j+1]
					secondPart := utilities.FlipLine(bloc[j+1 : j+2+count])
					/*fmt.Printf("Voici firstPart :  %v\n", firstPart)
					fmt.Printf("Voici secondPart : %v\n", secondPart)*/
					if len(firstPart) != len(secondPart) {
						panic("Problème de dimension")
					} else {
						for k := range firstPart {
							if firstPart[k] != secondPart[k] {
								check = false
								break
							}
						}
					}
					count++
				}
				/*fmt.Printf("Voici check : %t\n", check)
				fmt.Printf("Voici count : %d\n", count)*/
				if check {
					sum += (j + 1) * 100
					symmetry = true
				}
			}
		}
		if !symmetry {
			blocTrans := utilities.Transpose(bloc)
			for j := range blocTrans {
				if j < len(blocTrans)-1 && blocTrans[j] == blocTrans[j+1] {
					check := true
					count := 1

					for j-count >= 0 && j+count+1 < len(blocTrans) && check {
						firstPart := blocTrans[j-count : j+1]
						secondPart := utilities.FlipLine(blocTrans[j+1 : j+2+count])
						/*fmt.Printf("Voici firstPart :  %v\n", firstPart)
						fmt.Printf("Voici secondPart : %v\n", secondPart)*/
						if len(firstPart) != len(secondPart) {
							panic("Problème de dimension")
						} else {
							for k := range firstPart {
								if ok, index := utilities.StringsIdentiquesAUnCaracterePres(firstPart[k], secondPart[k]); !modified && ok {
									modified = false
									if blocTrans[j-count+k][index] == '.' {
										blocTrans[j-count+k] = blocTrans[j-count+k][0:index] + "#" + blocTrans[j-count+k][index+1:]
									} else if blocTrans[j-count+k][index] == '#' {
										blocTrans[j-count+k] = blocTrans[j-count+k][0:index] + "." + blocTrans[j-count+k][index+1:]
									} else {
										panic("Caractère non reconnu")
									}
									count--
								}
								check = false
								break
							}
						}
						count++
					}
					/*fmt.Printf("Voici check : %t\n", check)
					fmt.Printf("Voici count : %d\n", count)*/
					if check {
						sum += j + 1
					}
				}
			}
		}
		fmt.Printf("Voici sum : %d\n", sum)
	}

	return sum
}
