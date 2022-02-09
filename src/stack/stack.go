package stack

type Stack struct {
	values []float64
	length int
}

func NewStack(values []float64) *Stack {
	return &Stack{values, len(values)}
}
