package maps

import "github.com/akitasoftware/go-utils/optionals"

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Upsert(k K, v V, onConflict func(v, newV V) V) {
	newV := v
	if oldV, exists := m[k]; exists {
		newV = onConflict(oldV, newV)
	}
	m[k] = newV
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
