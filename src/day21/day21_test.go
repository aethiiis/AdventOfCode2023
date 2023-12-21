package main

import (
	"advent_of_code_2023/src/utilities"
	"testing"
)

func TestPart1Example(t *testing.T) {
	path1 := "input_day21_test"
	lines1 := utilities.ReadLines(path1)

	res1 := Part1(lines1)
	exp1 := 11687500
	if res1 != exp1 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp1, res1)
	}
}

func TestPart1(t *testing.T) {
	path := "input_day21"
	lines := utilities.ReadLines(path)
	res := Part1(lines)
	exp := 812609846
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}
func TestPart2(t *testing.T) {
	path := "input_day21"
	lines := utilities.ReadLines(path)

	res := Part2(lines)
	exp := 245114020323037
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}

func BenchmarkPart1(b *testing.B) {
	path := "input_day21"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	path := "input_day21"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part2(lines)
	}
}
