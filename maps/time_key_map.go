package maps

import (
	"time"

	"github.com/akitasoftware/go-utils/optionals"
	"github.com/akitasoftware/go-utils/sets"
	"github.com/akitasoftware/go-utils/slices"
)

func getInternalMapKey(t time.Time) int64 {
	return t.Unix()
}

func getReverseKey(key int64) time.Time {
	return time.Unix(key, 0)
}

type TimeMap[V any] struct {
	internalMap Map[int64, V]
}

func NewTimeMap[V any]() TimeMap[V] {
	return TimeMap[V]{
		internalMap: NewMap[int64, V](),
	}
}

func (m TimeMap[V]) Put(k time.Time, v V) {
	key := getInternalMapKey(k)
	m.internalMap.Put(key, v)
}

func (m TimeMap[V]) Upsert(k time.Time, v V, onConflict func(v, newV V) V) {
	key := getInternalMapKey(k)
	m.internalMap.Upsert(key, v, onConflict)
}

// If the key k is not already in the map, then it is entered into the map with
// the value v.
func (m TimeMap[V]) PutIfAbsent(k time.Time, v V) {
	key := getInternalMapKey(k)
	m.internalMap.PutIfAbsent(key, v)
}

// If the key k is not already in the map, then it is entered into the map with
// the result of calling the supplied function. If the function returns an
// error, then the map is not modified, and the error is returned.
func (m TimeMap[V]) ComputeIfAbsent(k time.Time, computeValue func() (V, error)) error {
	key := getInternalMapKey(k)
	return m.internalMap.ComputeIfAbsent(key, computeValue)
}

// If the key k is not already in the map, then it is entered into the map with
// the result of calling the supplied function.
func (m TimeMap[V]) ComputeIfAbsentNoError(k time.Time, computeValue func() V) {
	key := getInternalMapKey(k)
	m.internalMap.ComputeIfAbsentNoError(key, computeValue)
}

func (m TimeMap[V]) Add(other TimeMap[V], onConflict func(v, newV V) V) {
	m.internalMap.Add(other.internalMap, onConflict)
}

func (m TimeMap[V]) Get(k time.Time) optionals.Optional[V] {
	key := getInternalMapKey(k)
	return m.internalMap.Get(key)
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied function is called, and the resulting
// value is entered into the map and returned.
func (m TimeMap[V]) GetOrCompute(k time.Time, computeValue func() (V, error)) (V, error) {
	key := getInternalMapKey(k)
	return m.internalMap.GetOrCompute(key, computeValue)
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied function is called, and the resulting
// value is entered into the map and returned.
func (m TimeMap[V]) GetOrComputeNoError(k time.Time, computeValue func() V) V {
	key := getInternalMapKey(k)
	return m.internalMap.GetOrComputeNoError(key, computeValue)
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the default Go value is returned.
func (m TimeMap[V]) GetOrDefault(k time.Time) V {
	key := getInternalMapKey(k)
	return m.internalMap.GetOrDefault(key)
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied value is entered into the map and
// returned.
func (m TimeMap[V]) GetOrValue(k time.Time, value V) V {
	key := getInternalMapKey(k)
	return m.internalMap.GetOrValue(key, value)
}

func (m TimeMap[V]) ContainsKey(k time.Time) bool {
	key := getInternalMapKey(k)
	return m.internalMap.ContainsKey(key)
}

func (m TimeMap[V]) Delete(k time.Time) {
	key := getInternalMapKey(k)
	m.internalMap.Delete(key)
}

func (m TimeMap[V]) IsEmpty() bool {
	return m.internalMap.IsEmpty()
}

func (m TimeMap[V]) Size() int {
	return m.internalMap.Size()
}

func (m TimeMap[V]) Keys() []time.Time {
	keys := m.internalMap.Keys()
	return slices.Map(keys, getReverseKey)
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
	return m.internalMap.Values()
}
