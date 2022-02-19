package lexer

import (
	tok "github.com/force4760/esostack/src/tokens"
)

//////////////////////////////////////////////////////////////
// Lexer                                                    //
//////////////////////////////////////////////////////////////

// Lexer type
// Responsible for Tokenizing the provided src code
type Lexer struct {
	input  string
	length int
	chrs   []string
	Tokens []tok.Token
}

// Create a new Lexer with the provided input
// chrs and tokens are inicialized to their 0 values
func New(input string) (*Lexer, error) {
	if !isValidBrace(input) {
		return nil, ErrParen("{ }")
	} else if !isValidParen(input) {
		return nil, ErrParen("( )")
	}

	le := &Lexer{
		input:  input,
		length: len(input),
	}

	return le, nil
}

// Turn src code into Tokens
// Return an error if it finds an INVALID Token
func (le *Lexer) Tokenize() error {
	// Split all the words in the src code
	le.Split()

	for i := range le.chrs {
		token, err := le.GetToken(i)

		if err != nil {
			return err
		}

		le.Tokens = append(le.Tokens, token)
	}

	return nil
}
