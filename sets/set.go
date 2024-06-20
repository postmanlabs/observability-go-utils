package sets

import (
	"encoding/json"
	"sort"

	"github.com/akitasoftware/go-utils/optionals"
	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vs ...T) Set[T] {
	s := make(Set[T], len(vs))
	for _, v := range vs {
		s.Insert(v)
	}
	return s
}

// Create a set from a slice, applying a mapping function first to
// translate to a comparable type.
func FromSlice[T1 any, T2 comparable](vs []T1, f func(T1) T2) Set[T2] {
	s := make(Set[T2], len(vs))
	for _, v := range vs {
		s.Insert(f(v))
	}
	return s
}

// Applies f to each element of the slice in order, removes any None results,
// and converts the Some results to a Set.
func FromFilteredSlice[T1 any, T2 comparable](vs []T1, f func(T1) optionals.Optional[T2]) Set[T2] {
	s := make(Set[T2], len(vs))
	for _, v := range vs {
		if u, exists := f(v).Get(); exists {
			s.Insert(u)
		}
	}
	return s
}

func (s Set[T]) Equals(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for elt := range s {
		if _, exists := other[elt]; !exists {
			return false
		}
	}
	return true
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Size() int {
	return len(s)
}

// Converts v to an optional value, depending on whether it is a member of s.
// Returns Some(v) if s contains v. Returns None otherwise.
func (s Set[T]) Get(v T) optionals.Optional[T] {
	if s.Contains(v) {
		return optionals.Some(v)
	}
	return optionals.None[T]()
}

func (s Set[T]) Contains(v T) bool {
	return s.ContainsAny(v)
}

func (s Set[T]) ContainsAny(vs ...T) bool {
	for _, v := range vs {
		_, exists := s[v]
		if exists {
			return true
		}
	}
	return false
}

func (s Set[T]) ContainsAll(vs ...T) bool {
	for _, v := range vs {
		_, exists := s[v]
		if !exists {
			return false
		}
	}
	return true
}

func (s Set[T]) Insert(vs ...T) {
	for _, v := range vs {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Delete(vs ...T) {
	for _, v := range vs {
		delete(s, v)
	}
}

func (s Set[T]) Clear() {
	for v := range s {
		delete(s, v)
	}
}

func (s Set[T]) Union(other Set[T]) {
	for k := range other {
		s.Insert(k)
	}
}

func (s Set[T]) Intersect(other Set[T]) {
	var toDelete []T
	for k := range s {
		if _, exists := other[k]; !exists {
			toDelete = append(toDelete, k)
		}
	}
	for _, k := range toDelete {
		delete(s, k)
	}
}

func (s Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.AsSlice())
}

func (s *Set[T]) UnmarshalJSON(text []byte) error {
	var slice []T
	if err := json.Unmarshal(text, &slice); err != nil {
		return errors.Wrapf(err, "failed to unmarshal stringset")
	}
	*s = make(Set[T], len(slice))
	for _, elt := range slice {
		(*s)[elt] = struct{}{}
	}
	return nil
}

func (s Set[T]) Clone() Set[T] {
	clonedMap := maps.Clone(s)
	if clonedMap == nil {
		clonedMap = NewSet[T]()
	}
	return clonedMap
}

// AsSlice returns the set as a slice in a nondeterministic order.
func (s Set[T]) AsSlice() []T {
	rv := make([]T, 0, len(s))
	for x := range s {
		rv = append(rv, x)
	}
	return rv
}

// Returns the set as an OrderedSet. Changes to the returned OrderedSet will be
// reflected in this set.
func AsOrderedSet[T constraints.Ordered](s Set[T]) OrderedSet[T] {
	return OrderedSet[T](s)
}

// Creates a new set from the intersection of sets.
func Intersect[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return Set[T]{}
	}

	// Sort by set length.  Starting with the smallest set reduces
	// the work we need to do.
	sort.Slice(sets, func(i, j int) bool {
		return len(sets[i]) < len(sets[j])
	})

	base := sets[0].Clone()
	for _, next := range sets[1:] {
		base.Intersect(next)
	}

	return base
}

// Applies the given function to each element of a set. Returns the resulting
// set of function outputs.
func Map[T, U comparable](ts Set[T], f func(T) U) Set[U] {
	return FilterMap(ts, func(t T) optionals.Optional[U] {
		return optionals.Some(f(t))
	})
}

// Applies the given function to each element of a set. Removes any None results
// and returns the rest.
func FilterMap[T, U comparable](ts Set[T], f func(T) optionals.Optional[U]) Set[U] {
	result := NewSet[U]()
	for t := range ts {
		if u, exists := f(t).Get(); exists {
			result.Insert(u)
		}
	}
	return result
}
