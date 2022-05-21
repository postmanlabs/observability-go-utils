package maps

import (
	"encoding/json"

	"github.com/akitasoftware/go-utils/optionals"
	"github.com/pkg/errors"
)

// A map whose keys are complex enough that the map is represented in JSON as
// a list of key-value pairs.
type ComplexKeyMap[K comparable, V any] map[K]V

func (m ComplexKeyMap[K, V]) Put(k K, v V) {
	m[k] = v
}

func (m ComplexKeyMap[K, V]) Upsert(k K, v V, onConflict func(v, newV V) V) {
	newV := v
	if oldV, exists := m[k]; exists {
		newV = onConflict(oldV, newV)
	}
	m[k] = newV
}

func (m ComplexKeyMap[K, V]) Add(other ComplexKeyMap[K, V], onConflict func(v, newV V) V) {
	for k, v := range other {
		m.Upsert(k, v, onConflict)
	}
}

func (m ComplexKeyMap[K, V]) Get(k K) optionals.Optional[V] {
	v, exists := m[k]
	if exists {
		return optionals.Some(v)
	}
	return optionals.None[V]()
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied function is called, and the resulting
// value is entered into the map and returned.
func (m ComplexKeyMap[K, V]) GetOrCompute(k K, computeValue func() (V, error)) (V, error) {
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

// A version of GetOrCompute that is guaranteed to not error.
func (m ComplexKeyMap[K, V]) GetOrComputeNoError(k K, computeValue func() V) V {
	if v, exists := m[k]; exists {
		return v
	}

	v := computeValue()
	m[k] = v
	return v
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied default value is entered into the map
// and returned.
func (m ComplexKeyMap[K, V]) GetOrDefault(k K, defaultValue V) V {
	v, exists := m[k]
	if !exists {
		v = defaultValue
		m[k] = v
	}
	return v
}

func (m ComplexKeyMap[K, V]) ContainsKey(k K) bool {
	_, exists := m[k]
	return exists
}

func (m ComplexKeyMap[K, V]) Delete(k K) {
	delete(m, k)
}

func (m ComplexKeyMap[K, V]) IsEmpty() bool {
	return len(m) == 0
}

type SliceElt[K comparable, V any] struct {
	Key   K `json:"key"`
	Value V `json:"value"`
}

func (m ComplexKeyMap[K, V]) MarshalJSON() ([]byte, error) {
	slice := make([]SliceElt[K, V], 0, len(m))
	for k, v := range m {
		slice = append(slice, SliceElt[K, V]{
			Key:   k,
			Value: v,
		})
	}

	return json.Marshal(slice)
}

func (m *ComplexKeyMap[K, V]) UnmarshalJSON(text []byte) error {
	var slice []SliceElt[K, V]
	if err := json.Unmarshal(text, &slice); err != nil {
		return errors.Wrapf(err, "failed to unmarshal ComplexKeyMap")
	}

	*m = make(ComplexKeyMap[K, V], len(slice))
	for _, elt := range slice {
		(*m)[elt.Key] = elt.Value
	}
	return nil
}
