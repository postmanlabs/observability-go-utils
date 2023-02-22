package stacks

import (
	"fmt"
	"testing"

	"github.com/akitasoftware/go-utils/slices"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
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
		for i, q := range []Stack[int]{NewSliceStack[int]()} {
			label := func(msg string) string {
				return fmt.Sprintf("%s - %d: %s", tc.name, i, msg)
			}

			assert.True(t, q.IsEmpty(), label("empty"))

			for _, v := range tc.input {
				q.Push(v)
			}
			reversedInput := slices.Reverse(tc.input)

			assert.Equal(t, len(tc.input), q.Size(), label("length"))

			if len(tc.input) > 0 {
				assert.False(t, q.IsEmpty(), label("not empty"))
			}

			// Test ForEach. It should behave as if we had popped each element.
			foreachOutput := make([]int, 0, q.Size())
			q.ForEach(func(v int) {
				foreachOutput = append(foreachOutput, v)
			})
			assert.Equal(t, reversedInput, foreachOutput, label("foreach output"))

			// Pop all elements, and check that the resulting list is equal to the
			// input, reversed.
			output := make([]int, q.Size())
			for i := range output {
				if v, exists := q.Pop().Get(); exists {
					output[i] = v
				}
			}

			assert.Equal(t, reversedInput, output, label("output"))
		}
	}
}
