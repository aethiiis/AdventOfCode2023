package main

import (
	"advent_of_code_2023/src/utilities"
)

type quad struct {
	X  int
	Y  int
	dX int
	dY int
}

func Part1_2(lines []string) int {
	var tab [][]rune
	tab = make([][]rune, len(lines))
	for i := range tab {
		tab[i] = make([]rune, len(lines[i]))
		for j := range tab[i] {
			tab[i][j] = rune(lines[i][j])
		}
	}

	count := calc(&tab, quad{X: 0, Y: 0, dX: 0, dY: 1})
	return count
}

func calc(tab *[][]rune, coords quad) int {
	start := coords
	seen := make(map[quad]bool)
	seen[start] = true
	frontier := utilities.Queue{}
	frontier.Put(start)
	nrj := make([][]rune, len(*tab))
	for i := range nrj {
		nrj[i] = make([]rune, len((*tab)[i]))
		for j := range nrj[i] {
			nrj[i][j] = '.'
		}
	}
	nrj[start.X][start.Y] = '#'

	for !frontier.IsEmpty() {
		/*fmt.Println("Voici nrjList : ")
		for i := range nrj {
			fmt.Println(string(nrj[i]))
		}*/
		current := frontier.Get().(quad)
		X := current.X
		Y := current.Y
		dX := current.dX
		dY := current.dY

		X += dX
		Y += dY

		if X < 0 || X >= len(*tab) || Y < 0 || Y >= len((*tab)[0]) {
			continue
		} else {
			char := (*tab)[X][Y]

			if char == '.' || (char == '-' && dY != 0) || (char == '|' && dX != 0) {
				if _, ok := seen[quad{X: X, Y: Y, dX: dX, dY: dY}]; !ok {
					seen[quad{X: X, Y: Y, dX: dX, dY: dY}] = true
					frontier.Put(quad{X: X, Y: Y, dX: dX, dY: dY})
					nrj[X][Y] = '#'
				}
			} else if char == '/' {
				dX, dY = -dY, -dX
				if _, ok := seen[quad{X: X, Y: Y, dX: dX, dY: dY}]; !ok {
					seen[quad{X: X, Y: Y, dX: dX, dY: dY}] = true
					frontier.Put(quad{X: X, Y: Y, dX: dX, dY: dY})
					nrj[X][Y] = '#'
				}
			} else if char == '\u005C' {
				dX, dY = dY, dX
				if _, ok := seen[quad{X: X, Y: Y, dX: dX, dY: dY}]; !ok {
					seen[quad{X: X, Y: Y, dX: dX, dY: dY}] = true
					frontier.Put(quad{X: X, Y: Y, dX: dX, dY: dY})
					nrj[X][Y] = '#'
				}
			} else {
				if char == '|' {
					if _, ok := seen[quad{X: X, Y: Y, dX: 1, dY: 0}]; !ok {
						seen[quad{X: X, Y: Y, dX: 1, dY: 0}] = true
						frontier.Put(quad{X: X, Y: Y, dX: 1, dY: 0})
						nrj[X][Y] = '#'
					}
					if _, ok := seen[quad{X: X, Y: Y, dX: -1, dY: 0}]; !ok {
						seen[quad{X: X, Y: Y, dX: -1, dY: 0}] = true
						frontier.Put(quad{X: X, Y: Y, dX: -1, dY: 0})
						nrj[X][Y] = '#'
					}
				} else if char == '-' {
					if _, ok := seen[quad{X: X, Y: Y, dX: 0, dY: 1}]; !ok {
						seen[quad{X: X, Y: Y, dX: 0, dY: 1}] = true
						frontier.Put(quad{X: X, Y: Y, dX: 0, dY: 1})
						nrj[X][Y] = '#'
					}
					if _, ok := seen[quad{X: X, Y: Y, dX: 0, dY: -1}]; !ok {
						seen[quad{X: X, Y: Y, dX: 0, dY: -1}] = true
						frontier.Put(quad{X: X, Y: Y, dX: 0, dY: -1})
						nrj[X][Y] = '#'
					}
				} else {
					panic("Caractère non-reconnu")
				}
			}
		}
	}
	count := 0
	for i := range nrj {
		for j := range nrj[i] {
			if nrj[i][j] == '#' {
				count++
			}
		}
	}
	return count
}
