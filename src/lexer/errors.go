package lexer

import "github.com/Force4760/pipes/io/colorize"

// Error representing an Invalid Token
type invalidError struct {
	ch string
}

func (e *invalidError) Error() string {
	return colorize.Colorize(
		"Error: Found and Invalid Identifier: "+e.ch,
		colorize.RED,
	)
}

func InvalidError(ch string) error {
	return &invalidError{ch: ch}
}

// Error representing an Invalid Number of Parentheses or Brackets
// () or [] do not match
type parenError struct {
	ch string
}

func (e *parenError) Error() string {
	return colorize.Colorize(
		"Error: The number of "+e.ch+" does not match",
		colorize.RED,
	)
}

func ParenError(ch string) error {
	return &parenError{ch: ch}
}
