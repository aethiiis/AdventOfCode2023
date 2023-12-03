package utilities

func SymbolSensor(position []int, i int, lines []string, carteGear [][]int) bool {
	ok := false
	// Gestion des cas de bord

	// Cas tout en haut
	if i == 0 {
		// Cas toute la première ligne
		if position[0] == 0 && position[1] == len(lines[i]) {
			// Ligne du bas juste en bas
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i+1][index])) {
					if rune(lines[i+1][index]) == '*' {
						carteGear[i+1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
		} else if position[0] == 0 { // Cas tout en haut à gauche
			// Caractère à droite
			if IsNotPointDigit(rune(lines[i][position[1]])) {
				if rune(lines[i][position[1]]) == '*' {
					carteGear[i][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Ligne du bas juste en bas
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i+1][index])) {
					if rune(lines[i+1][index]) == '*' {
						carteGear[i+1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
			// Ligne du bas diagonale droite
			if IsNotPointDigit(rune(lines[i+1][position[1]])) {
				if rune(lines[i+1][position[1]]) == '*' {
					carteGear[i+1][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
		} else if position[1] == len(lines[i]) { // Cas tout en haut à droite
			// Caractère à gauche
			if IsNotPointDigit(rune(lines[i][position[0]-1])) {
				if rune(lines[i][position[0]-1]) == '*' {
					carteGear[i][position[0]-1] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Ligne du bas juste en bas
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i+1][index])) {
					if rune(lines[i+1][index]) == '*' {
						carteGear[i+1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
			// Ligne du bas diagonale gauche
			if IsNotPointDigit(rune(lines[i+1][position[0]-1])) {
				if rune(lines[i+1][position[0]-1]) == '*' {
					carteGear[i+1][position[0]-1] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
		} else { // Cas basique tout en haut
			// Caractère à gauche
			if IsNotPointDigit(rune(lines[i][position[0]-1])) {
				if rune(lines[i][position[0]-1]) == '*' {
					carteGear[i][position[0]-1] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Caractère à droite
			if IsNotPointDigit(rune(lines[i][position[1]])) {
				if rune(lines[i][position[1]]) == '*' {
					carteGear[i][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Ligne du bas juste en bas
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i+1][index])) {
					if rune(lines[i+1][index]) == '*' {
						carteGear[i+1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
			// Ligne du bas diagonale gauche
			if IsNotPointDigit(rune(lines[i+1][position[0]-1])) {
				if rune(lines[i+1][position[0]-1]) == '*' {
					carteGear[i+1][position[0]-1] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Ligne du bas diagonale droite
			if IsNotPointDigit(rune(lines[i+1][position[1]])) {
				if rune(lines[i+1][position[1]]) == '*' {
					carteGear[i+1][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
		}
	} else if i == len(lines)-1 { // Cas tout en bas
		// Cas toute la dernière ligne
		if position[0] == 0 && position[1] == len(lines[i]) {
			// Ligne du dessus juste au-dessus
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i-1][index])) {
					if rune(lines[i-1][index]) == '*' {
						carteGear[i-1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
		} else if position[0] == 0 { // Cas tout en bas à gauche
			// Ligne du dessus juste au-dessus
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i-1][index])) {
					if rune(lines[i-1][index]) == '*' {
						carteGear[i-1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
			// Ligne du dessus diagonale droite
			if IsNotPointDigit(rune(lines[i-1][position[1]])) {
				if rune(lines[i-1][position[1]]) == '*' {
					carteGear[i-1][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Caractère à droite
			if IsNotPointDigit(rune(lines[i][position[1]])) {
				if rune(lines[i][position[1]]) == '*' {
					carteGear[i][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
		} else if position[1] == len(lines[i]) { // Cas tout en bas à droite
			// Ligne du dessus juste au-dessus
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i-1][index])) {
					if rune(lines[i-1][index]) == '*' {
						carteGear[i-1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
			// Ligne du dessus diagonale gauche
			if IsNotPointDigit(rune(lines[i-1][position[0]-1])) {
				if rune(lines[i-1][position[0]-1]) == '*' {
					carteGear[i-1][position[0]-1] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Caractère à gauche
			if IsNotPointDigit(rune(lines[i][position[0]-1])) {
				if rune(lines[i][position[0]-1]) == '*' {
					carteGear[i][position[0]-1] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
		} else { // Cas basique tout en bas
			// Ligne du dessus juste au-dessus
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i-1][index])) {
					if rune(lines[i-1][index]) == '*' {
						carteGear[i-1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
			// Ligne du dessus diagonale gauche
			if IsNotPointDigit(rune(lines[i-1][position[0]-1])) {
				if rune(lines[i-1][position[0]-1]) == '*' {
					carteGear[i-1][position[0]-1] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Ligne du dessus diagonale droite
			if IsNotPointDigit(rune(lines[i-1][position[1]])) {
				if rune(lines[i-1][position[1]]) == '*' {
					carteGear[i-1][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Caractère à gauche
			if IsNotPointDigit(rune(lines[i][position[0]-1])) {
				if rune(lines[i][position[0]-1]) == '*' {
					carteGear[i][position[0]-1] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Caractère à droite
			if IsNotPointDigit(rune(lines[i][position[1]])) {
				if rune(lines[i][position[1]]) == '*' {
					carteGear[i][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
		}
	} else if position[0] == 0 { // Cas tout à gauche
		// Cas toute la ligne
		if position[1] == len(lines[i]) {
			// Ligne du dessus juste au-dessus
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i-1][index])) {
					if rune(lines[i-1][index]) == '*' {
						carteGear[i-1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
			// Ligne du bas juste en bas
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i+1][index])) {
					if rune(lines[i+1][index]) == '*' {
						carteGear[i+1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
		} else { // Cas basique tout à gauche
			// Ligne du dessus juste au-dessus
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i-1][index])) {
					if rune(lines[i-1][index]) == '*' {
						carteGear[i-1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
			// Ligne du dessus diagonale droite
			if IsNotPointDigit(rune(lines[i-1][position[1]])) {
				if rune(lines[i-1][position[1]]) == '*' {
					carteGear[i-1][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Caractère à droite
			if IsNotPointDigit(rune(lines[i][position[1]])) {
				if rune(lines[i][position[1]]) == '*' {
					carteGear[i][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Ligne du bas diagonale droite
			if IsNotPointDigit(rune(lines[i+1][position[1]])) {
				if rune(lines[i+1][position[1]]) == '*' {
					carteGear[i+1][position[1]] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
			// Ligne du bas juste en bas
			for index := position[0]; index < position[1]; index++ {
				if IsNotPointDigit(rune(lines[i+1][index])) {
					if rune(lines[i+1][index]) == '*' {
						carteGear[i+1][index] += 1
						for j := position[0]; j < position[1]; j++ {
							carteGear[i][j] = -1
						}
					}
					ok = true
					return ok
				}
			}
		}
	} else if position[1] == len(lines[i]) { // Cas tout à droite
		// Cas basique tout à droite
		// Ligne du dessus juste au-dessus
		for index := position[0]; index < position[1]; index++ {
			if IsNotPointDigit(rune(lines[i-1][index])) {
				if rune(lines[i-1][index]) == '*' {
					carteGear[i-1][index] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
		}
		// Ligne du dessus diagonale gauche
		if IsNotPointDigit(rune(lines[i-1][position[0]-1])) {
			if rune(lines[i-1][position[0]-1]) == '*' {
				carteGear[i-1][position[0]-1] += 1
				for j := position[0]; j < position[1]; j++ {
					carteGear[i][j] = -1
				}
			}
			ok = true
			return ok
		}
		// Caractère à gauche
		if IsNotPointDigit(rune(lines[i][position[0]-1])) {
			if rune(lines[i][position[0]-1]) == '*' {
				carteGear[i][position[0]-1] += 1
				for j := position[0]; j < position[1]; j++ {
					carteGear[i][j] = -1
				}
			}
			ok = true
			return ok
		}
		// Ligne du bas diagonale gauche
		if IsNotPointDigit(rune(lines[i+1][position[0]-1])) {
			if rune(lines[i+1][position[0]-1]) == '*' {
				carteGear[i+1][position[0]-1] += 1
				for j := position[0]; j < position[1]; j++ {
					carteGear[i][j] = -1
				}
			}
			ok = true
			return ok
		}
		// Ligne du bas juste en bas
		for index := position[0]; index < position[1]; index++ {
			if IsNotPointDigit(rune(lines[i+1][index])) {
				if rune(lines[i+1][index]) == '*' {
					carteGear[i+1][index] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
		}
	} else { // Cas basique
		// Ligne du dessus diagonale gauche
		if IsNotPointDigit(rune(lines[i-1][position[0]-1])) {
			if rune(lines[i-1][position[0]-1]) == '*' {
				carteGear[i-1][position[0]-1] += 1
				for j := position[0]; j < position[1]; j++ {
					carteGear[i][j] = -1
				}
			}
			ok = true
			return ok
		}
		// Ligne du dessus diagonale droite
		if IsNotPointDigit(rune(lines[i-1][position[1]])) {
			if rune(lines[i-1][position[1]]) == '*' {
				carteGear[i-1][position[1]] += 1
				for j := position[0]; j < position[1]; j++ {
					carteGear[i][j] = -1
				}
			}
			ok = true
			return ok
		}
		// Ligne du dessus juste au-dessus
		for index := position[0]; index < position[1]; index++ {
			if IsNotPointDigit(rune(lines[i-1][index])) {
				if rune(lines[i-1][index]) == '*' {
					carteGear[i-1][index] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
		}
		// Caractère à gauche
		if IsNotPointDigit(rune(lines[i][position[0]-1])) {
			if rune(lines[i][position[0]-1]) == '*' {
				carteGear[i][position[0]-1] += 1
				for j := position[0]; j < position[1]; j++ {
					carteGear[i][j] = -1
				}
			}
			ok = true
			return ok
		}
		// Caractère à droite
		if IsNotPointDigit(rune(lines[i][position[1]])) {
			if rune(lines[i][position[1]]) == '*' {
				carteGear[i][position[1]] += 1
				for j := position[0]; j < position[1]; j++ {
					carteGear[i][j] = -1
				}
			}
			ok = true
			return ok
		}
		// Ligne du bas diagonale gauche
		if IsNotPointDigit(rune(lines[i+1][position[0]-1])) {
			if rune(lines[i+1][position[0]-1]) == '*' {
				carteGear[i+1][position[0]-1] += 1
				for j := position[0]; j < position[1]; j++ {
					carteGear[i][j] = -1
				}
			}
			ok = true
			return ok
		}
		// Ligne du bas diagonale droite
		if IsNotPointDigit(rune(lines[i+1][position[1]])) {
			if rune(lines[i+1][position[1]]) == '*' {
				carteGear[i+1][position[1]] += 1
				for j := position[0]; j < position[1]; j++ {
					carteGear[i][j] = -1
				}
			}
			ok = true
			return ok
		}
		// Ligne du bas juste en bas
		for index := position[0]; index < position[1]; index++ {
			if IsNotPointDigit(rune(lines[i+1][index])) {
				if rune(lines[i+1][index]) == '*' {
					carteGear[i+1][index] += 1
					for j := position[0]; j < position[1]; j++ {
						carteGear[i][j] = -1
					}
				}
				ok = true
				return ok
			}
		}
	}
	return ok
}
