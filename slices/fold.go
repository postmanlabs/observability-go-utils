package slices

func Fold[T1, T2 any](slice []T1, init T2, f func(T2, T1) T2) T2 {
	result, _ := FoldIndexWithErr(
		slice,
		init,
		func(accum T2, _ int, elt T1) (T2, error) {
			return f(accum, elt), nil
		},
	)
	return result
}

// Like Fold, but if f returns a non-nil element, iteration immediately stops,
// and the error is returned.
func FoldWithErr[T1, T2 any](slice []T1, init T2, f func(T2, T1) (T2, error)) (T2, error) {
	return FoldIndexWithErr(
		slice,
		init,
		func(accum T2, _ int, elt T1) (T2, error) {
			return f(accum, elt)
		},
	)
}

// Like Fold, but f also takes in the element's index.
func FoldIndex[T1, T2 any](slice []T1, init T2, f func(T2, int, T1) T2) T2 {
	result, _ := FoldIndexWithErr(
		slice,
		init,
		func(accum T2, idx int, elt T1) (T2, error) {
			return f(accum, idx, elt), nil
		},
	)
	return result
}

// Like FoldWithErr, but f also takes in the element's index.
func FoldIndexWithErr[T1, T2 any](slice []T1, init T2, f func(T2, int, T1) (T2, error)) (T2, error) {
	result := init
	for idx, elt := range slice {
		var err error
		result, err = f(result, idx, elt)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}
