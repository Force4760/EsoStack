package interpret

import (
	"bufio"
	"fmt"
	"io"

	"github.com/force4760/esostack/src/interpreter/stack"
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
			err := Interpret(line, stk)

			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
