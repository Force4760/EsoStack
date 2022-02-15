package stack

import (
	"fmt"

	"github.com/Force4760/pipes/io/colorize"
)

//////////////////////////////////////////////////////////////
// STACK                                                    //
//////////////////////////////////////////////////////////////

var delim = colorize.Colorize(" | ", colorize.GREEN)

// Stack data structure
type Stack struct {
	values []float64
	length int
}

// Create a new stack
func NewStack(values []float64) *Stack {
	return &Stack{values, len(values)}
}

// String Representation of the stack
//
// | 1.00 | 5.55 | .... | ( len )
func (s *Stack) String() string {
	// | 1.00 | 2.00 | ... |
	repr := delim

	for _, i := range s.values {
		repr += fmt.Sprintf("%.2f%s", i, delim)
	}

	// ( len )
	repr += colorize.Colorize(
		fmt.Sprintf("  ( %d )", s.length),
		colorize.BLUE,
	)

	return repr
}
