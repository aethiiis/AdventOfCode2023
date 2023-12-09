package utilities

func EstVide(list []int) bool {
	for i := range list {
		if list[i] != 0 {
			return false
		}
	}
	return true
}
