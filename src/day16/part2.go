package main

func Part2(lines []string) int {
	var maxE int
	maxE = 0
	// On teste pour toutes les cases des côtés gauche et droit, ligne par ligne
	for x := range lines {
		var gauche Ray
		gauche = Ray{
			x:  x,
			y:  -1,
			dx: 0,
			dy: 1,
		}
		var droite Ray
		droite = Ray{
			x:  x,
			y:  len(lines[0]),
			dx: 0,
			dy: -1,
		}
		maxE = maxInt(maxE, calc(gauche, lines))
		maxE = maxInt(maxE, calc(droite, lines))
	}
	// On teste pour toutes les cases des côtés haut et bas, colonne par colonne
	for y := range lines[0] {
		var haut Ray
		haut = Ray{
			x:  -1,
			y:  y,
			dx: 1,
			dy: 0,
		}
		var bas Ray
		bas = Ray{
			x:  len(lines),
			y:  y,
			dx: -1,
			dy: 0,
		}
		maxE = maxInt(maxE, calc(haut, lines))
		maxE = maxInt(maxE, calc(bas, lines))
	}
	return maxE
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
