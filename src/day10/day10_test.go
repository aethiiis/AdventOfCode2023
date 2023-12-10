package main

import (
	"advent_of_code_2023/src/utilities"
	"testing"
)

func TestPart1Example(t *testing.T) {
	path1 := "input_day10_test1"
	lines1 := utilities.ReadLines(path1)

	res1 := Part1(lines1)
	exp1 := 4
	if res1 != exp1 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp1, res1)
	}

	path2 := "input_day10_test2"
	lines2 := utilities.ReadLines(path2)

	res2 := Part1(lines2)
	exp2 := 8
	if res2 != exp2 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp2, res2)
	}
}

func TestPart1(t *testing.T) {
	path := "input_day10"
	lines := utilities.ReadLines(path)
	res := Part1(lines)
	exp := 6733
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}
func TestPart2Example(t *testing.T) {
	path3 := "input_day10_test3"
	lines3 := utilities.ReadLines(path3)

	res3 := Part2(lines3)
	exp3 := 4
	if res3 != exp3 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp3, res3)
	}

	path4 := "input_day10_test4"
	lines4 := utilities.ReadLines(path4)

	res4 := Part2(lines4)
	exp4 := 4
	if res4 != exp4 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp4, res4)
	}

	path5 := "input_day10_test5"
	lines5 := utilities.ReadLines(path5)

	res5 := Part2(lines5)
	exp5 := 8
	if res5 != exp5 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp5, res5)
	}

	path6 := "input_day10_test6"
	lines6 := utilities.ReadLines(path6)

	res6 := Part2(lines6)
	exp6 := 10
	if res6 != exp6 {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp6, res6)
	}
}
func TestPart2(t *testing.T) {
	path := "input_day10"
	lines := utilities.ReadLines(path)

	res := Part2(lines)
	exp := 435
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}

func BenchmarkPart1(b *testing.B) {
	path := "input_day10"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	path := "input_day10"
	lines := utilities.ReadLines(path)

	for i := 0; i < b.N; i++ {
		Part2(lines)
	}
}
