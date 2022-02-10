package ast

import "github.com/Force4760/pipes/src/tokens"

type ForExpression struct {
	Token tokens.Token // the token.FOR token
	Body  []Expression
}

// Implementing the Expression interface
func (f *ForExpression) expressionNode() {}

// Implementing the Node interface
func (f *ForExpression) TokenLiteral() string { return f.Token.Literal }
