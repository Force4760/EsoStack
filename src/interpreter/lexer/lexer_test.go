package lexer

import (
	"testing"

	"github.com/force4760/esostack/src/interpreter/tokens"
	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	t.Run("Data Types", func(t *testing.T) {
		input := "2 -3 2.3 -2.3 true false e pi phi π ϕ"
		got, _ := New(input)

		got.Tokenize()

		want := &Lexer{
			input:  input,
			length: len(input),
			chrs:   []string{"2", "-3", "2.3", "-2.3", "true", "false", "e", "pi", "phi", "π", "ϕ"},
			Tokens: []tokens.Token{
				{Type: tokens.INT, Literal: "2"},
				{Type: tokens.INT, Literal: "-3"},
				{Type: tokens.FLOAT, Literal: "2.3"},
				{Type: tokens.FLOAT, Literal: "-2.3"},
				{Type: tokens.BOOL, Literal: "true"},
				{Type: tokens.BOOL, Literal: "false"},
				{Type: tokens.FLOAT, Literal: NEPER},
				{Type: tokens.FLOAT, Literal: PI},
				{Type: tokens.FLOAT, Literal: PHI},
				{Type: tokens.FLOAT, Literal: PI},
				{Type: tokens.FLOAT, Literal: PHI},
			},
		}

		assert.Equal(t, got, want)
	})
	t.Run("Operators", func(t *testing.T) {
		input := "+ - * / % ^ > < >= <= == != && || XX !! and or xor not"
		got, _ := New(input)

		got.Tokenize()

		want := &Lexer{
			input:  input,
			length: len(input),
			chrs:   []string{"+", "-", "*", "/", "%", "^", ">", "<", ">=", "<=", "==", "!=", "&&", "||", "XX", "!!", "and", "or", "xor", "not"},
			Tokens: []tokens.Token{
				{Type: tokens.ADD, Literal: "+"},
				{Type: tokens.SUB, Literal: "-"},
				{Type: tokens.MUL, Literal: "*"},
				{Type: tokens.DIV, Literal: "/"},
				{Type: tokens.MOD, Literal: "%"},
				{Type: tokens.POW, Literal: "^"},
				{Type: tokens.GREAT, Literal: ">"},
				{Type: tokens.LESS, Literal: "<"},
				{Type: tokens.GREATEQ, Literal: ">="},
				{Type: tokens.LESSEQ, Literal: "<="},
				{Type: tokens.EQUAL, Literal: "=="},
				{Type: tokens.DIFF, Literal: "!="},
				{Type: tokens.AND, Literal: "&&"},
				{Type: tokens.OR, Literal: "||"},
				{Type: tokens.XOR, Literal: "XX"},
				{Type: tokens.NOT, Literal: "!!"},
				{Type: tokens.AND, Literal: "and"},
				{Type: tokens.OR, Literal: "or"},
				{Type: tokens.XOR, Literal: "xor"},
				{Type: tokens.NOT, Literal: "not"},
			},
		}

		assert.Equal(t, got, want)
	})
	t.Run("Math Functions", func(t *testing.T) {
		input := "√x sqrt x! fact exp log sin cos tan ↑ max ↓ min ± sim ⌊x⌋ floor ⌈x⌉ ceil |x| abs"
		got, _ := New(input)

		got.Tokenize()

		want := &Lexer{
			input:  input,
			length: len(input),
			chrs:   []string{"√x", "sqrt", "x!", "fact", "exp", "log", "sin", "cos", "tan", "↑", "max", "↓", "min", "±", "sim", "⌊x⌋", "floor", "⌈x⌉", "ceil", "|x|", "abs"},
			Tokens: []tokens.Token{
				{Type: tokens.SQRT, Literal: "√x"},
				{Type: tokens.SQRT, Literal: "sqrt"},
				{Type: tokens.FACT, Literal: "x!"},
				{Type: tokens.FACT, Literal: "fact"},
				{Type: tokens.EXP, Literal: "exp"},
				{Type: tokens.LOG, Literal: "log"},
				{Type: tokens.SIN, Literal: "sin"},
				{Type: tokens.COS, Literal: "cos"},
				{Type: tokens.TAN, Literal: "tan"},
				{Type: tokens.MAX, Literal: "↑"},
				{Type: tokens.MAX, Literal: "max"},
				{Type: tokens.MIN, Literal: "↓"},
				{Type: tokens.MIN, Literal: "min"},
				{Type: tokens.SIM, Literal: "±"},
				{Type: tokens.SIM, Literal: "sim"},
				{Type: tokens.FLOOR, Literal: "⌊x⌋"},
				{Type: tokens.FLOOR, Literal: "floor"},
				{Type: tokens.CEIL, Literal: "⌈x⌉"},
				{Type: tokens.CEIL, Literal: "ceil"},
				{Type: tokens.ABS, Literal: "|x|"},
				{Type: tokens.ABS, Literal: "abs"},
			},
		}

		assert.Equal(t, got, want)
	})
	t.Run("Stack Functions", func(t *testing.T) {
		input := "swap drop dup dup2 rot over if for {} size"
		got, _ := New(input)

		got.Tokenize()

		want := &Lexer{
			input:  input,
			length: len(input),
			chrs:   []string{"swap", "drop", "dup", "dup2", "rot", "over", "if", "for", "{", "}", "size"},
			Tokens: []tokens.Token{
				{Type: tokens.SWAP, Literal: "swap"},
				{Type: tokens.DROP, Literal: "drop"},
				{Type: tokens.DUP, Literal: "dup"},
				{Type: tokens.DUP2, Literal: "dup2"},
				{Type: tokens.ROT, Literal: "rot"},
				{Type: tokens.OVER, Literal: "over"},
				{Type: tokens.IF, Literal: "if"},
				{Type: tokens.FOR, Literal: "for"},
				{Type: tokens.LBRACE, Literal: "{"},
				{Type: tokens.RBRACE, Literal: "}"},
				{Type: tokens.SIZE, Literal: "size"},
			},
		}

		assert.Equal(t, got, want)
	})
}

