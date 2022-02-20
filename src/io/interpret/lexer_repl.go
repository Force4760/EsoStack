package interpret

import (
	"bufio"
	"fmt"
	"io"
)

//////////////////////////////////////////////////////////////
// LEXER REPL COMMAND                                       //
//////////////////////////////////////////////////////////////

func Lexer(in io.Reader) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)

		// Get the input
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		lex, err := lexerReturn(line)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(lex.Tokens)
		}

	}
}
