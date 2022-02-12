package ast

import tok "github.com/Force4760/pipes/src/tokens"

// AST for an If expression
// if { --- then --- } { --- else ---}
type IfExpression struct {
	Token tok.Token    `json:"token"`
	Then  []Expression `json:"then"`
	Else  []Expression `json:"else"`
}

// Implementing the Expression interface
func (i *IfExpression) expressionNode() {}

// Implementing the Node interface
func (i *IfExpression) Node() string {
	//      If{---------then---------}{---------else---------}
	return "If{ " + toStr(i.Then) + "}{ " + toStr(i.Else) + "}"
}
