package goutils

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Abs returns the absolute value of x.
func Abs[T constraints.Integer | constraints.Float](x T) T {
	var zero T
	if x < zero {
		return -x
	}
	return x
}

// RoundUp returns the nearest large integer (ceil).
func RoundUp[T constraints.Integer](x float64) T {
	return T(math.Ceil(x))
}

// Round returns the nearest integer, rounding half away from zero.
func Round[T constraints.Integer](x float64) T {
	return T(math.Round(x))
}

// RoundDown returns the nearest small integer.
func RoundDown[T constraints.Integer](x float64) T {
	return T(x)
}

// Min returns the smallest value of numbers or zero.
func Min[T constraints.Integer | constraints.Float](numbers ...T) T {
	if len(numbers) == 0 {
		var zero T
		return zero
	}

	res := numbers[0]
	for _, num := range numbers {
		if num < res {
			res = num
		}
	}
	return res
}

// Max returns the largest value of numbers or zero.
func Max[T constraints.Integer | constraints.Float](numbers ...T) T {
	if len(numbers) == 0 {
		var zero T
		return zero
	}

	res := numbers[0]
	for _, num := range numbers {
		if num > res {
			res = num
		}
	}
	return res
}
