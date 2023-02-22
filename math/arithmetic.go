package math

import (
	"github.com/akitasoftware/go-utils/constraints"
	go_constraints "golang.org/x/exp/constraints"
)

func Add[T constraints.Number](x, y T) T {
	return x + y
}

func Min[T constraints.Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Max[T constraints.Number](x, y T) T {
	if x > y {
		return x
	}
	return y
}

// Assumes positive inputs. Adapted from
// https://en.wikipedia.org/wiki/Euclidean_algorithm#Implementations.
func GCD[T go_constraints.Integer](a, b T) T {
	for b > 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Assumes positive inputs.
func LCM[T go_constraints.Integer](a, b T) T {
	return a * (b / GCD(a, b))
}
