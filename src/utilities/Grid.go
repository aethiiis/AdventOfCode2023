package utilities

type Pos struct {
	X, Y int
}

type Grid[T comparable] map[Pos]T

func (grid Grid[int]) NeighboursInt(pos Pos) ([]Pos, []int) {
	var listePos []Pos
	var listeValue []int
	if c, ok := grid[Pos{pos.X + 1, pos.Y}]; ok {
		listePos = append(listePos, Pos{pos.X, pos.Y - 1})
		listeValue = append(listeValue, c)
	}
	if c, ok := grid[Pos{pos.X, pos.Y - 1}]; ok {
		listePos = append(listePos, Pos{pos.X, pos.Y + 1})
		listeValue = append(listeValue, c)
	}
	if c, ok := grid[Pos{pos.X, pos.Y + 1}]; ok {
		listePos = append(listePos, Pos{pos.X - 1, pos.Y})
		listeValue = append(listeValue, c)
	}
	if c, ok := grid[Pos{pos.X, pos.Y - 1}]; ok {
		listePos = append(listePos, Pos{pos.X + 1, pos.Y})
		listeValue = append(listeValue, c)
	}
	return listePos, listeValue
}

func (grid Grid[rune]) NeighboursRune(pos Pos) ([]Pos, []rune) {
	var listePos []Pos
	var listeValue []rune
	if c, ok := grid[Pos{pos.X + 1, pos.Y}]; ok {
		listePos = append(listePos, Pos{pos.X, pos.Y - 1})
		listeValue = append(listeValue, c)
	}
	if c, ok := grid[Pos{pos.X, pos.Y - 1}]; ok {
		listePos = append(listePos, Pos{pos.X, pos.Y + 1})
		listeValue = append(listeValue, c)
	}
	if c, ok := grid[Pos{pos.X, pos.Y + 1}]; ok {
		listePos = append(listePos, Pos{pos.X - 1, pos.Y})
		listeValue = append(listeValue, c)
	}
	if c, ok := grid[Pos{pos.X, pos.Y - 1}]; ok {
		listePos = append(listePos, Pos{pos.X + 1, pos.Y})
		listeValue = append(listeValue, c)
	}
	return listePos, listeValue
}

func ConvertirTableauRuneGrid(tab [][]rune) Grid[rune] {
	grid := make(Grid[rune])
	for i := range tab {
		for j := range tab[i] {
			grid[Pos{
				X: i,
				Y: j,
			}] = tab[i][j]
		}
	}
	return grid
}
