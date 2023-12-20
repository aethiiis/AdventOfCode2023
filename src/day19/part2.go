package main

import (
	"strconv"
	"strings"
)

func Part2(lines []string) int {
	workflows := workflowsMaker(lines)
	// Construction de l'arbre
	var listePaths [][]string
	var listeNonPaths [][]string
	PrintPaths("in", []string{}, workflows, &listePaths, &listeNonPaths)
	sum := 0
	for _, path := range listePaths {
		more := count(path, workflows)
		sum += more
	}
	return sum
}

func PrintPaths(root string, path []string, workflows map[string][]Rule, listPaths *[][]string, listNonPaths *[][]string) {
	path = append(path, root)
	rules := workflows[root]
	for _, rule := range rules {
		if rule.dest == "A" {
			var pathEnd []string
			pathEnd = make([]string, len(path))
			for i := range path {
				pathEnd[i] = path[i]
			}
			pathEnd = append(pathEnd, rule.dest)
			*listPaths = append(*listPaths, pathEnd)
		} else if rule.dest == "R" {
			var pathEnd []string
			pathEnd = make([]string, len(path))
			for i := range path {
				pathEnd[i] = path[i]
			}
			pathEnd = append(pathEnd, rule.dest)
			*listNonPaths = append(*listNonPaths, pathEnd)
		} else {
			PrintPaths(rule.dest, path, workflows, listPaths, listNonPaths)
		}
	}
}

func count(path []string, workflows map[string][]Rule) int {
	intervalleX := []int{1, 4000}
	intervalleM := []int{1, 4000}
	intervalleA := []int{1, 4000}
	intervalleS := []int{1, 4000}
	var conditions []string
	for i, step := range path[:len(path)-1] {
		rules := workflows[step]
		for j, rule := range rules {
			if rule.cond == "" {
				continue
			} else if rule.dest == path[i+1] {
				conditions = append(conditions, rule.cond)
				if path[i+1] == "A" {
					rule = Rule{cond: rule.cond, dest: "A*"}
					rules[j] = rule
					workflows[step] = rules
				}
				if path[i+1] == "R" {
					rule = Rule{cond: rule.cond, dest: "R*"}
					rules[j] = rule
					workflows[step] = rules
				}
				break
			} else {
				conditions = append(conditions, "!"+rule.cond)
			}
		}
	}
	for _, cond := range conditions {
		if cond[0] == '!' {
			symbol := cond[1]
			inequality := cond[2]
			number, _ := strconv.Atoi(cond[3:])
			if symbol == 'x' {
				if inequality == '<' {
					if intervalleX[0] <= number {
						intervalleX[0] = number
					}
				} else if inequality == '>' {
					if intervalleX[1] >= number {
						intervalleX[1] = number
					}
				} else {
					panic("Problème de format")
				}
			} else if symbol == 'm' {
				if inequality == '<' {
					if intervalleM[0] <= number {
						intervalleM[0] = number
					}
				} else if inequality == '>' {
					if intervalleM[1] >= number {
						intervalleM[1] = number
					}
				} else {
					panic("Problème de format")
				}
			} else if symbol == 's' {
				if inequality == '<' {
					if intervalleS[0] <= number {
						intervalleS[0] = number
					}
				} else if inequality == '>' {
					if intervalleS[1] >= number {
						intervalleS[1] = number
					}
				} else {
					panic("Problème de format")
				}
			} else if symbol == 'a' {
				if inequality == '<' {
					if intervalleA[0] <= number {
						intervalleA[0] = number
					}
				} else if inequality == '>' {
					if intervalleA[1] >= number {
						intervalleA[1] = number
					}
				} else {
					panic("Problème de format")
				}
			} else {
				panic("Problème de formatage")
			}
		} else {
			symbol := cond[0]
			inequality := cond[1]
			number, _ := strconv.Atoi(cond[2:])
			if symbol == 'x' {
				if inequality == '<' {
					if intervalleX[1] > number {
						intervalleX[1] = number - 1
					}
				} else if inequality == '>' {
					if intervalleX[0] < number {
						intervalleX[0] = number + 1
					}
				} else {
					panic("Problème de format")
				}
			} else if symbol == 'm' {
				if inequality == '<' {
					if intervalleM[1] > number {
						intervalleM[1] = number - 1
					}
				} else if inequality == '>' {
					if intervalleM[0] < number {
						intervalleM[0] = number + 1
					}
				} else {
					panic("Problème de format")
				}
			} else if symbol == 's' {
				if inequality == '<' {
					if intervalleS[1] > number {
						intervalleS[1] = number - 1
					}
				} else if inequality == '>' {
					if intervalleS[0] < number {
						intervalleS[0] = number + 1
					}
				} else {
					panic("Problème de format")
				}
			} else if symbol == 'a' {
				if inequality == '<' {
					if intervalleA[1] > number {
						intervalleA[1] = number - 1
					}
				} else if inequality == '>' {
					if intervalleA[0] < number {
						intervalleA[0] = number + 1
					}
				} else {
					panic("Problème de format")
				}
			} else {
				panic("Problème de formatage")
			}
		}
	}
	return (intervalleX[1] - intervalleX[0] + 1) * (intervalleM[1] - intervalleM[0] + 1) * (intervalleA[1] - intervalleA[0] + 1) * (intervalleS[1] - intervalleS[0] + 1)
}

func workflowsMaker(lines []string) map[string][]Rule {
	var lineSep int
	for i, line := range lines {
		if len(line) == 0 {
			lineSep = i
			break
		}
	}
	workflowsList := lines[:lineSep]
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
	return dict
}
