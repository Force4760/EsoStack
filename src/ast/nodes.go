package ast

import (
	"github.com/force4760/stackit/src/stack"
	"github.com/force4760/stackit/src/tokens"
)

//////////////////////////////////////////////////////////////
// Expression                                               //
//////////////////////////////////////////////////////////////

type Expression interface {
	Literal() string
	Type() tokens.TokenType
	Eval(*stack.Stack) error
}
