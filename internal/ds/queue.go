package ds

// Queue is a simple FIFO queue.
type Queue[T any] []T

func (q Queue[T]) Push(v T) Queue[T] {
	return append(q, v)
}

func (q Queue[T]) Pop() (T, Queue[T]) {
	return q[0], q[1:]
}

func (q Queue[T]) IsEmpty() bool {
	return len(q) == 0
}
