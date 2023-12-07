package utilities

func IsInTwo(list []int, x int) bool {
	count := 0
	for i := range list {
		if x == list[i] {
			count++
			if count == 2 {
				return true
			}
		}
	}
	return false
}
