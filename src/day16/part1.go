package main

import (
	"advent_of_code_2023/src/utilities"
	"fmt"
)

type Tile struct {
	Pos       utilities.Pos
	Mirror    rune
	Energized bool
}

type Beam struct {
	Case      Tile
	Direction rune
}

type Grid struct {
	data map[utilities.Pos]Tile
}

func MakeGrid() Grid {
	return Grid{
		data: make(map[utilities.Pos]Tile),
	}
}

func (grid Grid) Neighbours(pos utilities.Pos) []Tile {
	var listePos []utilities.Pos
	var listeValue []Tile
	if c, ok := grid.data[utilities.Pos{X: pos.X + 1, Y: pos.Y}]; ok {
		listePos = append(listePos, utilities.Pos{X: pos.X + 1, Y: pos.Y})
		listeValue = append(listeValue, c)
	}
	if c, ok := grid.data[utilities.Pos{X: pos.X - 1, Y: pos.Y}]; ok {
		listePos = append(listePos, utilities.Pos{X: pos.X - 1, Y: pos.Y})
		listeValue = append(listeValue, c)
	}
	if c, ok := grid.data[utilities.Pos{X: pos.X, Y: pos.Y + 1}]; ok {
		listePos = append(listePos, utilities.Pos{X: pos.X, Y: pos.Y + 1})
		listeValue = append(listeValue, c)
	}
	if c, ok := grid.data[utilities.Pos{X: pos.X, Y: pos.Y - 1}]; ok {
		listePos = append(listePos, utilities.Pos{X: pos.X, Y: pos.Y - 1})
		listeValue = append(listeValue, c)
	}
	return listeValue
}

