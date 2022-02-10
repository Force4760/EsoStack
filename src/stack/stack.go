package stack

// Stack data structure
type Stack struct {
	values []float64
	length int
}

// Create a new stack
func NewStack(values []float64) *Stack {
	return &Stack{values, len(values)}
}
