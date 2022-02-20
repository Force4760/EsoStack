package ast

import (
	"github.com/force4760/esostack/src/interpreter/stack"
	"github.com/force4760/esostack/src/interpreter/tokens"
)

//////////////////////////////////////////////////////////////
// Expression                                               //
//////////////////////////////////////////////////////////////

type Expression interface {
	Literal() string
	Type() tokens.TokenType
	Eval(*stack.Stack) error
}
