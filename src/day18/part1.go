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

	grid := make(map[Case]bool)
	current := Case{
		X: 0,
		Y: 0,
	}
	grid[current] = true
	vertices := make([][2]int, 0)
	vertices = append(vertices, [2]int{current.X, current.Y})
	for _, instruction := range instructions {
		if instruction.Dir == 'U' {
			for j := 1; j <= instruction.Len; j++ {
				grid[Case{X: current.X - j, Y: current.Y}] = true
			}
			current.X -= instruction.Len
			if current.X == 0 && current.Y == 0 {
				break
			}
			vertices = append(vertices, [2]int{current.X, current.Y})
		} else if instruction.Dir == 'D' {
			for j := 1; j <= instruction.Len; j++ {
				grid[Case{X: current.X + j, Y: current.Y}] = true
			}
			current.X += instruction.Len
			if current.X == 0 && current.Y == 0 {
				break
			}
			vertices = append(vertices, [2]int{current.X, current.Y})
		} else if instruction.Dir == 'L' {
			for j := 1; j <= instruction.Len; j++ {
				grid[Case{X: current.X, Y: current.Y - j}] = true
			}
			current.Y -= instruction.Len
			vertices = append(vertices, [2]int{current.X, current.Y})
		} else if instruction.Dir == 'R' {
			for j := 1; j <= instruction.Len; j++ {
				grid[Case{X: current.X, Y: current.Y + j}] = true
			}
			current.Y += instruction.Len
			if current.X == 0 && current.Y == 0 {
				break
			}
			vertices = append(vertices, [2]int{current.X, current.Y})
		} else {
			panic("Invalid direction")
		}
	}
	return int(CalculerAirePolygone(vertices, grid)) + 1
}

func CalculerAirePolygone(vertices [][2]int, grid map[Case]bool) float64 {
	n := len(vertices)
	if n < 3 {
		// Un polygone doit avoir au moins trois sommets.
		return 0.0
	}

	// Appliquer l'algorithme de Shoelace pour calculer l'aire.
	aire := 0.0
	for i := 0; i < n-1; i++ {
		aire += float64(vertices[i][0]*vertices[i+1][1] - vertices[i+1][0]*vertices[i][1])
	}

	// Ajouter le dernier segment (du dernier au premier sommet).
	aire += float64(vertices[n-1][0]*vertices[0][1] - vertices[0][0]*vertices[n-1][1])

	// Prendre la valeur absolue et diviser par 2.
	if aire < 0 {
		aire = -aire
	}
	for range grid {
		aire++
	}
	aire = 0.5 * (aire)

	return aire
}
