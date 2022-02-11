package ast

// Convert a list of expressions to a string
func toStr(e []Expression) string {
	result := ""

	// Add every element to result
	for _, i := range e {
		result += i.Node() + " "
	}

	return result
}
