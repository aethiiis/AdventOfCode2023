package main

import (
	"container/heap"
)

type Lava struct {
	Cost int
	X    int
	Y    int
	dX   int
	dY   int
	Con  int
}

type Cas struct {
	X   int
	Y   int
	dX  int
	dY  int
	Con int
}

// PriorityQueue est une PriorityQueue pour la struct Lava
type PriorityQueue []*Lava

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Lava)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Part1(lines []string) int {
	parcourus := make(map[Cas]bool)
	frontier := make(PriorityQueue, 0)
	heap.Init(&frontier)
	start := Lava{Cost: 0, X: 0, Y: 0, dX: 0, dY: 0, Con: 0}
	heap.Push(&frontier, &start)

	for frontier.Len() != 0 {
		current := heap.Pop(&frontier).(*Lava)
		cost := current.Cost
		x := current.X
		y := current.Y
		dx := current.dX
		dy := current.dY
		con := current.Con
		//fmt.Printf("%d %d %d %d %d %d\n", cost, x, y, dx, dy, con)

		if x == len(lines)-1 && y == len(lines[0])-1 {
			return cost
		}

		currentCas := Cas{
			X:   x,
			Y:   y,
			dX:  dx,
			dY:  dy,
			Con: con,
		}
		if _, ok := parcourus[currentCas]; ok {
			continue
		}

		parcourus[currentCas] = true

		if con < 3 && (dx != 0 || dy != 0) {
			nextX := x + dx
			nextY := y + dy
			if nextX >= 0 && nextX < len(lines) && nextY >= 0 && nextY < len(lines[0]) {
				next := Lava{
					Cost: cost + int(lines[nextX][nextY]) - 48,
					X:    nextX,
					Y:    nextY,
					dX:   dx,
					dY:   dy,
					Con:  con + 1,
				}
				heap.Push(&frontier, &next)
			}
		}

		directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, direction := range directions {
			nextDx := direction[0]
			nextDy := direction[1]
			if (nextDx != dx || nextDy != dy) && (nextDx != -dx || nextDy != -dy) {
				nextX := x + nextDx
				nextY := y + nextDy
				if nextX >= 0 && nextX < len(lines) && nextY >= 0 && nextY < len(lines[0]) {
					next := Lava{
						Cost: cost + int(lines[nextX][nextY]) - 48,
						X:    nextX,
						Y:    nextY,
						dX:   nextDx,
						dY:   nextDy,
						Con:  1,
					}
					heap.Push(&frontier, &next)
				}
			}
		}
	}
	return 0
}
