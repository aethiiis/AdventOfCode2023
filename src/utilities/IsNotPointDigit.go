package utilities

import "unicode"

func IsNotPointDigit(char rune) bool {
	if char == '.' {
		return false
	} else if unicode.IsDigit(char) {
		return false
	} else {
		return true
	}
}
