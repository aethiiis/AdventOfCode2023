package main

import (
	"advent_of_code_2023/src/utilities"
	"testing"
)

func TestPart1Example(t *testing.T) {
	path1 := "input_day13_test"
	lines1 := utilities.ReadLines(path1)

	res1 := Part1(lines1)
	exp1 := 405
	if res1 != exp1 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp1, res1)
	}
}

func TestPart1(t *testing.T) {
	path := "input_day13"
	lines := utilities.ReadLines(path)
	res := Part1(lines)
	exp := 43614
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}
func TestPart2Example(t *testing.T) {
	path3 := "input_day13_test"
	lines3 := utilities.ReadLines(path3)

	res3 := Part2(lines3)
	exp3 := 400
	if res3 != exp3 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp3, res3)
	}
}
func TestPart2(t *testing.T) {
	path := "input_day13"
	lines := utilities.ReadLines(path)

	res := Part2(lines)
	exp := 791134099634
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}

func BenchmarkPart1(b *testing.B) {
	path := "input_day13"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	path := "input_day13"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part2(lines)
	}
}
