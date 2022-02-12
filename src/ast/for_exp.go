package ast

import "github.com/Force4760/pipes/src/tokens"

// AST for a FOR expression
// for { --- }
type ForExpression struct {
	Token tokens.Token `json:"token"`
	Body  []Expression `json:"body"`
}

// Implementing the Expression interface
func (f *ForExpression) expressionNode() {}

// Implementing the Node interface
func (f *ForExpression) Node() string {
	//      For{----------body--------}
	return "For{ " + toStr(f.Body) + "}"
}
