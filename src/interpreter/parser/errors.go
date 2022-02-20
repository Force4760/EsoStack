package parser

import (
	"github.com/force4760/esostack/src/io/colorize"
)

// Error representing a Misplaced Token
type misplacedError struct {
	msg string
}

// Create a colorized Misplaced Error
func MisplacedError(msg string) error {
	return &misplacedError{
		colorize.Error(msg),
	}
}

// Implement the error interface
func (e *misplacedError) Error() string {
	return e.msg
}

// Standard Errors for the Parser Package

var ErrNoBraceAfterIf = MisplacedError("Expected { after if")
var ErrNoBraceAfterFor = MisplacedError("Expected { after for")
var ErrExtraLBrace = MisplacedError("Found a non-expected {")
var ErrExtraRBrace = MisplacedError("Found a non-expected }")
var ErrNoElse = MisplacedError("Expected an else expression after if as \"if { then } { else }\"")
