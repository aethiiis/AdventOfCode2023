package utilities

import (
	"container/heap"
)

// Item represents an item in the priority queue.
type Item struct {
	Value    Pos // The value of the item.
	Priority int // The priority of the item.
	index    int // The index of the item in the heap.
}

// PriorityQueue implements a priority queue based on a min-heap.
type PriorityQueue []*Item

// Len returns the length of the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less returns true if the item at index i has higher priority than the item at index j.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

// Swap swaps the items at indices i and j.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the highest priority item from the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an item in the priority queue.
func (pq *PriorityQueue) update(item *Item, value Pos, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}
