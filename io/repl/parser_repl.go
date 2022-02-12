package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Force4760/pipes/src/ast"
	"github.com/Force4760/pipes/src/parser"
)

func Parser(in io.Reader) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)

		// Get the input
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		parsed, err := ParserReturn(line)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(parsed.ToJson())
		}

	}
}

func ParserReturn(line string) (*ast.FunctionFile, error) {
	lex, err := LexerReturn(line)

	if err != nil {
		return nil, err
	}

	parse := parser.New(lex)

	return parse.ParseProgram()
}
