package utilities

func Sum2ProductList(list []int) int {
	somme := 0
	for i := 0; i < len(list); i += 2 {
		somme += list[i] * list[i+1]
	}
	return somme
}
