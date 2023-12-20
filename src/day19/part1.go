package main

import (
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Rule struct {
	cond string
	dest string
}

func Part1(lines []string) int {
	workflows, parts := traitement(lines)
	sum := 0
	for _, part := range parts {
		next := parcours(part, workflows["in"])
		for next != "A" && next != "R" {
			next = parcours(part, workflows[next])
		}
		if next == "A" {
			sum += part.x + part.m + part.a + part.s
		}
	}
	return sum
}

func traitement(lines []string) (map[string][]Rule, []Part) {
	var lineSep int
	for i, line := range lines {
		if len(line) == 0 {
			lineSep = i
			break
		}
	}
	workflowsList := lines[:lineSep]
	partsList := lines[lineSep+1:]
	dict := make(map[string][]Rule)
	for i := range workflowsList {
		nameWorkflow := strings.Split(workflowsList[i], "{")
		name := nameWorkflow[0]
		workflow := strings.Trim(nameWorkflow[1], "}")
		workflowSep := strings.Split(workflow, ",")
		finalWorkflow := make([]Rule, len(workflowSep))
		for j := range workflowSep {
			tmp := strings.Split(workflowSep[j], ":")
			if len(tmp) == 2 {
				finalWorkflow[j] = Rule{
					cond: tmp[0],
					dest: tmp[1],
				}
			} else if len(tmp) == 1 {
				finalWorkflow[j] = Rule{
					dest: tmp[0],
				}
			} else {
				panic("Problème de format")
			}
		}
		dict[name] = finalWorkflow
	}
	parts := make([]Part, len(partsList))
	for i := range parts {
		nums := strings.Split(strings.Trim(partsList[i], "{}"), ",")
		x, _ := strconv.Atoi(strings.Split(nums[0], "=")[1])
		m, _ := strconv.Atoi(strings.Split(nums[1], "=")[1])
		a, _ := strconv.Atoi(strings.Split(nums[2], "=")[1])
		s, _ := strconv.Atoi(strings.Split(nums[3], "=")[1])
		parts[i] = Part{
			x: x,
			m: m,
			a: a,
			s: s,
		}
	}
	return dict, parts
}

func parcours(part Part, workflow []Rule) string {
	for _, rule := range workflow {
		// Cas de fin
		if rule.cond == "" {
			return rule.dest
		}
		symbol := rule.cond[0]
		inequality := rule.cond[1]
		number, _ := strconv.Atoi(rule.cond[2:])
		if symbol == 'x' {
			if inequality == '<' {
				if part.x < number {
					return rule.dest
				}
			} else if inequality == '>' {
				if part.x > number {
					return rule.dest
				}
			} else {
				panic("Problème de format")
			}
		} else if symbol == 'm' {
			if inequality == '<' {
				if part.m < number {
					return rule.dest
				}
			} else if inequality == '>' {
				if part.m > number {
					return rule.dest
				}
			} else {
				panic("Problème de format")
			}
		} else if symbol == 's' {
			if inequality == '<' {
				if part.s < number {
					return rule.dest
				}
			} else if inequality == '>' {
				if part.s > number {
					return rule.dest
				}
			} else {
				panic("Problème de format")
			}
		} else if symbol == 'a' {
			if inequality == '<' {
				if part.a < number {
					return rule.dest
				}
			} else if inequality == '>' {
				if part.a > number {
					return rule.dest
				}
			} else {
				panic("Problème de format")
			}
		} else {
			panic("Problème de formatage")
		}
	}
	return ""
}
