package stacks

import "github.com/akitasoftware/go-utils/optionals"

type Stack[T any] interface {
	// Adds an element to the top of the stack.
	Push(T)

	// Removes an element from the top of the stack. Returns None if the stack is
	// empty.
	Pop() optionals.Optional[T]

	// Returns (but does not remove) the element at the top of the stack. Returns
	// None if the stack is empty.
	Peek() optionals.Optional[T]

	// Returns true if the stack is empty.
	IsEmpty() bool

	// Returns the number of elements on the stack.
	Size() int

	// Calls the given function with each element in the stack, from top to
	// bottom.
	ForEach(func(T))
}

// Returns a new stack containing the given elements. Equivalent to creating an
// empty stack and calling Push with each element in turn.
func NewStack[T any](elements ...T) Stack[T] {
	return NewSliceStack(elements...)
}
