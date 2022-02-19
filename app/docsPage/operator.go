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

// Create the markdown representation of an operator struct
//
// ### [title]
// * char: [0.chars]
// * args: [0.args]
// * desc: [0.description]
func (o *operator) String() string {
	return fmt.Sprintf(
		`
### %s
* char: %s
* args: %d
* desc: %s
    `,
		o.title,
		o.chars,
		o.args,
		o.description,
	)
}

// Convert a list of operators to their markdown representation
func operatorList(ops []operator) string {
	result := ""

	for _, i := range ops {
		result += i.String()
	}

	return result
}
