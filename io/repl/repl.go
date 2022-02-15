package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Force4760/pipes/src/stack"
)

//////////////////////////////////////////////////////////////
// NORMAL REPL COMMAND                                      //
//////////////////////////////////////////////////////////////

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
			err := replReturn(line, stk)

			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func replReturn(line string, stk *stack.Stack) error {
	parsed, err := parserReturn(line)

	if err != nil {
		return err
	}

	err = parsed.Eval(stk)

	return err
}
