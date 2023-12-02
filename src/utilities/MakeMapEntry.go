package utilities

import (
	"fmt"
	"strings"
)

func MakeMapEntry(resultat map[string][][]string, input string) {

	// Diviser la chaîne par ":"
	parts := strings.SplitN(input, ":", 2)
	if len(parts) != 2 {
		fmt.Println("Format de chaîne incorrect.")
	}

	// Extraire la clé et la valeur
	cle := strings.TrimSpace(parts[0])
	valeurBrute := strings.TrimSpace(parts[1])

	// Diviser la valeur par ";"
	listeGrandes := strings.Split(valeurBrute, ";")

	// Nettoyer les espaces des éléments de la grande liste
	for i, grande := range listeGrandes {
		listeGrandes[i] = strings.TrimSpace(grande)

		// Diviser chaque grande liste par ","
		listePetites := strings.Split(grande, ",")

		// Nettoyer les espaces des éléments de chaque petite liste
		for j, petite := range listePetites {
			listePetites[j] = strings.TrimSpace(petite)
		}

		// Ajouter chaque petite liste à la grande liste
		resultat[cle] = append(resultat[cle], listePetites)
	}
}
