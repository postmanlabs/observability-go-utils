package maps

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/akitasoftware/go-utils/math"
	"github.com/akitasoftware/go-utils/optionals"
	"github.com/stretchr/testify/assert"
)

func TestBasicComplexKeyMapOperations(t *testing.T) {
	m := ComplexKeyMap[[2]string, int]{}

	assert.True(t, m.IsEmpty())

	foo := [2]string{"foo", "bar"}
	bar := [2]string{"bar", "baz"}
	baz := [2]string{"baz", "qux"}
	qux := [2]string{"qux", "quux"}

	m.Upsert(foo, 1, math.Add[int])
	assert.False(t, m.IsEmpty())
	assert.Equal(t, ComplexKeyMap[[2]string, int]{foo: 1}, m)

	m.Add(ComplexKeyMap[[2]string, int]{foo: 2, bar: 1}, math.Add[int])
	assert.Equal(t, ComplexKeyMap[[2]string, int]{foo: 3, bar: 1}, m)

	m.Put(foo, 42)
	assert.Equal(t, ComplexKeyMap[[2]string, int]{foo: 42, bar: 1}, m)
	assert.Equal(t, m.Get(foo), optionals.Some(42))
	assert.Equal(t, m.GetOrValue(foo, 19), 42)
	assert.False(t, m.ContainsKey(baz))
	assert.Equal(t, m.GetOrValue(baz, 19), 19)
	assert.True(t, m.ContainsKey(baz))
	assert.Equal(t, m.Get(baz), optionals.Some(19))

	result, err := m.GetOrCompute(foo, func() (int, error) { return 37, fmt.Errorf("error") })
	assert.NoError(t, err)
	assert.Equal(t, result, 42)

	result, err = m.GetOrCompute(foo, func() (int, error) { return 37, nil })
	assert.NoError(t, err)
	assert.Equal(t, result, 42)

	result = m.GetOrComputeNoError(foo, func() int { return 37 })
	assert.Equal(t, result, 42)

	assert.False(t, m.ContainsKey(qux))

	_, err = m.GetOrCompute(qux, func() (int, error) { return 37, fmt.Errorf("error") })
	assert.Error(t, err)

	result, err = m.GetOrCompute(qux, func() (int, error) { return 37, nil })
	assert.NoError(t, err)
	assert.Equal(t, result, 37)
	assert.Equal(t, m.Get(qux), optionals.Some(37))

	assert.True(t, m.ContainsKey(qux))
	m.Delete(qux)
	assert.False(t, m.ContainsKey(qux))

	result = m.GetOrComputeNoError(qux, func() int { return 37 })
	assert.Equal(t, result, 37)
	assert.Equal(t, m.Get(qux), optionals.Some(37))
}

func TestComplexKeyMapJson(t *testing.T) {
	m := ComplexKeyMap[[2]string, int]{}
	m.Put([2]string{"foo", "bar"}, 3)
	m.Put([2]string{"bar", "baz"}, 2)
	m.Put([2]string{"baz", "qux"}, 1)

	bytes, err := json.Marshal(m)
	assert.NoError(t, err)

	var deserialized ComplexKeyMap[[2]string, int]
	err = json.Unmarshal(bytes, &deserialized)
	assert.NoError(t, err)

	assert.Equal(t, deserialized, m, "m == unmarshal(marshal(m))")
}
