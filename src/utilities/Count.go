package utilities

func Count(element rune, liste []rune) int {
	count := 0
	for i := range liste {
		if element == liste[i] {
			count++
		}
	}
	return count
}
