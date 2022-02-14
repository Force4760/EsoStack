package ast

import (
	"github.com/Force4760/pipes/src/stack"
	"github.com/Force4760/pipes/src/tokens"
)

type Expression interface {
	Literal() string
	Type() tokens.TokenType
	Eval(*stack.Stack) error
}
