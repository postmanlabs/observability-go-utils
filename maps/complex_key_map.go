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
	Map[K, V](m).Put(k, v)
}

func (m ComplexKeyMap[K, V]) Upsert(k K, v V, onConflict func(v, newV V) V) {
	Map[K, V](m).Upsert(k, v, onConflict)
}

func (m ComplexKeyMap[K, V]) Add(other ComplexKeyMap[K, V], onConflict func(v, newV V) V) {
	Map[K, V](m).Add(Map[K, V](other), onConflict)
}

func (m ComplexKeyMap[K, V]) Get(k K) optionals.Optional[V] {
	return Map[K, V](m).Get(k)
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied function is called, and the resulting
// value is entered into the map and returned.
func (m ComplexKeyMap[K, V]) GetOrCompute(k K, computeValue func() (V, error)) (V, error) {
	return Map[K, V](m).GetOrCompute(k, computeValue)
}

// A version of GetOrCompute that is guaranteed to not error.
func (m ComplexKeyMap[K, V]) GetOrComputeNoError(k K, computeValue func() V) V {
	return Map[K, V](m).GetOrComputeNoError(k, computeValue)
}

// Returns the value associated with the given key k. If the key does not
// already exist in the map, the supplied default value is entered into the map
// and returned.
func (m ComplexKeyMap[K, V]) GetOrDefault(k K, defaultValue V) V {
	return Map[K, V](m).GetOrDefault(k, defaultValue)
}

func (m ComplexKeyMap[K, V]) ContainsKey(k K) bool {
	return Map[K, V](m).ContainsKey(k)
}

func (m ComplexKeyMap[K, V]) Delete(k K) {
	Map[K, V](m).Delete(k)
}

func (m ComplexKeyMap[K, V]) IsEmpty() bool {
	return Map[K, V](m).IsEmpty()
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
