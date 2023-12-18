package main

import (
	"strconv"
	"strings"
)

func Part2(lines []string) int {
	var instructions Instructions
	instructions = make(Instructions, 0)
	for i := range lines {
		elements := strings.Split(lines[i], " ")
		dirN := elements[2][len(elements[2])-2]
		var dir rune
		if dirN == 48 {
			dir = 'R'
		} else if dirN == 49 {
			dir = 'D'
		} else if dirN == 50 {
			dir = 'L'
		} else if dirN == 51 {
			dir = 'U'
		}
		/*Len, _ := hexToDecimal(elements[2][:len(elements[2])-1])*/
		Len, _ := strconv.ParseInt(elements[2][2:len(elements[2])-2], 16, 64)
		/*fmt.Println(dir, Len, elements[2][2:len(elements[2])-2])*/
		instructions = append(instructions, Instruction{
			Dir: dir,
			Len: int(Len),
			Rgb: elements[2],
		})
	}
	//fmt.Println(instructions)
	current := Case{
		X: 0,
		Y: 0,
	}
	vertices := make([][2]int, 0)
	vertices = append(vertices, [2]int{current.X, current.Y})
	count := 1
	for _, instruction := range instructions {
		if instruction.Dir == 'U' {
			count += instruction.Len
			current.X -= instruction.Len
			if current.X == 0 && current.Y == 0 {
				break
			}
			vertices = append(vertices, [2]int{current.X, current.Y})
		} else if instruction.Dir == 'D' {
			count += instruction.Len
			current.X += instruction.Len
			if current.X == 0 && current.Y == 0 {
				break
			}
			vertices = append(vertices, [2]int{current.X, current.Y})
		} else if instruction.Dir == 'L' {
			count += instruction.Len
			current.Y -= instruction.Len
			vertices = append(vertices, [2]int{current.X, current.Y})
		} else if instruction.Dir == 'R' {
			count += instruction.Len
			current.Y += instruction.Len
			if current.X == 0 && current.Y == 0 {
				break
			}
			vertices = append(vertices, [2]int{current.X, current.Y})
		} else {
			panic("Invalid direction")
		}
	}
	return int(CalculerAirePolygone(vertices, count)) + 1
}
