package slices

import "github.com/akitasoftware/go-utils/optionals"

// Applies the predicate f to each element of slice in order, and returns those
// elements that satisfy f.
func Filter[T any](slice []T, f func(T) bool) []T {
	result, _ := FilterMapIndexWithErr(slice, func(_ int, t T) (optionals.Optional[T], error) {
		if f(t) {
			return optionals.Some(t), nil
		}
		return optionals.None[T](), nil
	})
	return result
}

// Applies the predicate f to each element of slice in order, and returns those
// elements that satisfy f. If f returns a non-nil error on any element,
// iteration immediately stops, and the error is returned.
func FilterWithErr[T any](slice []T, f func(T) (bool, error)) ([]T, error) {
	return FilterMapIndexWithErr(slice, func(_ int, t T) (optionals.Optional[T], error) {
		if include, err := f(t); !include || err != nil {
			return optionals.None[T](), err
		}
		return optionals.Some(t), nil
	})
}

// Like Filter, but f also takes in the element's index.
func FilterIndex[T any](slice []T, f func(int, T) bool) []T {
	result, _ := FilterMapIndexWithErr(slice, func(idx int, t T) (optionals.Optional[T], error) {
		if f(idx, t) {
			return optionals.Some(t), nil
		}
		return optionals.None[T](), nil
	})
	return result
}

// Like FilterWithErr, but f also takes in the element's index.
func FilterIndexWithErr[T any](slice []T, f func(int, T) (bool, error)) ([]T, error) {
	return FilterMapIndexWithErr(slice, func(idx int, t T) (optionals.Optional[T], error) {
		if include, err := f(idx, t); !include || err != nil {
			return optionals.None[T](), err
		}
		return optionals.Some(t), nil
	})
}