func Part1(lines []string) int {
	var grid Grid
	grid = MakeGrid()
	for i := range lines {
		for j := range lines[i] {
			grid.data[utilities.Pos{
				X: i,
				Y: j,
			}] = Tile{Pos: utilities.Pos{X: i, Y: j}, Mirror: rune(lines[i][j]), Energized: false}
			/*fmt.Printf("%c", rune(lines[i][j]))*/
		}
	}
	/*fmt.Println()*/

	count := 0
	nrjList := make([][]rune, len(lines))
	for i := range nrjList {
		nrjList[i] = make([]rune, len(lines[i]))
		for j := range nrjList[i] {
			nrjList[i][j] = '.'
		}
	}

	var frontier utilities.Queue

	var start Beam
	start = Beam{Case: grid.data[utilities.Pos{X: 0, Y: 0}], Direction: 'd'}
	start.Case.Energized = true
	grid.data[utilities.Pos{X: 0, Y: 0}] = start.Case
	nrjList[0][0] = '#'

	frontier.Put(start)

	reached := make(map[Beam]bool)
	reached[start] = true
	count++

	for !frontier.IsEmpty() {
		//time.Sleep(1 * time.Second)
		var current Beam
		current = frontier.Get().(Beam)

		var nexts []Beam

		if current.Direction == 'h' {
			if current.Case.Mirror == '.' {
				if current.Case.Pos.X > 0 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'h'}
					nexts = append(nexts, next)
				}
			} else if current.Case.Mirror == '/' {
				if current.Case.Pos.Y < len(lines[0])-1 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'd'}
					nexts = append(nexts, next)
				}
			} else if current.Case.Mirror == '\u005C' {
				if current.Case.Pos.Y > 0 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'g'}
					nexts = append(nexts, next)
				}
			} else if current.Case.Mirror == '|' {
				if current.Case.Pos.X > 0 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'h'}
					nexts = append(nexts, next)
				}
			} else if current.Case.Mirror == '-' {
				if current.Case.Pos.Y > 0 {
					nextTileG := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}].Mirror, Energized: true}
					nextG := Beam{Case: nextTileG, Direction: 'g'}
					nexts = append(nexts, nextG)
				}
				if current.Case.Pos.Y < len(lines[0])-1 {
					nextTileD := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}].Mirror, Energized: true}
					nextD := Beam{Case: nextTileD, Direction: 'd'}
					nexts = append(nexts, nextD)
				}
			} else {
				fmt.Printf("Voici le caratère non-valide : %c\n", current.Case.Mirror)
				panic("Caractère non-valide")
			}
		} else if current.Direction == 'd' {
			if current.Case.Mirror == '.' {
				if current.Case.Pos.Y < len(lines[0])-1 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'd'}
					nexts = append(nexts, next)
				}
			} else if current.Case.Mirror == '/' {
				if current.Case.Pos.X > 0 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'h'}
					nexts = append(nexts, next)
				}
			} else if current.Case.Mirror == '\u005C' {
				if current.Case.Pos.X < len(lines)-1 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'b'}
					nexts = append(nexts, next)
				}
			} else if current.Case.Mirror == '|' {
				if current.Case.Pos.X > 0 {
					nextTileH := Tile{Pos: utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
					nextH := Beam{Case: nextTileH, Direction: 'h'}
					nexts = append(nexts, nextH)
				}
				if current.Case.Pos.X < len(lines)-1 {
					nextTileB := Tile{Pos: utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
					nextB := Beam{Case: nextTileB, Direction: 'b'}
					nexts = append(nexts, nextB)
				}
			} else if current.Case.Mirror == '-' {
				if current.Case.Pos.Y < len(lines[0])-1 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'd'}
					nexts = append(nexts, next)
				}
			} else {
				fmt.Printf("Voici le caratère non-valide : %c\n", current.Case.Mirror)
				panic("Caractère non-valide")
			}
		} else if current.Direction == 'b' {
			if current.Case.Pos.X < len(lines)-1 {
				if current.Case.Mirror == '.' {
					if current.Case.Pos.X < len(lines)-1 {
						nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
						next := Beam{Case: nextTile, Direction: 'b'}
						nexts = append(nexts, next)
					}
				} else if current.Case.Mirror == '/' {
					if current.Case.Pos.Y > 0 {
						nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}].Mirror, Energized: true}
						next := Beam{Case: nextTile, Direction: 'g'}
						nexts = append(nexts, next)
					}
				} else if current.Case.Mirror == '\u005C' {
					if current.Case.Pos.Y < len(lines[0])-1 {
						nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}].Mirror, Energized: true}
						next := Beam{Case: nextTile, Direction: 'd'}
						nexts = append(nexts, next)
					}
				} else if current.Case.Mirror == '|' {
					if current.Case.Pos.X < len(lines)-1 {
						nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
						next := Beam{Case: nextTile, Direction: 'b'}
						nexts = append(nexts, next)
					}
				} else if current.Case.Mirror == '-' {
					if current.Case.Pos.Y > 0 {
						nextTileG := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}].Mirror, Energized: true}
						nextG := Beam{Case: nextTileG, Direction: 'g'}
						nexts = append(nexts, nextG)
					}
					if current.Case.Pos.Y < len(lines[0])-1 {
						nextTileD := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y + 1}].Mirror, Energized: true}
						nextD := Beam{Case: nextTileD, Direction: 'd'}
						nexts = append(nexts, nextD)
					}
				} else {
					fmt.Printf("Voici le caratère non-valide : %c\n", current.Case.Mirror)
					panic("Caractère non-valide")
				}
			}
		} else if current.Direction == 'g' {
			if current.Case.Mirror == '.' {
				if current.Case.Pos.Y > 0 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'g'}
					nexts = append(nexts, next)
				}
			} else if current.Case.Mirror == '/' {
				if current.Case.Pos.X < len(lines[0])-1 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'b'}
					nexts = append(nexts, next)
				}
			} else if current.Case.Mirror == '\u005C' {
				if current.Case.Pos.X > 0 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'h'}
					nexts = append(nexts, next)
				}
			} else if current.Case.Mirror == '|' {
				if current.Case.Pos.X > 0 {
					nextTileH := Tile{Pos: utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X - 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
					nextH := Beam{Case: nextTileH, Direction: 'h'}
					nexts = append(nexts, nextH)
				}
				if current.Case.Pos.X < len(lines)-1 {
					nextTileB := Tile{Pos: utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X + 1, Y: current.Case.Pos.Y}].Mirror, Energized: true}
					nextB := Beam{Case: nextTileB, Direction: 'b'}
					nexts = append(nexts, nextB)
				}
			} else if current.Case.Mirror == '-' {
				if current.Case.Pos.Y > 0 {
					nextTile := Tile{Pos: utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}, Mirror: grid.data[utilities.Pos{X: current.Case.Pos.X, Y: current.Case.Pos.Y - 1}].Mirror, Energized: true}
					next := Beam{Case: nextTile, Direction: 'g'}
					nexts = append(nexts, next)
				}
			} else {
				fmt.Printf("Voici le caratère non-valide : %c\n", current.Case.Mirror)
				panic("Caractère non-valide")
			}
		} else {
			panic("Valeur de Direction non-valide")
		}

		for _, beam := range nexts {
			if _, ok := reached[beam]; !ok {
				frontier.Put(beam)
				reached[beam] = true
				nrjList[beam.Case.Pos.X][beam.Case.Pos.Y] = '#'
				if grid.data[beam.Case.Pos].Energized == false {
					count++
					nrj := grid.data[beam.Case.Pos]
					nrj.Energized = true
					grid.data[beam.Case.Pos] = nrj
				}
			}
		}
		/*fmt.Println("Voici nrjList : ")
		for i := range nrjList {
			fmt.Println(string(nrjList[i]))
		}*/
		//time.Sleep(50 * time.Millisecond)
	}
	return count
}
