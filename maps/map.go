package maps

import (
	"github.com/akitasoftware/go-utils/optionals"
	"github.com/akitasoftware/go-utils/sets"
)

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{}
}

func (m Map[K, V]) Put(k K, v V) {
	m[k] = v
}

func (m Map[K, V]) Upsert(k K, v V, onConflict func(v, newV V) V) {
	newV := v
	if oldV, exists := m[k]; exists {
		newV = onConflict(oldV, newV)
	}
	m[k] = newV
}

// If the key k is not already in the map, then it is entered into the map with
// the value v.
func (m Map[K, V]) PutIfAbsent(k K, v V) {
	m.GetOrValue(k, v)
}

// If the key k is not already in the map, then it is entered into the map with
// the result of calling the supplied function. If the function returns an
// error, then the map is not modified, and the error is returned.
func (m Map[K, V]) ComputeIfAbsent(k K, computeValue func() (V, error)) error {
	_, err := m.GetOrCompute(k, computeValue)
	return err
}

// If the key k is not already in the map, then it is entered into the map with
// the result of calling the supplied function.
func (m Map[K, V]) ComputeIfAbsentNoError(k K, computeValue func() V) {
	m.GetOrComputeNoError(k, computeValue)
}

func (m Map[K, V]) Add(other Map[K, V], onConflict func(v, newV V) V) {
	for k, v := range other {
		m.Upsert(k, v, onConflict)
	}
}

func (m Map[K, V]) Get(k K) optionals.Optional[V] {
	v, exists := m[k]
	if exists {
		return optionals.Some(v)
	}
	return optionals.None[V]()
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied function is called, and the resulting
// value is entered into the map and returned.
func (m Map[K, V]) GetOrCompute(k K, computeValue func() (V, error)) (V, error) {
	if v, exists := m[k]; exists {
		return v, nil
	}

	v, err := computeValue()
	if err != nil {
		return v, err
	}

	m[k] = v
	return v, nil
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied function is called, and the resulting
// value is entered into the map and returned.
func (m Map[K, V]) GetOrComputeNoError(k K, computeValue func() V) V {
	if v, exists := m[k]; exists {
		return v
	}

	v := computeValue()
	m[k] = v
	return v
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the default Go value is returned.
func (m Map[K, V]) GetOrDefault(k K) V {
	return m[k]
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied value is entered into the map and
// returned.
func (m Map[K, V]) GetOrValue(k K, value V) V {
	v, exists := m[k]
	if !exists {
		v = value
		m[k] = v
	}
	return v
}

func (m Map[K, V]) ContainsKey(k K) bool {
	_, exists := m[k]
	return exists
}

func (m Map[K, V]) Delete(k K) {
	delete(m, k)
}

func (m Map[K, V]) IsEmpty() bool {
	return len(m) == 0
}

func (m Map[K, V]) Size() int {
	return len(m)
}

func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[K, V]) KeySet() sets.Set[K] {
	keys := sets.NewSet[K]()
	for k := range m {
		keys.Insert(k)
	}
	return keys
}

func (m Map[K, V]) Values() []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
