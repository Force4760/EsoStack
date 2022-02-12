package parser

import (
	"github.com/Force4760/pipes/io/colorize"
)

// Error representing a Misplaced Token
type misplacedError struct {
	msg string
}

// Create a colorized Misplaced Error
func MisplacedError(msg string) error {
	return &misplacedError{
		colorize.Colorize(msg, colorize.RED),
	}
}

// Implement the error interface
func (e *misplacedError) Error() string {
	return e.msg
}

// Standard Errors for the Parser Package

var NoBraceAfterIfError = MisplacedError("Expected { after if")
var NoBraceAfterForError = MisplacedError("Expected { after for")
var ExtraLBraceError = MisplacedError("Found a non-expected {")
var ExtraRBraceError = MisplacedError("Found a non-expected }")
var NoElseError = MisplacedError("Expected an else expression after if as \"if { then } { else }\"")
