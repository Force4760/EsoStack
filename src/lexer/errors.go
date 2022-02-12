package lexer

import "github.com/Force4760/pipes/io/colorize"

// Error representing an Invalid Token
type invalidError struct {
	ch string
}

// Implement the error interface
func (e *invalidError) Error() string {
	return colorize.Colorize(
		"Error: Found an Invalid Character: "+e.ch,
		colorize.RED,
	)
}

// Create a new invalidError
func InvalidError(ch string) error {
	return &invalidError{ch: ch}
}

// Error representing an Invalid Number of Parentheses or Brackets
// () or [] do not match
type parenError struct {
	ch string
}

// Implement the error interface
func (e *parenError) Error() string {
	return colorize.Colorize(
		"Error: The number of "+e.ch+" does not match",
		colorize.RED,
	)
}

// Create a new parenError
func ParenError(ch string) error {
	return &parenError{ch: ch}
}
