package utilities

func IsNeighbourToZero(lines [][]rune, i int, j int) bool {
	if lines[i-1][j] == '0' {
		return true
	} else if lines[i][j-1] == '0' {
		return true
	} else if lines[i][j+1] == '0' {
		return true
	} else if lines[i+1][j] == '0' {
		return true
	} else {
		return false
	}
}
