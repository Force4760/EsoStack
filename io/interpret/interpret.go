package interpret

import (
	"github.com/force4760/stackit/src/ast"
	"github.com/force4760/stackit/src/lexer"
	"github.com/force4760/stackit/src/parser"
	"github.com/force4760/stackit/src/stack"
)

func lexerReturn(line string) (*lexer.Lexer, error) {
	// Create the lexer
	lex, err := lexer.New(line)

	if err != nil {
		return nil, err

	}

	// Tokenize
	err = lex.Tokenize()

	return lex, err
}

func parserReturn(line string) (*ast.FunctionFile, error) {
	lex, err := lexerReturn(line)

	if err != nil {
		return nil, err
	}

	parse := parser.New(lex)

	return parse.ParseProgram()
}

func Interpret(line string, stk *stack.Stack) error {
	parsed, err := parserReturn(line)

	if err != nil {
		return err
	}

	err = parsed.Eval(stk)

	return err
}
