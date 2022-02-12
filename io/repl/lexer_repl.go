package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Force4760/pipes/src/lexer"
)

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

		lex, err := LexerReturn(line)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(lex.Tokens)
		}

	}
}

func LexerReturn(line string) (*lexer.Lexer, error) {
	// Create the lexer
	lex, err := lexer.New(line)

	if err != nil {
		return nil, err

	}

	// Tokenize
	err = lex.Tokenize()

	return lex, err
}
