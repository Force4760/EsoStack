package ast

import tok "github.com/Force4760/pipes/src/tokens"

// AST for any kind of value
// 1 -2 3.3 -4.4 true false + - * / % ^ && || ........
type ValExpression struct {
	Token tok.Token `json:"token"`
}

// Implementing the Expression interface
func (v *ValExpression) expressionNode() {}

// Implementing the Node interface
func (v *ValExpression) Node() string { return v.Token.Literal }
