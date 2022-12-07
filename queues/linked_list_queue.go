package queues

import (
	"container/list"

	"github.com/akitasoftware/go-utils/optionals"
)

type LinkedListQueue[T any] struct {
	queue *list.List
}

func NewLinkedListQueue[T any](elements ...T) Queue[T] {
	rv := &LinkedListQueue[T]{queue: list.New()}

	for _, v := range elements {
		rv.Enqueue(v)
	}

	return rv
}

func (q *LinkedListQueue[T]) Enqueue(v T) {
	q.queue.PushBack(v)
}

func (q *LinkedListQueue[T]) Dequeue() optionals.Optional[T] {
	if q.queue.Len() == 0 {
		return optionals.None[T]()
	}

	return optionals.Some(q.queue.Remove(q.queue.Front()).(T))
}

func (q *LinkedListQueue[T]) Peek() optionals.Optional[T] {
	if elt := q.queue.Front(); elt == nil {
		return optionals.None[T]()
	} else {
		return optionals.Some((elt.Value).(T))
	}
}

func (q *LinkedListQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *LinkedListQueue[T]) Size() int {
	return q.queue.Len()
}

func (q *LinkedListQueue[T]) ForEach(f func(T)) {
	for element := q.queue.Front(); element != nil; element = element.Next() {
		f(element.Value.(T))
	}
}
