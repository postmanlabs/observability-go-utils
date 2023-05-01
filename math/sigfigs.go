package math

import (
	"math"
)

// Rounds the given number to n significant figures.
func RoundToSigFigs[T float32 | float64](x T, n int) T {
	if x == 0 {
		return 0
	}

	scale := math.Pow10(n - int(math.Ceil(math.Log10(math.Abs(float64(x))))))

	return T(math.Round(float64(x)*scale) / scale)
}

// Rounds the given number down to n significant figures.
func FloorToSigFigs[T float32 | float64](x T, n int) T {
	if x == 0 {
		return 0
	}

	scale := math.Pow10(n - int(math.Ceil(math.Log10(math.Abs(float64(x))))))

	return T(math.Floor(float64(x)*scale) / scale)
}

// Rounds the given number up to n significant figures.
func CeilToSigFigs[T float32 | float64](x T, n int) T {
	if x == 0 {
		return 0
	}

	scale := math.Pow10(n - int(math.Ceil(math.Log10(math.Abs(float64(x))))))

	return T(math.Ceil(float64(x)*scale) / scale)
}
