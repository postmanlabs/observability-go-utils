package slices

import "github.com/akitasoftware/go-utils/optionals"

// Apply f to each element of slice in order, returning the results.
func Map[T1, T2 any](slice []T1, f func(T1) T2) []T2 {
	result, _ := MapWithErr(slice, func(t T1) (T2, error) {
		return f(t), nil
	})
	return result
}

// Apply f to each element of slice in order, returning the results.  Returns
// an error if f returns a non-nil error on any element.
func MapWithErr[T1, T2 any](slice []T1, f func(T1) (T2, error)) (rv []T2, err error) {
	return FilterMapWithErr(slice, func(t1 T1) (optionals.Optional[T2], error) {
		t2, err := f(t1)
		return optionals.Some(t2), err
	})
}

// Like Map, but f also takes in the element's index.
func MapIndex[T1, T2 any](slice []T1, f func(int, T1) T2) []T2 {
	result, _ := MapIndexWithErr(slice, func(idx int, t T1) (T2, error) {
		return f(idx, t), nil
	})
	return result
}

// Like MapWithErr, but f also takes in the element's index.
func MapIndexWithErr[T1, T2 any](slice []T1, f func(int, T1) (T2, error)) (rv []T2, err error) {
	return FilterMapIndexWithErr(slice, func(idx int, t1 T1) (optionals.Optional[T2], error) {
		t2, err := f(idx, t1)
		return optionals.Some(t2), err
	})
}
