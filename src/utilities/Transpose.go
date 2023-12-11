package utilities

import "strings"

func Transpose(listeOrigine []string) []string {
	nLignes := len(listeOrigine)
	nColonnes := len(listeOrigine[0])

	// Créer une nouvelle liste avec les dimensions inversées
	nouvelleListe := make([]string, nColonnes)
	for i := range nouvelleListe {
		nouvelleListe[i] = strings.Repeat(" ", nLignes)
	}

	// Remplir la nouvelle liste avec les éléments transposés
	for i := 0; i < nLignes; i++ {
		for j := 0; j < nColonnes; j++ {
			nouvelleListe[j] = nouvelleListe[j][:i] + string(listeOrigine[i][j]) + nouvelleListe[j][i+1:]
		}
	}

	return nouvelleListe
}
