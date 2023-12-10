package utilities

func CheckAround(lines []string, xPos int, yPos int) [][]uint8 {
	tab := make([][]uint8, 3)
	for i := range tab {
		tab[i] = make([]uint8, 3)
	}

	if xPos == 0 && yPos == 0 { // Coin supérieur gauche

		tab[1][1] = 'C'                 // Le point en lui-même
		tab[1][2] = lines[xPos][yPos+1] // A droite

		tab[2][1] = lines[xPos+1][yPos] // En-dessous
		tab[2][2] = '.'                 // Diagonale inférieure droite

	} else if xPos == 0 && yPos == len(lines[0])-1 { // Coin supérieur droit

		tab[1][0] = lines[xPos][yPos-1] // A gauche
		tab[1][1] = 'C'                 // Le point en lui-même

		tab[2][0] = '.'                 // Diagonale inférieure gauche
		tab[2][1] = lines[xPos+1][yPos] // En-dessous

	} else if xPos == len(lines) && yPos == 0 { // Coin inférieur gauche

		tab[0][1] = lines[xPos-1][yPos] // Au-dessus
		tab[0][2] = '.'                 // Diagonale supérieure droite

		tab[1][1] = 'C'                 // Le point en lui-même
		tab[1][2] = lines[xPos][yPos+1] // A droite

	} else if xPos == len(lines) && yPos == len(lines[0])-1 { // Coin inférieur droit

		tab[0][0] = '.'                 // Diagonale supérieure gauche
		tab[0][1] = lines[xPos-1][yPos] // Au-dessus

		tab[1][0] = lines[xPos][yPos-1] // A gauche
		tab[1][1] = 'C'                 // Le point en lui-même

	} else if xPos == 0 { // Côté supérieur

		tab[1][0] = lines[xPos][yPos-1] // A gauche
		tab[1][1] = 'C'                 // Le point en lui-même
		tab[1][2] = lines[xPos][yPos+1] // A droite

		tab[2][0] = '.'                 // Diagonale inférieure gauche
		tab[2][1] = lines[xPos+1][yPos] // En-dessous
		tab[2][2] = '.'                 // Diagonale inférieure droite

	} else if xPos == len(lines) { // Côté inférieur

		tab[0][0] = '.'                 // Diagonale supérieure gauche
		tab[0][1] = lines[xPos-1][yPos] // Au-dessus
		tab[0][2] = '.'                 // Diagonale supérieure droite

		tab[1][0] = lines[xPos][yPos-1] // A gauche
		tab[1][1] = 'C'                 // Le point en lui-même
		tab[1][2] = lines[xPos][yPos+1] // A droite

	} else if yPos == 0 { // Côté gauche

		tab[0][1] = lines[xPos-1][yPos] // Au-dessus
		tab[0][2] = '.'                 // Diagonale supérieure droite

		tab[1][1] = 'C'                 // Le point en lui-même
		tab[1][2] = lines[xPos][yPos+1] // A droite

		tab[2][1] = lines[xPos+1][yPos] // En-dessous
		tab[2][2] = '.'                 // Diagonale inférieure droite

	} else if yPos == len(lines[0]) { // Côté droite

		tab[0][0] = '.'                 // Diagonale supérieure gauche
		tab[0][1] = lines[xPos-1][yPos] // Au-dessus

		tab[1][0] = lines[xPos][yPos-1] // A gauche
		tab[1][1] = 'C'                 // Le point en lui-même

		tab[2][0] = '.'                 // Diagonale inférieure gauche
		tab[2][1] = lines[xPos+1][yPos] // En-dessous

	} else {
		tab[0][0] = '.'                 // Diagonale supérieure gauche
		tab[0][1] = lines[xPos-1][yPos] // Au-dessus
		tab[0][2] = '.'                 // Diagonale supérieure droite

		tab[1][0] = lines[xPos][yPos-1] // A gauche
		tab[1][1] = 'C'                 // Le point en lui-même
		tab[1][2] = lines[xPos][yPos+1] // A droite

		tab[2][0] = '.'                 // Diagonale inférieure gauche
		tab[2][1] = lines[xPos+1][yPos] // En-dessous
		tab[2][2] = '.'                 // Diagonale inférieure droite
	}
	return tab
}
