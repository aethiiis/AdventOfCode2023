package main

import (
	"strconv"
	"strings"
)

type Lens struct {
	Label string
	Focal int
}

func Part2(lines []string) int {
	var boxes map[int][]Lens
	boxes = make(map[int][]Lens)
	words := strings.Split(lines[0], ",")

	for _, word := range words {
		numero, label, signe, focal := instruction(word)
		if signe {
			equal(&boxes, numero, label, focal)
		} else {
			dash(&boxes, numero, label)
		}

	}
	return power(boxes)
}

func instruction(word string) (int, string, bool, int) {
	var focal int
	if strings.Count(word, "=") == 1 {
		tmp := strings.Split(word, "=")
		focal, _ = strconv.Atoi(tmp[1])
		return hash(tmp[0]), tmp[0], true, focal
	} else if strings.Count(word, "-") == 1 {
		return hash(word[:len(word)-1]), word[:len(word)-1], false, -1
	} else {
		panic("Erreur dans le mot fourni")
	}
}

func dash(boxes *map[int][]Lens, numero int, label string) {
	liste, ok := (*boxes)[numero]
	if !ok {
		return
	} else {
		for i := 0; i < len(liste); i++ {
			if liste[i].Label == label {
				liste = append(liste[:i], liste[i+1:]...)
				(*boxes)[numero] = liste
				break
			}
		}
	}
}

func equal(boxes *map[int][]Lens, numero int, label string, focal int) {
	liste, ok := (*boxes)[numero]
	if !ok {
		(*boxes)[numero] = []Lens{{Label: label, Focal: focal}}
	} else {
		present := false
		for i := 0; i < len(liste); i++ {
			if liste[i].Label == label {
				liste[i] = Lens{Label: label, Focal: focal}
				present = true
				break
			}
		}
		if !present {
			liste = append(liste, Lens{Label: label, Focal: focal})
		}
		(*boxes)[numero] = liste
	}
}

func power(boxes map[int][]Lens) int {
	pow := 0
	for key, value := range boxes {
		for pos, lens := range value {
			pow += (key + 1) * (pos + 1) * lens.Focal
		}
	}
	return pow
}
