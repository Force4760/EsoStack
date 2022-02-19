package ast

import (
	"errors"

	"github.com/force4760/esostack/io/colorize"
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

var ErrUnknown error = errors.New(
	colorize.Error("Unknown AST Node, please report this bug"),
)
