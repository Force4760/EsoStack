package stack

import (
	"errors"
	"math"
)

var ErrNotEnoughtValuesOnTheStack error = errors.New("the operation needed more arguments than the ones available on the stack")
var ErrNotValid error = errors.New("the provided argument was not accepted by the operation")
var ErrDivideBy0 error = errors.New("it's not possible to divide by 0")

// Convert a boolean to a float
//
// true  -> 1
// false -> 0
func boolToInt(condition bool) float64 {
	if condition {
		return 1
	}
	return 0
}

// Predicate checking if two numbers are "Equal enough"
//
// Fixes the famous problem of (0.2 + 0.1 == 0.3) -> False
func isEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-15
}

// helper itterative factorial function
func factorial(num int) int {
	var result int = 1

	if num <= 1 {
		// num! = 1 if num <= 1

	} else {
		// num * (num-1) * (num-2) * (num-3) * ... * 3 * 2 * 1
		for i := 1; i <= num; i++ {
			result *= i
		}
	}

	return result
}
