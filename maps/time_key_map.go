package maps

import (
	"time"

	"github.com/akitasoftware/go-utils/sets"
)

func getInternalMapKey(t time.Time) int64 {
	return t.Unix()
}

func getReverseKey(key int64) time.Time {
	return time.Unix(key, 0)
}

type TimeMap[V any] struct {
	internalMap map[int64]V
}

func NewTimeMap[V any]() TimeMap[V] {
	return TimeMap[V]{
		internalMap: make(map[int64]V),
	}
}

func (tm TimeMap[V]) Put(k time.Time, v V) {
	key := getInternalMapKey(k)
	tm.internalMap[key] = v
}

func (m TimeMap[V]) Upsert(k time.Time, v V, onConflict func(v, newV V) V) {
	newV := v
	key := getInternalMapKey(k)
	if oldV, exists := m.internalMap[key]; exists {
		newV = onConflict(oldV, newV)
	}
	m.internalMap[key] = newV
}

// If the key k is not already in the map, then it is entered into the map with
// the result of calling the supplied function. If the function returns an
// error, then the map is not modified, and the error is returned.
func (m TimeMap[V]) ComputeIfAbsent(k time.Time, computeValue func() (V, error)) error {
	_, err := m.GetOrCompute(k, computeValue)
	return err
}

// If the key k is not already in the map, then it is entered into the map with
// the result of calling the supplied function.
func (m TimeMap[V]) ComputeIfAbsentNoError(k time.Time, computeValue func() V) {
	m.GetOrComputeNoError(k, computeValue)
}

func (m TimeMap[V]) Add(other TimeMap[V], onConflict func(v, newV V) V) {
	for k, v := range other.internalMap {
		timeKey := getReverseKey(k)
		m.Upsert(timeKey, v, onConflict)
	}
}

func (tm TimeMap[V]) Get(k time.Time) (V, bool) {
	key := getInternalMapKey(k)
	v, exists := tm.internalMap[key]
	return v, exists
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied function is called, and the resulting
// value is entered into the map and returned.
func (m TimeMap[V]) GetOrCompute(k time.Time, computeValue func() (V, error)) (V, error) {
	key := getInternalMapKey(k)
	if v, exists := m.internalMap[key]; exists {
		return v, nil
	}

	v, err := computeValue()
	if err != nil {
		return v, err
	}

	m.internalMap[key] = v
	return v, nil
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied function is called, and the resulting
// value is entered into the map and returned.
func (m TimeMap[V]) GetOrComputeNoError(k time.Time, computeValue func() V) V {
	key := getInternalMapKey(k)
	if v, exists := m.internalMap[key]; exists {
		return v
	}

	v := computeValue()
	m.internalMap[key] = v
	return v
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the default Go value is returned.
func (m TimeMap[V]) GetOrDefault(k time.Time) V {
	key := getInternalMapKey(k)
	return m.internalMap[key]
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied value is entered into the map and
// returned.
func (m TimeMap[V]) GetOrValue(k time.Time, value V) V {
	key := getInternalMapKey(k)
	v, exists := m.internalMap[key]
	if !exists {
		v = value
		m.internalMap[key] = v
	}
	return v
}

func (m TimeMap[V]) ContainsKey(k time.Time) bool {
	key := getInternalMapKey(k)
	_, exists := m.internalMap[key]
	return exists
}

func (m TimeMap[V]) Delete(k time.Time) {
	key := getInternalMapKey(k)
	delete(m.internalMap, key)
}

func (m TimeMap[V]) IsEmpty() bool {
	return len(m.internalMap) == 0
}

func (m TimeMap[V]) Size() int {
	return len(m.internalMap)
}

func (m TimeMap[V]) Keys() []time.Time {
	keys := make([]time.Time, 0, len(m.internalMap))
	for k := range m.internalMap {
		timeKey := getReverseKey(k)
		keys = append(keys, timeKey)
	}
	return keys
}

func (m TimeMap[V]) KeySet() sets.Set[time.Time] {
	keys := sets.NewSet[time.Time]()
	for k := range m.internalMap {
		timeKey := getReverseKey(k)
		keys.Insert(timeKey)
	}
	return keys
}

func (m TimeMap[V]) Values() []V {
	values := make([]V, 0, len(m.internalMap))
	for _, v := range m.internalMap {
		values = append(values, v)
	}
	return values
}
