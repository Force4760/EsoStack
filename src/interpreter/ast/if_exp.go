package ast

import (
	"github.com/force4760/esostack/src/interpreter/stack"
	"github.com/force4760/esostack/src/interpreter/tokens"
)

//////////////////////////////////////////////////////////////
// If {then} {else}                                         //
//////////////////////////////////////////////////////////////

// AST for an If expression
// if { --- then --- } { --- else ---}
type IfExpression struct {
	Token tokens.Token `json:"token"`
	Then  []Expression `json:"then"`
	Else  []Expression `json:"else"`
}

// Implementing the Node interface
func (i *IfExpression) Literal() string {
	return "If{ " + toStr(i.Then) + "}{ " + toStr(i.Else) + "}"
}

// Implementing the Expression interface
func (i *IfExpression) Type() tokens.TokenType {
	return i.Token.Type
}

// Evaluate the If expression
func (i *IfExpression) Eval(s *stack.Stack) error {
	value, err := s.IsTrue()

	if err != nil {
		return err
	}

	if value {
		// TRUE on the stack
		// Evaluate each Node in the if-expression's then branch
		for _, t := range i.Then {
			err := t.Eval(s)

			if err != nil {
				return err
			}
		}

	} else {
		// FALSE on the stack
		// Evaluate each Node in the if-expression's else branch
		for _, e := range i.Else {
			err := e.Eval(s)

			if err != nil {
				return err
			}
		}
	}

	return nil

}
