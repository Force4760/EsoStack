package parser

import (
	"testing"

	"github.com/Force4760/pipes/src/lexer"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	input := "5 4 5 6 +"

	lex, _ := lexer.New(input)
	_ = lex.Tokenize()

	want := &Parser{
		tokens:    lex.Tokens,
		curToken:  lex.Tokens[0],
		peekToken: lex.Tokens[1],
		index:     0,
		length:    len(lex.Tokens),
	}

	got := New(lex)

	assert.Equal(t, want, got, "The two words should be the same.")
}

func TestParserIf(t *testing.T) {
	t.Run("If test 1", func(t *testing.T) {
		input := "if{2 + 1}{2 4 5}"

		lex, _ := lexer.New(input)
		lex.Tokenize()
		pars := New(lex)

		got, _ := pars.ParseProgram()

		wanted := "If{ 2 + 1 }{ 2 4 5 } "

		assert.Equal(t, wanted, got.String(), "Wrong string representation")

	})
	t.Run("If test 2 (Chained Ifs)", func(t *testing.T) {
		input := "if{2 if{4 5 6}{2 3}+ 1}{2 4 5}"

		lex, _ := lexer.New(input)
		lex.Tokenize()
		pars := New(lex)

		got, _ := pars.ParseProgram()

		wanted := "If{ 2 If{ 4 5 6 }{ 2 3 } + 1 }{ 2 4 5 } "

		assert.Equal(t, wanted, got.String(), "Wrong string representation")

	})
	t.Run("No Brace After If test", func(t *testing.T) {
		input := "if 2 {2 if{4 5 6}{2 3}+ 1}{2 4 5}"

		lex, _ := lexer.New(input)
		lex.Tokenize()
		pars := New(lex)

		_, got := pars.ParseProgram()

		wanted := NoBraceAfterIfError

		assert.Equal(t, wanted, got, "Expecting a NoBraceAfterIfError error")
	})
	t.Run("No Else test", func(t *testing.T) {
		input := "if {2 if{4 5 6}{2 3}+ 1} 2"

		lex, _ := lexer.New(input)
		lex.Tokenize()
		pars := New(lex)

		_, got := pars.ParseProgram()

		wanted := NoElseError

		assert.Equal(t, wanted, got, "Expecting a NoElseError error")
	})
}

func TestParserFor(t *testing.T) {
	t.Run("For test 1", func(t *testing.T) {
		input := "for{2 + 1}"

		lex, _ := lexer.New(input)
		lex.Tokenize()
		pars := New(lex)

		got, _ := pars.ParseProgram()

		wanted := "For{ 2 + 1 } "

		assert.Equal(t, wanted, got.String(), "Wrong string representation")

	})
	t.Run("For test 2 (Chained For)", func(t *testing.T) {
		input := "for{2 + for{2 -56 4 sqrt}1}"

		lex, _ := lexer.New(input)
		lex.Tokenize()
		pars := New(lex)

		got, _ := pars.ParseProgram()

		wanted := "For{ 2 + For{ 2 -56 4 sqrt } 1 } "

		assert.Equal(t, wanted, got.String(), "Wrong string representation")

	})
	t.Run("No Brace After For test", func(t *testing.T) {
		input := "for 2 {2 if{4 5 6}{2 3}+ 1}{2 4 5}"

		lex, _ := lexer.New(input)
		lex.Tokenize()
		pars := New(lex)

		_, got := pars.ParseProgram()

		wanted := NoBraceAfterForError

		assert.Equal(t, wanted, got, "Expecting a NoBraceAfterForError error")
	})
}
