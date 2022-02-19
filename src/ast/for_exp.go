package ast

import (
	"github.com/force4760/stackit/src/stack"
	"github.com/force4760/stackit/src/tokens"
)

//////////////////////////////////////////////////////////////
// For {body}                                               //
//////////////////////////////////////////////////////////////

// AST for a FOR expression
// for { --- }
type ForExpression struct {
	Token tokens.Token `json:"token"`
	Body  []Expression `json:"body"`
}

// Implementing the Node interface
func (f *ForExpression) Literal() string {
	return "For{ " + toStr(f.Body) + "}"
}

// Implementing the Expression interface
func (f *ForExpression) Type() tokens.TokenType {
	return f.Token.Type
}

// Evaluate the FOR expression
func (f *ForExpression) Eval(s *stack.Stack) error {
	time, err := s.Pop()

	if err != nil {
		return err
	}

	// repeat time times
	for i := 0; i < int(time); i++ {
		// Evaluate each node on the for-loop's body
		for _, b := range f.Body {
			err := b.Eval(s)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
