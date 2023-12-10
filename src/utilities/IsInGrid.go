package utilities

func IsInGrid(runesTab2 [][]rune) bool {
	for i := range runesTab2 {
		for j := range runesTab2[i] {
			if (runesTab2[i][j] == '.' && IsNeighbourToZero(runesTab2, i, j)) || runesTab2[i][j] == ',' {
				return false
			}
		}
	}
	return true
}
