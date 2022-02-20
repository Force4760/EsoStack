package stack

import (
	"fmt"

	"github.com/force4760/esostack/io/colorize"
)

//////////////////////////////////////////////////////////////
// STACK                                                    //
//////////////////////////////////////////////////////////////

var delim = colorize.Colorize(" | ", colorize.GREEN)
var format = "\n║ %g \n"

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

func (s *Stack) StringNoColor() string {
	// | 1.00 | 2.00 | ... |
	repr := "| "

	for _, i := range s.values {
		repr += fmt.Sprintf("%.2f%s", i, " | ")
	}

	// ( len )
	repr += fmt.Sprintf("  ( %d )", s.length)

	return repr
}

// Get a list containing all values in the stack
func (s *Stack) VerticalNoColorString() string {
	repr := ""

	for _, val := range s.values {
		repr += fmt.Sprintf(
			format,
			val,
		)
	}

	return repr
}

func (s *Stack) GetVals() []float64 {
	return s.values
}
