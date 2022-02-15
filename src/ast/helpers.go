package ast

import (
	"errors"

	"github.com/Force4760/pipes/io/colorize"
)

// Convert a list of expressions to a string
func toStr(e []Expression) string {
	result := ""

	// Add every element to result
	for _, i := range e {
		result += i.Literal() + " "
	}

	return result
}

var ErrUnknown error = errors.New(colorize.Colorize("Unknown AST Node, please report this bug", colorize.RED))
