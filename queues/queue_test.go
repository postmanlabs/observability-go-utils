package queues

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "empty",
			input: []int{},
		},
		{
			name:  "singleton",
			input: []int{1},
		},
		{
			name:  "list",
			input: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tests {
		for i, q := range []Queue[int]{NewLinkedListQueue[int]()} {
			label := func(msg string) string {
				return fmt.Sprintf("%s - %d: %s", tc.name, i, msg)
			}

			assert.True(t, q.IsEmpty(), label("empty"))

			for _, v := range tc.input {
				q.Enqueue(v)
			}

			assert.Equal(t, len(tc.input), q.Size(), label("length"))

			if len(tc.input) > 0 {
				assert.False(t, q.IsEmpty(), label("not empty"))
			}

			// Test ForEach.  It should behave as if we had dequeued each
			// element.
			foreachOutput := make([]int, 0, q.Size())
			q.ForEach(func(v int) {
				foreachOutput = append(foreachOutput, v)
			})
			assert.Equal(t, tc.input, foreachOutput, label("foreach output"))

			// Dequeue all elements, and check that the resulting list is equal
			// to the input.
			output := make([]int, q.Size())
			for i := range output {
				if v, exists := q.Dequeue().Get(); exists {
					output[i] = v
				}
			}

			assert.Equal(t, tc.input, output, label("output"))
		}
	}
}
