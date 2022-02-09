package stack

import (
	"errors"
	"math"
)

var ErrNotEnoughtValuesOnTheStack error = errors.New("the operation needed more arguments than the ones available on the stack")
var ErrNotValid error = errors.New("the provided argument was not accepted by the operation")
var ErrDivideBy0 error = errors.New("it's not possible to divide by 0")

func boolToInt(condition bool) float64 {
	if condition {
		return 1
	}
	return 0
}

func isEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-15
}
