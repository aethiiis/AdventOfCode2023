package main

type Ray struct {
	x  int
	y  int
	dx int
	dy int
}

type Place struct {
	x int
	y int
}

func Part1(lines []string) int {
	var start Ray
	start = Ray{
		x:  0,
		y:  -1,
		dx: 0,
		dy: 1,
	}
	return calc(start, lines)
}

func calc(start Ray, tab []string) int {
	// Set pour éviter de retomber sur les mêmes éléments depuis le même angle
	var seen map[Ray]bool
	seen = make(map[Ray]bool)
	// File pour stocker les éléments à aborder ainsi que leur angle
	var queue []Ray
	queue = make([]Ray, 0)
	queue = append(queue, start)
	// Tant qu'il reste des éléments dans la file
	for len(queue) != 0 {
		current := queue[0]
		x := current.x
		y := current.y
		dx := current.dx
		dy := current.dy
		queue = queue[1:]
		// On met à jour les coordonnées de la position
		x += dx
		y += dy
		// Si cela nous amène en dehors de la grille, on passe
		if x < 0 || x >= len(tab) || y < 0 || y >= len(tab[0]) {
			continue
		}
		// On récupère le caractère de la grille
		char := tab[x][y]
		if char == '.' || (char == '-' && dy != 0) || (char == '|' && dx != 0) {
			// Cas où l'on ne change pas de direction
			next := Ray{
				x:  x,
				y:  y,
				dx: dx,
				dy: dy,
			}
			if _, ok := seen[next]; !ok {
				seen[next] = true
				queue = append(queue, next)
			}
		} else if char == '/' {
			// Cas '/' :
			// dx = 1 (bas) --> dy = -1 (gauche) || dx = -1 (haut) --> dy = 1 (droite)
			// dy = 1 (droite) --> dx = -1 (haut) || dy = -1 (gauche) --> dx = 1 (bas)
			dx, dy = -dy, -dx
			next := Ray{
				x:  x,
				y:  y,
				dx: dx,
				dy: dy,
			}
			if _, ok := seen[next]; !ok {
				seen[next] = true
				queue = append(queue, next)
			}
		} else if char == '\\' {
			// Cas '\' :
			// dx = 1 (bas) --> dy = 1 (droite) || dx = -1 (haut) --> dy = -1 (gauche)
			// dy = 1 (droite) --> dx = 1 (bas) || dy = -1 (gauche) --> dx = -1 (haut)
			dx, dy = dy, dx
			next := Ray{
				x:  x,
				y:  y,
				dx: dx,
				dy: dy,
			}
			if _, ok := seen[next]; !ok {
				seen[next] = true
				queue = append(queue, next)
			}
		} else if char == '|' {
			// Cas '|' :
			// dy = 1 || dy = -1 --> dx = 1 && dx = -1
			nextB := Ray{
				x:  x,
				y:  y,
				dx: 1,
				dy: 0,
			}
			if _, ok := seen[nextB]; !ok {
				seen[nextB] = true
				queue = append(queue, nextB)
			}
			nextH := Ray{
				x:  x,
				y:  y,
				dx: -1,
				dy: 0,
			}
			if _, ok := seen[nextH]; !ok {
				seen[nextH] = true
				queue = append(queue, nextH)
			}
		} else if char == '-' {
			// Cas '-' :
			// dx = 1 || dx = -1 --> dy = 1 && dy = -1
			nextD := Ray{
				x:  x,
				y:  y,
				dx: 0,
				dy: 1,
			}
			if _, ok := seen[nextD]; !ok {
				seen[nextD] = true
				queue = append(queue, nextD)
			}
			nextG := Ray{
				x:  x,
				y:  y,
				dx: 0,
				dy: -1,
			}
			if _, ok := seen[nextG]; !ok {
				seen[nextG] = true
				queue = append(queue, nextG)
			}
		}
	}
	// Set pour compter le nombre de points distincts énergisés
	var coordinates map[Place]bool
	coordinates = make(map[Place]bool)
	// On ajoute chaque point énergisé de seen dans coordinates
	for ray := range seen {
		x := ray.x
		y := ray.y
		coords := Place{
			x: x,
			y: y,
		}
		if _, ok := coordinates[coords]; !ok {
			coordinates[coords] = true
		}
	}
	// On compte le nombre d'éléments dans coordinates
	var count int
	count = 0
	for range coordinates {
		count++
	}
	return count
}
