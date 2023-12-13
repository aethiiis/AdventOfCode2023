package utilities

import "strings"

func StringsIdentiquesAUnCaracterePres(s1, s2 string) (bool, int) {
	// Vérifie si les chaînes ont la même longueur
	if len(s1) != len(s2) {
		return false, -1
	}

	// Compte le nombre de caractères différents
	diffCount := 0
	index := -2
	for i := 0; i < len(s1); i++ {
		if !strings.EqualFold(string(s1[i]), string(s2[i])) {
			diffCount++
			index = i
		}
	}

	// Retourne true si le nombre de caractères différents est au plus 1
	if diffCount == 0 {
		return false, -2
	}
	return diffCount == 1, index
}
