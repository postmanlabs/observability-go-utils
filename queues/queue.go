package queues

import (
	"github.com/akitasoftware/go-utils/optionals"
)

// A FIFO queue.
type Queue[T any] interface {
	// Adds an element to the tail of the queue.
	Enqueue(T)

	// Removes and returns an element from the head of the queue.  Returns None
	// if the queue is empty.
	Dequeue() optionals.Optional[T]

	// Returns (but does not remove) an element from the head of the queue.
	// Returns None if the queue is empty.
	Peek() optionals.Optional[T]

	// Returns true if the queue is empty.
	IsEmpty() bool

	// Returns the length of the queue.
	Size() int

	// Iterates through the queue from head to tail.
	ForEach(func(T))
}

// Returns a new FIFO queue containing the given elements. Equivalent to
// creating an empty queue and calling Enqueue with each element in turn.
func NewQueue[T any](elements ...T) Queue[T] {
	return NewLinkedListQueue(elements...)
}
