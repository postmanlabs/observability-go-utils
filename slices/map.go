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

// Applies f to each element of slice in order, removes any None results, and
// returns the rest.
func FilterMap[T1, T2 any](slice []T1, f func(T1) optionals.Optional[T2]) []T2 {
	result, _ := FilterMapWithErr(slice, func(t T1) (optionals.Optional[T2], error) {
		return f(t), nil
	})
	return result
}

// Applies f to each element of slice in order, removes any None results, and
// returns the rest. Returns an error if f returns a non-nil error on any
// element.
func FilterMapWithErr[T, U any](slice []T, f func(T) (optionals.Optional[U], error)) ([]U, error) {
	if slice == nil {
		return nil, nil
	}

	result := make([]U, len(slice))
	for _, t := range slice {
		u_opt, err := f(t)
		if err != nil {
			return nil, err
		}
		if u, exists := u_opt.Get(); exists {
			result = append(result, u)
		}
	}

	return result, nil
}
