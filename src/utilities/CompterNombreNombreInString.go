package utilities

import "unicode"

func CompterNombreNombreInString(str string) int {
	nombreDeNombres := 0
	enNombre := false

	for _, char := range str {
		if unicode.IsDigit(char) {
			// Si le caractère est un chiffre, et que nous ne sommes pas déjà en train de compter un nombre
			if !enNombre {
				nombreDeNombres++
				enNombre = true
			}
		} else {
			// Si le caractère n'est pas un chiffre, indiquer que nous ne sommes pas en train de compter un nombre
			enNombre = false
		}
	}

	return nombreDeNombres
}
