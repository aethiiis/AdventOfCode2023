package day6

import "testing"

func TestPart1Example(t *testing.T) {
	res := Part1("input_day6_test")
	exp := 288
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}

func TestPart1(t *testing.T) {
	res := Part1("input_day6")
	exp := 588588
	if res != exp {
		t.Errorf("Mauvaise valeur : on s'attendait à %d mais on a obtenu %d.\n", exp, res)
	}
}