func TestLexerBraces(t *testing.T) {
	t.Run("Braces Error", func(t *testing.T) {
		input := "{ test"
		_, got := New(input)

		want := ErrParen("{ }")

		assert.Equal(t, got, want)
	})
	t.Run("Braces Error 2", func(t *testing.T) {
		input := "test }"
		_, got := New(input)

		want := ErrParen("{ }")

		assert.Equal(t, got, want)
	})
	t.Run("Braces Error 3", func(t *testing.T) {
		input := " { test }"
		_, got := New(input)

		var want error = nil

		assert.Equal(t, got, want)
	})
}

func TestLexerParens(t *testing.T) {
	t.Run("Parens Error 1", func(t *testing.T) {
		input := "( fa a (fas) as ag )"
		_, got := New(input)

		var want error = nil

		assert.Equal(t, got, want)
	})
	t.Run("Parens Error 2", func(t *testing.T) {
		input := "( fas) as ag )"
		_, got := New(input)

		want := ErrParen("( )")

		assert.Equal(t, got, want)
	})
	t.Run("Parens Error 3", func(t *testing.T) {
		input := "( (fas) (as ag )"
		_, got := New(input)

		want := ErrParen("( )")

		assert.Equal(t, got, want)
	})
}
func TestLexerInvalid(t *testing.T) {
	t.Run("Invalid Error 1", func(t *testing.T) {
		input := "{ $ }"
		lex, _ := New(input)

		got := lex.Tokenize()
		want := ErrInvalid("$")

		assert.Equal(t, got, want)
	})
	t.Run("Invalid Error 2", func(t *testing.T) {
		input := "{ dgff dsg dsg test }"
		lex, _ := New(input)

		got := lex.Tokenize()
		want := ErrInvalid("dgff")

		assert.Equal(t, got, want)
	})
}

func TestIsInt(t *testing.T) {
	t.Run("Is Int positive", func(t *testing.T) {
		got := isInt("123456")
		want := true

		assert.Equal(t, got, want)
	})
	t.Run("Is Int negative", func(t *testing.T) {
		got := isInt("-1233456")
		want := true

		assert.Equal(t, got, want)
	})
	t.Run("Is Int float -> False", func(t *testing.T) {
		got := isInt("-12.3456")
		want := false

		assert.Equal(t, got, want)
	})
	t.Run("Is Int nonInt -> False", func(t *testing.T) {
		got := isInt("saffasfa")
		want := false

		assert.Equal(t, got, want)
	})
}

func TestIsFloat(t *testing.T) {
	t.Run("Is Float positive", func(t *testing.T) {
		got := isFloat("12.3456")
		want := true

		assert.Equal(t, got, want)
	})
	t.Run("Is Float negative", func(t *testing.T) {
		got := isFloat("-12.33456")
		want := true

		assert.Equal(t, got, want)
	})
	t.Run("Is Float int -> True", func(t *testing.T) {
		got := isFloat("-123456")
		want := true

		assert.Equal(t, got, want)
	})
	t.Run("Is Float nonInt -> False", func(t *testing.T) {
		got := isFloat("saffasfa")
		want := false

		assert.Equal(t, got, want)
	})
}

func TestIsValid(t *testing.T) {
	t.Run("valid {}", func(t *testing.T) {
		got := isValidBrace("abc {} asfasf asf}}gg { f {")
		want := true

		assert.Equal(t, got, want)
	})
	t.Run("not valid {}", func(t *testing.T) {
		got := isValidBrace("abc { asfasf asf}}gg { f {")
		want := false

		assert.Equal(t, got, want)
	})
	t.Run("valid ()", func(t *testing.T) {
		got := isValidParen("abc () asfasf asf))gg ( f (")
		want := true

		assert.Equal(t, got, want)
	})
	t.Run("not valid ()", func(t *testing.T) {
		got := isValidParen("abc ( asfasf asf))gg ( f (")
		want := false

		assert.Equal(t, got, want)
	})
}
