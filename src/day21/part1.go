package main

type Plots struct {
	x    int
	y    int
	step int
}

func isInMap(seen map[Plots]bool, x int, y int) bool {
	for key := range seen {
		if key.x == x && key.y == y {
			return true
		}
	}
	return false
}

func Part1(lines []string) int {
	var frontier []Plots
	var seen map[Plots]bool
	seen = make(map[Plots]bool)
	var start Plots
	// Localiser S
	for i, line := range lines {
		for j, char := range line {
			if char == 'S' {
				start = Plots{
					x:    i,
					y:    j,
					step: 0,
				}
			}
		}
	}
	frontier = append(frontier, start)
	// Différenciation input/exemple
	var maxSteps int
	if len(lines) == 11 {
		maxSteps = 6
	} else {
		maxSteps = 64
	}
	lastStep := 0
	// Début du parcours
	count := 0
	for len(frontier) != 0 {
		current := frontier[0]
		x := current.x
		y := current.y
		step := current.step
		frontier = frontier[1:]

		if step+1 > maxSteps && !isInMap(seen, x+1, y) {
			count++
			continue
		}
		if step > lastStep {
			seen = make(map[Plots]bool)
		}
		if x+1 < len(lines) && (lines[x+1][y] == '.' || lines[x+1][y] == 'S') && !isInMap(seen, x+1, y) {
			next := Plots{
				x:    x + 1,
				y:    y,
				step: step + 1,
			}
			frontier = append(frontier, next)
			seen[next] = true
		}
		if x-1 >= 0 && (lines[x-1][y] == '.' || lines[x-1][y] == 'S') && !isInMap(seen, x-1, y) {
			next := Plots{
				x:    x - 1,
				y:    y,
				step: step + 1,
			}
			frontier = append(frontier, next)
			seen[next] = true
		}
		if y+1 < len(lines[0]) && (lines[x][y+1] == '.' || lines[x][y+1] == 'S') && !isInMap(seen, x, y+1) {
			next := Plots{
				x:    x,
				y:    y + 1,
				step: step + 1,
			}
			frontier = append(frontier, next)
			seen[next] = true
		}
		if y-1 >= 0 && (lines[x][y-1] == '.' || lines[x][y-1] == 'S') && !isInMap(seen, x, y-1) {
			next := Plots{
				x:    x,
				y:    y - 1,
				step: step + 1,
			}
			frontier = append(frontier, next)
			seen[next] = true
		}
		lastStep = step
	}
	return count
}
