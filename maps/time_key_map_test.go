package maps

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTimeMap(t *testing.T) {
	tm := NewTimeMap[int]()
	assert.NotNil(t, tm)
}

func TestPutAndGet(t *testing.T) {
	tm := NewTimeMap[int]()
	now := time.Now()
	tm.Put(now, 123)
	val, exists := tm.Get(now)
	assert.True(t, exists)
	assert.Equal(t, 123, val)
}

func TestUpsert(t *testing.T) {
	tm := NewTimeMap[int]()
	now := time.Now()
	tm.Put(now, 1)
	tm.Upsert(now, 2, func(oldVal, newVal int) int {
		return oldVal + newVal
	})
	val, exists := tm.Get(now)
	assert.True(t, exists)
	assert.Equal(t, 3, val)
}

func TestComputeIfAbsent(t *testing.T) {
	tm := NewTimeMap[int]()
	now := time.Now()
	err := tm.ComputeIfAbsent(now, func() (int, error) {
		return 42, nil
	})
	assert.NoError(t, err)
	val, exists := tm.Get(now)
	assert.True(t, exists)
	assert.Equal(t, 42, val)
}

func TestComputeIfAbsentNoError(t *testing.T) {
	tm := NewTimeMap[int]()
	now := time.Now()
	tm.ComputeIfAbsentNoError(now, func() int {
		return 42
	})
	val, exists := tm.Get(now)
	assert.True(t, exists)
	assert.Equal(t, 42, val)
}

func TestAdd(t *testing.T) {
	tm1 := NewTimeMap[int]()
	tm2 := NewTimeMap[int]()
	now := time.Now()
	tm1.Put(now, 1)
	tm2.Put(now.Add(time.Hour), 2)
	tm1.Add(tm2, func(v1, v2 int) int {
		return v1 + v2
	})
	assert.Equal(t, 2, tm1.Size())
}

func TestDelete(t *testing.T) {
	tm := NewTimeMap[int]()
	now := time.Now()
	tm.Put(now, 123)
	tm.Delete(now)
	_, exists := tm.Get(now)
	assert.False(t, exists)
}

func TestIsEmptyAndSize(t *testing.T) {
	tm := NewTimeMap[int]()
	assert.True(t, tm.IsEmpty())
	assert.Equal(t, 0, tm.Size())
	now := time.Now()
	tm.Put(now, 123)
	assert.False(t, tm.IsEmpty())
	assert.Equal(t, 1, tm.Size())
}

func TestKeys(t *testing.T) {
	tm := NewTimeMap[int]()
	now := time.Now()
	tm.Put(now, 123)
	keys := tm.Keys()
	assert.Equal(t, 1, len(keys))
}

func TestValues(t *testing.T) {
	tm := NewTimeMap[int]()
	now := time.Now()
	tm.Put(now, 123)
	values := tm.Values()
	assert.Contains(t, values, 123)
}

func TestContainsKey(t *testing.T) {
	tm := NewTimeMap[int]()
	now := time.Now()
	assert.False(t, tm.ContainsKey(now))
	tm.Put(now, 123)
	assert.True(t, tm.ContainsKey(now))
}
