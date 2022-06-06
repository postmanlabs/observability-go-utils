package queues

import "github.com/akitasoftware/go-utils/optionals"

// A FIFO queue.
//
// XXX Current implementation should be revisited: may not have great memory
// performance.
type Queue[T any] struct {
	elements []T
}

// Returns a new FIFO queue containing the given elements. Equivalent to
// creating an empty queue and calling Enqueue with each element in turn.
func NewQueue[T any](elements ...T) *Queue[T] {
	return &Queue[T]{
		elements: elements,
	}
}

func (q *Queue[T]) Enqueue(t T) {
	q.elements = append(q.elements, t)
}

func (q *Queue[T]) Dequeue() optionals.Optional[T] {
	if len(q.elements) == 0 {
		return optionals.None[T]()
	}
	result := q.elements[0]
	q.elements = q.elements[1:]
	return optionals.Some(result)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}
