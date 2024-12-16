package ds

import (
	"container/heap"
)

// priorityQueue is an implementation of the heap.Interface.
type priorityQueue[T any] struct {
	data []T
	cmp  func(a, b T) bool
}

func NewPriorityQueue[T any](cmp func(a, b T) bool) heap.Interface {
	return &priorityQueue[T]{
		data: nil,
		cmp:  cmp,
	}
}

func (pq *priorityQueue[T]) Len() int {
	return len(pq.data)
}

func (pq *priorityQueue[T]) Less(i, j int) bool {
	return pq.cmp(pq.data[i], pq.data[j])
}

func (pq *priorityQueue[T]) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *priorityQueue[T]) Push(x any) {
	pq.data = append(pq.data, x.(T))
}

func (pq *priorityQueue[T]) Pop() any {
	n := pq.Len()
	item := pq.data[n-1]
	pq.data = pq.data[:n-1]
	return item
}
