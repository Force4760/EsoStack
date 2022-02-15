package ast

import (
	"strconv"

	"github.com/Force4760/pipes/src/stack"
	"github.com/Force4760/pipes/src/tokens"
)

// AST for any kind of value
type ValExpression struct {
	Token tokens.Token `json:"token"`
}

// Implementing the Expression interface
func (v *ValExpression) Type() tokens.TokenType {
	return v.Token.Type
}

// Implementing the Node interface
func (v *ValExpression) Literal() string {
	return v.Token.Literal
}

// Evaluate a "value" Node
func (v *ValExpression) Eval(s *stack.Stack) error {
	var err error = nil

	switch v.Type() {
	case tokens.ADD:
		err = s.Plus()
	case tokens.SUB:
		err = s.Minus()
	case tokens.MUL:
		err = s.Times()
	case tokens.DIV:
		err = s.Div()
	case tokens.MOD:
		err = s.Modulus()
	case tokens.POW:
		err = s.Power()

	case tokens.AND:
		err = s.And()
	case tokens.OR:
		err = s.Or()
	case tokens.XOR:
		err = s.Xor()
	case tokens.NOT:
		err = s.Not()

	case tokens.LESS:
		err = s.Lesser()
	case tokens.LESSEQ:
		err = s.LesserEq()
	case tokens.GREAT:
		err = s.Greater()
	case tokens.GREATEQ:
		err = s.GreaterEq()
	case tokens.EQUAL:
		err = s.Equal()
	case tokens.DIFF:
		err = s.Diferent()

	case tokens.SIN:
		err = s.Sin()
	case tokens.COS:
		err = s.Cos()
	case tokens.TAN:
		err = s.Tan()
	case tokens.SQRT:
		err = s.Sqrt()
	case tokens.FACT:
		err = s.Fact()
	case tokens.EXP:
		err = s.Exp()
	case tokens.LOG:
		err = s.Log()
	case tokens.CEIL:
		err = s.Ceil()
	case tokens.FLOOR:
		err = s.Floor()
	case tokens.ABS:
		err = s.Abs()
	case tokens.MAX:
		err = s.Max()
	case tokens.MIN:
		err = s.Min()
	case tokens.SIM:
		err = s.Simetric()

	case tokens.SWAP:
		err = s.Swap()
	case tokens.DUP:
		err = s.Dup()
	case tokens.DUP2:
		err = s.Dup2()
	case tokens.DROP:
		err = s.Drop()
	case tokens.OVER:
		err = s.Over()
	case tokens.ROT:
		err = s.Rot()

	case tokens.BOOL:
		// push boolean
		if v.Literal() == "true" {
			s.Push(1)
		} else {
			s.Push(0)
		}

	case tokens.INT:
		// push integer
		num, _ := strconv.ParseFloat(v.Literal(), 64)
		s.Push(num)

	case tokens.FLOAT:
		// push float
		num, _ := strconv.ParseFloat(v.Literal(), 64)
		s.Push(num)

	default:
		return ErrUnknown
	}

	return err
}
