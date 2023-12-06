package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

func ConcatenerNombres(liste []string) int {
	// Concaténer les chaînes de caractères
	resultatStr := strings.Join(liste, "")

	// Convertir la chaîne résultante en entier
	resultat, err := strconv.Atoi(resultatStr)
	if err != nil {
		fmt.Println("Erreur de conversion :", err)
	}

	return resultat
}
