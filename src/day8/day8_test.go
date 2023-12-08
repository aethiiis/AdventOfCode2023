package main

import (
	"advent_of_code_2023/src/utilities"
	"testing"
)

func TestPart1Example1(t *testing.T) {
	path := "input_day8_test1"
	lines := utilities.ReadLines(path)

	res := Part1(lines)
	exp := 2
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}

func TestPart1Example2(t *testing.T) {
	path := "input_day8_test2"
	lines := utilities.ReadLines(path)

	res := Part1(lines)
	exp := 6
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}

func TestPart1(t *testing.T) {
	path := "input_day8"
	lines := utilities.ReadLines(path)
	res := Part1(lines)
	exp := 20093
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}
func TestPart2Example3(t *testing.T) {
	path := "input_day8_test3"
	lines := utilities.ReadLines(path)

	var res64 int64
	res64 = Part2(lines).Int64()
	var exp64 int64
	exp64 = 6

	if res64 != exp64 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp64, res64)
	}
}

func TestPart2(t *testing.T) {
	path := "input_day8"
	lines := utilities.ReadLines(path)

	var res64 int64
	res64 = Part2(lines).Int64()
	var exp64 int64
	exp64 = 22103062509257

	if res64 != exp64 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp64, res64)
	}
}

func BenchmarkPart1(b *testing.B) {
	path := "input_day8"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	path := "input_day8"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part2(lines)
	}
}
