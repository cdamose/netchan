// This example demonstrates a priority queue built using the heap interface.
package lib

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value interface{} // The value of the item; arbitrary.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index    int // The index of the item in the heap.
	sourceId int // payload containing the value's source id
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	items    []*Item
	lessFunc func(a, b interface{}) bool
}

func NewPriorityQueue(lessFunc func(a, b interface{}) bool) *PriorityQueue {
	pq := &PriorityQueue{}
	pq.items = make([]*Item, 0)
	pq.lessFunc = lessFunc
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue) Enqueue(x interface{}, sourceId int) {
	heap.Push(pq, &Item{value: x, sourceId: sourceId})
}

func (pq *PriorityQueue) Dequeue() (interface{}, int) {
	item := heap.Pop(pq).(*Item)
	return item.value, item.sourceId
}

func (pq *PriorityQueue) Len() int { return len(pq.items) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.lessFunc(pq.items[i].value, pq.items[j].value)
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(*Item)
	item.index = n
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	pq.items = old[0 : n-1]
	return item
}
