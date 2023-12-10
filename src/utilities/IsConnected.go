package utilities

func IsConnected(tab [][]uint8) [][]uint8 {
	// Initialisation
	connection := make([][]uint8, 3)

	for i := range connection {
		connection[i] = make([]uint8, 3)
	}

	count := 0

	// Checks
	if tab[1][1] != 'C' {
		panic("Mauvais format de tableau")
	}

	// [0, 1]
	if tab[0][1] == '|' || tab[0][1] == '7' || tab[0][1] == 'F' {
		connection[0][1] = 1
		count++
	}

	// [1, 0]
	if tab[1][0] == '-' || tab[1][0] == 'L' || tab[1][0] == 'F' {
		connection[1][0] = 1
		count++
	}

	// [1, 2]
	if tab[1][2] == '-' || tab[1][2] == 'J' || tab[1][2] == '7' {
		connection[1][2] = 1
		count++
	}

	// [2, 1]
	if tab[2][1] == '|' || tab[2][1] == 'L' || tab[2][1] == 'J' {
		connection[2][1] = 1
		count++
	}

	// Checks
	if count != 2 {
		panic("Mauvais nombre de connections")
	}

	return connection
}
