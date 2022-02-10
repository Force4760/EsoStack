package ast

import tok "github.com/Force4760/pipes/src/tokens"

type ValExpression struct {
	Token tok.Token // the token of the function
}

// Implementing the Expression interface
func (v *ValExpression) expressionNode() {}

// Implementing the Node interface
func (v *ValExpression) TokenLiteral() string { return v.Token.Literal }
