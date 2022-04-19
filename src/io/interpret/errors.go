package interpret

import "github.com/force4760/esostack/src/io/colorize"

// Error representing a Misplaced Token
type notExistsError struct {
	msg string
}

// Create a colorized Misplaced Error
func NotExistsError(msg string) error {
	return &notExistsError{
		colorize.Error(msg),
	}
}

// Implement the error interface
func (e *notExistsError) Error() string {
	return e.msg
}
