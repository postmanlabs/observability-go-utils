package slices

// Apply f to each element of slice in order, returning the results.
func Map[T1, T2 any](slice []T1, f func(T1) T2) []T2 {
	result, _ := MapWithErr[T1, T2](slice, func(t T1) (T2, error) {
		return f(t), nil
	})
	return result
}

// Apply f to each element of slice in parallel and wait for the results.
func MapInParallel[T1, T2 any](slice []T1, f func(T1) T2) []T2 {
	result, _ := MapWithErrInParallel[T1, T2](slice, func(t T1) (T2, error) {
		return f(t), nil
	})
	return result
}

// Apply f to each element of slice in order, returning the results.  Returns
// an error if f returns a non-nil error on any element.
func MapWithErr[T1, T2 any](slice []T1, f func(T1) (T2, error)) (rv []T2, err error) {
	if slice == nil {
		return nil, nil
	}

	rv = make([]T2, len(slice))
	for i, v := range slice {
		rv[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}

	return rv, nil
}

// Apply f to each element of slice in parallel and wait for all results.
// Returns the first error if f returns a non-nil error on any element.
func MapWithErrInParallel[T1, T2 any](slice []T1, f func(T1) (T2, error)) ([]T2, error) {
	// Return early if slice is nil to ensure we don't transform a nil slice
	// to an empty slice.
	if slice == nil {
		return nil, nil
	}

	outChan := make(chan runInParallelResult[T2], len(slice))
	errChan := make(chan error, len(slice))

	// Start jobs.
	for idx, elt := range slice {
		curIdx := idx
		curElt := elt
		go func() {
			res, err := f(curElt)
			if err != nil {
				errChan <- err
			} else {
				outChan <- runInParallelResult[T2]{idx: curIdx, result: res}
			}
		}()
	}

	// Wait for responses.
	numResponses := 0
	results := make([]T2, len(slice))
	for numResponses < len(slice) {
		select {
		case res := <-outChan:
			results[res.idx] = res.result
			numResponses++

		case err := <-errChan:
			return nil, err
		}
	}

	return results, nil
}

type runInParallelResult[T any] struct {
	// The slice index of the input that led to this result.
	idx    int
	result T
}
