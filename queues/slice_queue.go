package queues

import "github.com/akitasoftware/go-utils/optionals"

// A FIFO queue.
//
// XXX Current implementation should be revisited: may not have great memory
// performance.
type SliceQueue[T any] struct {
	elements []T
}

// Returns a new FIFO queue containing the given elements. Equivalent to
// creating an empty queue and calling Enqueue with each element in turn.
func NewSliceQueue[T any](elements ...T) Queue[T] {
	return &SliceQueue[T]{
		elements: elements,
	}
}

func (q *SliceQueue[T]) Enqueue(t T) {
	q.elements = append(q.elements, t)
}

func (q *SliceQueue[T]) Dequeue() optionals.Optional[T] {
	if len(q.elements) == 0 {
		return optionals.None[T]()
	}
	result := q.elements[0]
	q.elements = q.elements[1:]
	return optionals.Some(result)
}

func (q *SliceQueue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *SliceQueue[T]) Size() int {
	return len(q.elements)
}

func (q *SliceQueue[T]) Peek() optionals.Optional[T] {
	if q.IsEmpty() {
		return optionals.None[T]()
	}
	return optionals.Some(q.elements[0])
}

func (q *SliceQueue[T]) ForEach(f func(T)) {
	for _, v := range q.elements {
		f(v)
	}
}
