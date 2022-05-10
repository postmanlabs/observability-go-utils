package sets

import (
	"encoding/json"
	"sort"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type OrderedSet[T constraints.Ordered] map[T]struct{}

func NewOrderedSet[T constraints.Ordered](vs ...T) OrderedSet[T] {
	s := make(OrderedSet[T], len(vs))
	for _, v := range vs {
		s.Insert(v)
	}
	return s
}

func (s OrderedSet[T]) Contains(v T) bool {
	return s.ContainsAny(v)
}

func (s OrderedSet[T]) ContainsAny(vs ...T) bool {
	for _, v := range vs {
		_, exists := s[v]
		if exists {
			return true
		}
	}
	return false
}

func (s OrderedSet[T]) ContainsAll(vs ...T) bool {
	for _, v := range vs {
		_, exists := s[v]
		if !exists {
			return false
		}
	}
	return true
}

func (s OrderedSet[T]) Insert(vs ...T) {
	for _, v := range vs {
		s[v] = struct{}{}
	}
}

func (s OrderedSet[T]) Delete(vs ...T) {
	for _, v := range vs {
		delete(s, v)
	}
}

func (s OrderedSet[T]) Union(other OrderedSet[T]) {
	for k := range other {
		s.Insert(k)
	}
}

func (s OrderedSet[T]) Intersect(other OrderedSet[T]) {
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

// Marshals as a sorted slice.
func (s OrderedSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.AsSlice())
}

func (s *OrderedSet[T]) UnmarshalJSON(text []byte) error {
	var slice []T
	if err := json.Unmarshal(text, &slice); err != nil {
		return errors.Wrapf(err, "failed to unmarshal stringset")
	}
	*s = make(OrderedSet[T], len(slice))
	for _, elt := range slice {
		(*s)[elt] = struct{}{}
	}
	return nil
}

func (s OrderedSet[T]) Clone() OrderedSet[T] {
	return maps.Clone(s)
}

// AsSlice returns the set as a sorted slice.
func (s OrderedSet[T]) AsSlice() []T {
	rv := make([]T, 0, len(s))
	for x, _ := range s {
		rv = append(rv, x)
	}
	slices.Sort(rv)
	return rv
}

// Creates a new set from the intersection of sets.
func IntersectOrdered[T constraints.Ordered](sets ...OrderedSet[T]) OrderedSet[T] {
	if len(sets) == 0 {
		return OrderedSet[T]{}
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
