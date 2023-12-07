package main

import (
	"advent_of_code_2023/src/utilities"
	"testing"
)

func TestPart1Example(t *testing.T) {
	path := "input_day3_test"
	lines := utilities.ReadLines(path)

	res := Part1(lines)
	exp := 4361
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}

func TestPart1(t *testing.T) {
	path := "input_day3"
	lines := utilities.ReadLines(path)
	res := Part1(lines)
	exp := 533784
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}
func TestPart2Example(t *testing.T) {
	path := "input_day3_test"
	lines := utilities.ReadLines(path)

	res := Part2(lines)
	exp := 467835
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}
func TestPart2(t *testing.T) {
	path := "input_day3"
	lines := utilities.ReadLines(path)

	res := Part2(lines)
	exp := 78826761
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}

func BenchmarkPart1(b *testing.B) {
	path := "input_day3"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	path := "input_day3"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part2(lines)
	}
}
