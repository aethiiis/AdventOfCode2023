package utilities

func ConvertirTableauRunesListeStrings(runes [][]rune) []string {
	if len(runes) == 0 {
		return nil
	}

	lines := make([]string, len(runes))
	for i, listRunes := range runes {
		lines[i] = string(listRunes)
	}
	return lines
}
