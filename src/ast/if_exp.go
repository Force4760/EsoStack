package ast

import tok "github.com/Force4760/pipes/src/tokens"

type IfExpression struct {
	Token tok.Token // the token.IF token
	Then  []Expression
	Else  []Expression
}

// Implementing the Expression interface
func (i *IfExpression) expressionNode() {}

// Implementing the Node interface
func (i *IfExpression) TokenLiteral() string { return i.Token.Literal }