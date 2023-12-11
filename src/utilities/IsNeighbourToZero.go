package utilities

func IsNeighbourToZero(lines [][]rune, i int, j int) bool {
	return lines[i-1][j] == '0' || lines[i][j-1] == '0' || lines[i][j+1] == '0' || lines[i+1][j] == '0'
}
