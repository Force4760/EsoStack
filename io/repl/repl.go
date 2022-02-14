package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Force4760/pipes/src/stack"
)

func Repl(in io.Reader) {
	scanner := bufio.NewScanner(in)

	stk := stack.NewStack([]float64{})

	for {
		fmt.Print(PROMPT)

		// Get the input
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		if !Commands(line, stk) {
			err := ReplReturn(line, stk)

			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func ReplReturn(line string, stk *stack.Stack) error {
	parsed, err := ParserReturn(line)

	if err != nil {
		return err
	}

	err = parsed.Eval(stk)

	return err
}
