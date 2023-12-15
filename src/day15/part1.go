package main

import (
	"strings"
)

func Part1(lines []string) int {
	words := strings.Split(lines[0], ",")
	sum := 0
	for _, word := range words {
		sum += hash(word)
	}
	return sum
}

func hash(word string) int {
	value := 0
	for _, char := range word {
		value += int(char)
		value *= 17
		value = value % 256
	}
	return value
}
