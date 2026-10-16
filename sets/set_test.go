package sets

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/akitasoftware/go-utils/optionals"
	"github.com/stretchr/testify/assert"
)

func TestBasicSetOperations(t *testing.T) {
	s := NewSet[int]()
	assert.Equal(t, len(s), 0)
	assert.Equal(t, map[int]struct{}(s), map[int]struct{}{})

	s.Insert(1)
	assert.Equal(t, s, NewSet(1))

	s.Intersect(NewSet(1, 2))
	assert.Equal(t, s, NewSet(1))

	s.Union(NewSet(1, 2))
	assert.Equal(t, s, NewSet(1, 2))

	s.Delete(1)
	assert.Equal(t, s, NewSet(2))
}

func TestSetJson(t *testing.T) {
	s := NewSet(3, 2, 1)

	bs, err := json.Marshal(s)
	assert.NoError(t, err)

	var deserialized Set[int]
	err = json.Unmarshal(bs, &deserialized)
	assert.NoError(t, err)

	assert.Equal(t, deserialized, s, "s == unmarshal(marshal(s))")
}

func TestSetIntersect(t *testing.T) {
	testCases := []struct {
		name     string
		sets     []Set[int]
		expected Set[int]
	}{
		{
			name:     "empty",
			sets:     nil,
			expected: NewSet[int](),
		},
		{
			name:     "overlap",
			sets:     []Set[int]{NewSet(1, 2), NewSet(2, 3)},
			expected: NewSet(2),
		},
		{
			name:     "no overlap",
			sets:     []Set[int]{NewSet(1, 2), NewSet(3, 4)},
			expected: NewSet[int](),
		},
	}

	for _, tc := range testCases {
		intersected := Intersect(tc.sets...)
		assert.Equal(t, tc.expected, intersected, tc.name)
	}
}

func TestSetFromSlice(t *testing.T) {
	slice := []string{"0x1", "0x2", "0x3", "0xf", "xyz"}

	e1 := NewSet[int](0, 1, 2, 3, 15)
	s1 := FromSlice(slice, func(s string) int {
		v, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return 0
		} else {
			return int(v)
		}
	})
	assert.Equal(t, e1, s1)

	e2 := NewSet[int](1, 2, 3)
	s2 := FromFilteredSlice(slice, func(s string) optionals.Optional[int] {
		v, err := strconv.ParseInt(s, 0, 0)
		if err != nil || v > 10 {
			return optionals.None[int]()
		} else {
			return optionals.Some(int(v))
		}
	})
	assert.Equal(t, e2, s2)
}
