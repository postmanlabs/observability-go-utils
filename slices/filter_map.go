package slices

import "github.com/akitasoftware/go-utils/optionals"

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
	return FilterMapIndexWithErr(slice, func(_ int, t T) (optionals.Optional[U], error) {
		return f(t)
	})
}

// Like FilterMap, but f also takes in the element's index.
func FilterMapIndex[T1, T2 any](slice []T1, f func(int, T1) optionals.Optional[T2]) []T2 {
	result, _ := FilterMapIndexWithErr(slice, func(idx int, t T1) (optionals.Optional[T2], error) {
		return f(idx, t), nil
	})
	return result
}

// Like FilterMapWithErr, but f also takes in the element's index.
func FilterMapIndexWithErr[T, U any](slice []T, f func(int, T) (optionals.Optional[U], error)) ([]U, error) {
	if slice == nil {
		return nil, nil
	}

	result := make([]U, 0, len(slice))
	for idx, t := range slice {
		u_opt, err := f(idx, t)
		if err != nil {
			return nil, err
		}
		if u, exists := u_opt.Get(); exists {
			result = append(result, u)
		}
	}

	return result, nil
}
