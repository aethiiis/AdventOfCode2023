package main

import (
	"advent_of_code_2023/src/utilities"
	"testing"
)

func TestPart1Example(t *testing.T) {
	path := "input_day6_test"
	lines := utilities.ReadLines(path)

	res := Part1(lines)
	exp := 288
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}

func TestPart1(t *testing.T) {
	path := "input_day6"
	lines := utilities.ReadLines(path)
	res := Part1(lines)
	exp := 588588
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}
func TestPart2Example(t *testing.T) {
	path := "input_day6_test"
	lines := utilities.ReadLines(path)

	res := Part2(lines)
	exp := 71503
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}
func TestPart2(t *testing.T) {
	path := "input_day6"
	lines := utilities.ReadLines(path)

	res := Part2(lines)
	exp := 34655848
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}

func BenchmarkPart1(b *testing.B) {
	path := "input_day6"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	path := "input_day6"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part2(lines)
	}
}
