package utilities

func MapMaker(lines []string) map[string][][]string {
	str := ""
	resultat := make(map[string][][]string)

	for i := 0; i < len(lines); i++ {
		str = lines[i]
		MakeMapEntry(resultat, str)
	}
	return resultat
}
