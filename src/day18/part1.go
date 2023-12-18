package main

import (
	"strconv"
	"strings"
)

type Instruction struct {
	Dir rune
	Len int
	Rgb string
}

type Case struct {
	X int
	Y int
}

type Instructions []Instruction

func Part1(lines []string) int {
	var instructions Instructions
	instructions = make(Instructions, 0)
	for i := range lines {
		elements := strings.Split(lines[i], " ")
		dir := lines[i][0]
		Len, _ := strconv.Atoi(elements[1])
		instructions = append(instructions, Instruction{
			Dir: rune(dir),
			Len: Len,
			Rgb: elements[2],
		})
	}

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
	return CalculerAirePolygone(vertices, count) + 1
}

func CalculerAirePolygone(vertices [][2]int, count int) int {
	n := len(vertices)
	if n < 3 {
		// Un polygone doit avoir au moins trois sommets.
		return 0
	}

	// Appliquer l'algorithme de Shoelace pour calculer l'aire.
	aire := 0
	for i := 0; i < n-1; i++ {
		aire += vertices[i][0]*vertices[i+1][1] - vertices[i+1][0]*vertices[i][1]
	}

	// Ajouter le dernier segment (du dernier au premier sommet).
	aire += vertices[n-1][0]*vertices[0][1] - vertices[0][0]*vertices[n-1][1]

	// Prendre la valeur absolue et diviser par 2.
	if aire < 0 {
		aire = -aire
	}
	aire += count
	aire = aire / 2

	return aire
}
