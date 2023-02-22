package stacks

import "github.com/akitasoftware/go-utils/optionals"

type SliceStack[T any] struct {
	elements []T
}

func NewSliceStack[T any](elements ...T) *SliceStack[T] {
	return &SliceStack[T]{
		elements: elements,
	}
}

func (stack *SliceStack[T]) Push(element T) {
	stack.elements = append(stack.elements, element)
}

func (stack *SliceStack[T]) Pop() optionals.Optional[T] {
	var result optionals.Optional[T]
	if stack.Size() > 0 {
		result = optionals.Some(stack.elements[len(stack.elements)-1])
		stack.elements = stack.elements[:len(stack.elements)-1]
	}
	return result
}

func (stack *SliceStack[T]) Peek() optionals.Optional[T] {
	var result optionals.Optional[T]
	if stack.Size() > 0 {
		result = optionals.Some(stack.elements[len(stack.elements)-1])
	}
	return result
}

func (stack *SliceStack[T]) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *SliceStack[T]) Size() int {
	return len(stack.elements)
}

func (stack *SliceStack[T]) ForEach(f func(T)) {
	for idx := len(stack.elements) - 1; idx >= 0; idx-- {
		f(stack.elements[idx])
	}
}
