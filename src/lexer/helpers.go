package lexer

import (
	"strconv"
	"strings"
)

const (
	PI    = "3.1415926536"
	NEPER = "2.7182818285"
	PHI   = "1.6180339887"
)

// Predicate checking if a string could represent an integer
// Can contain -0123456789
func isInt(ch string) bool {
	_, err := strconv.ParseInt(ch, 10, 8)

	return err == nil
}

// Predicate checking if a string could represent a float
// Can contain -.0123456789
func isFloat(ch string) bool {
	_, err := strconv.ParseFloat(ch, 32)

	return err == nil
}

// Predicate checking if the {} are balanced
func isValidBrace(input string) bool {
	numL := strings.Count(input, "{")
	numR := strings.Count(input, "}")

	return numL == numR
}
