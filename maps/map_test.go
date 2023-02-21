package maps

import (
	"sort"
	"testing"

	"github.com/akitasoftware/go-utils/math"
	"github.com/stretchr/testify/assert"
)

func TestBasicOps(t *testing.T) {
	m := Map[string, int]{}

	m.Upsert("foo", 1, math.Add[int])
	assert.Equal(t, Map[string, int]{"foo": 1}, m)

	m.Add(Map[string, int]{"foo": 2, "bar": 1}, math.Add[int])
	assert.Equal(t, Map[string, int]{"foo": 3, "bar": 1}, m)

	sortedKeys := m.Keys()
	sort.Strings(sortedKeys)
	assert.Equal(t, []string{"bar", "foo"}, sortedKeys)

	sortedValues := m.Values()
	sort.Ints(sortedValues)
	assert.Equal(t, []int{1, 3}, sortedValues)
}
