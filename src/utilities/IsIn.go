package utilities

func IsIn(list []int, x int) bool {
	for i := range list {
		if x == list[i] {
			return true
		}
	}
	return false
}
