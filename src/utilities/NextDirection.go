package utilities

func NextDirection(pipe []int, runeTab [][]rune) []int {
	if len(pipe) != 2 {
		panic("Problème de dimension")
	}

	newPipe := make([]int, 2)

	if runeTab[pipe[0]][pipe[1]] == '|' {

		if runeTab[pipe[0]-1][pipe[1]] == '*' {
			newPipe[0] = pipe[0] + 1
		} else if runeTab[pipe[0]+1][pipe[1]] == '*' {
			newPipe[0] = pipe[0] - 1
		}
		newPipe[1] = pipe[1]

	} else if runeTab[pipe[0]][pipe[1]] == '-' {

		if runeTab[pipe[0]][pipe[1]-1] == '*' {
			newPipe[1] = pipe[1] + 1
		} else if runeTab[pipe[0]][pipe[1]+1] == '*' {
			newPipe[1] = pipe[1] - 1
		} else {
			panic("On ne trouve pas où aller ensuite")
		}
		newPipe[0] = pipe[0]

	} else if runeTab[pipe[0]][pipe[1]] == 'L' {

		if runeTab[pipe[0]-1][pipe[1]] == '*' {
			newPipe[0] = pipe[0]
			newPipe[1] = pipe[1] + 1
		} else if runeTab[pipe[0]][pipe[1]+1] == '*' {
			newPipe[0] = pipe[0] - 1
			newPipe[1] = pipe[1]
		} else {
			panic("On ne trouve pas où aller ensuite")
		}

	} else if runeTab[pipe[0]][pipe[1]] == 'J' {

		if runeTab[pipe[0]-1][pipe[1]] == '*' {
			newPipe[0] = pipe[0]
			newPipe[1] = pipe[1] - 1
		} else if runeTab[pipe[0]][pipe[1]-1] == '*' {
			newPipe[0] = pipe[0] - 1
			newPipe[1] = pipe[1]
		} else {
			panic("On ne trouve pas où aller ensuite")
		}

	} else if runeTab[pipe[0]][pipe[1]] == '7' {

		if runeTab[pipe[0]+1][pipe[1]] == '*' {
			newPipe[0] = pipe[0]
			newPipe[1] = pipe[1] - 1
		} else if runeTab[pipe[0]][pipe[1]-1] == '*' {
			newPipe[0] = pipe[0] + 1
			newPipe[1] = pipe[1]
		} else {
			panic("On ne trouve pas où aller ensuite")
		}

	} else if runeTab[pipe[0]][pipe[1]] == 'F' {

		if runeTab[pipe[0]+1][pipe[1]] == '*' {
			newPipe[0] = pipe[0]
			newPipe[1] = pipe[1] + 1
		} else if runeTab[pipe[0]][pipe[1]+1] == '*' {
			newPipe[0] = pipe[0] + 1
			newPipe[1] = pipe[1]
		} else {
			panic("On ne trouve pas où aller ensuite")
		}

	} else {
		panic("Problème de valeur de 'argument")
	}

	runeTab[pipe[0]][pipe[1]] = '*'

	return newPipe
}
