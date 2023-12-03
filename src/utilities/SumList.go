package utilities

func SumList(list []int) int {
	somme := 0
	for _, element := range list {
		somme += element
	}
	return somme
}
