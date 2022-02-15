package tokens

//////////////////////////////////////////////////////////////
// TOKENS                                                   //
//////////////////////////////////////////////////////////////

// Token struct
// Represents a token in the src code
type Token struct {
	Type    TokenType
	Literal string
}

// Create a new token
func NewToken(tokenType TokenType, ch string) Token {
	return Token{Type: tokenType, Literal: ch}
}

// End Of File Token
var EOFToken = Token{EOF, ""}
