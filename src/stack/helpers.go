package stack

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var ErrNotEnoughtValuesOnTheStack error = errors.New("the operation needed more arguments than the ones available on the stack")
var ErrNotValid error = errors.New("the provided argument was not accepted by the operation")
var ErrDivideBy0 error = errors.New("it's not possible to divide by 0")

func boolToInt(condition bool) float64 {
	if condition {
		return 1
	}
	return 0
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}
