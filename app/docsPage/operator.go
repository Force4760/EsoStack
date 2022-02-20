package docs

import "fmt"

//////////////////////////////////////////////////////////////
// OPERATORS                                                //
//////////////////////////////////////////////////////////////

// Operator Documentation struct
type operator struct {
	title       string
	chars       string
	args        int
	description string
}

// Create the one-line representation of an operator struct
//
// Title: chars
func (o *operator) OneLineString() string {
	return fmt.Sprintf(
		"%s: %s",
		o.title,
		o.chars,
	)
}
