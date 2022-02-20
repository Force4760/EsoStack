package parser

import (
	"github.com/force4760/esostack/src/interpreter/tokens"
)

// Predicate checking if the current Token is of type t
func (p *Parser) curTokenIs(t tokens.TokenType) bool {
	return p.curToken.Type == t
}

// Predicate checking if the next (peek) Token is of type t
func (p *Parser) peekTokenIs(t tokens.TokenType) bool {
	return p.peekToken.Type == t
}

// Predicate checking if the next (peek) Token is of type t
// In that case it advances to the nex token
func (p *Parser) expectPeek(t tokens.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

// Advance to the next token
// If there is no "next token" it will give an EOFToken (End Of File)
func (p *Parser) nextToken() {
	p.index += 1

	if p.index < p.length-1 {
		// Access 2 valid tokens
		p.curToken = p.tokens[p.index]
		p.peekToken = p.tokens[p.index+1]

	} else if p.index == p.length-1 {
		// Try to access the last Token + 1
		p.curToken = p.tokens[p.index]
		p.peekToken = tokens.EOFToken

	} else {
		// Try to access 2 Tokens after the end of the program
		p.curToken = tokens.EOFToken
		p.peekToken = tokens.EOFToken

	}
}
