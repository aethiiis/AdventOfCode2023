package utilities

func ConvertirListeStringsTableauRunes(listeDeChaines []string) [][]rune {
	// Vérifier que la liste n'est pas vide
	if len(listeDeChaines) == 0 {
		return nil
	}

	// Longueur des chaînes dans la liste
	nbLigne := len(listeDeChaines)
	nbColonne := len(listeDeChaines[0])

	// Initialiser la liste de listes de runes
	listeDeListes := make([][]rune, nbLigne)

	// Parcourir chaque chaîne dans la liste
	for i, chaine := range listeDeChaines {
		// Convertir la chaîne en une liste de runes
		runes := []rune(chaine)

		// Vérifier que la longueur de la chaîne est correcte
		if len(runes) != nbColonne {
			panic("Les chaînes n'ont pas la même longueur.")
		}

		// Ajouter chaque rune à la liste correspondante
		listeDeListes[i] = runes
	}

	return listeDeListes
}
