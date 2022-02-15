package parser

import (
	"github.com/Force4760/pipes/src/ast"
	"github.com/Force4760/pipes/src/tokens"
)

// Parse expression of the type
// if {then} {else}
// Where "then" and "else" are lists of expressions
func (p *Parser) parseIfExpression() (*ast.IfExpression, error) {
	stmt := &ast.IfExpression{Token: p.curToken}

	// THEN
	if !p.expectPeek(tokens.LBRACE) {
		return nil, ErrNoBraceAfterIf
	}
	p.nextToken()

	// Process THEN until the parser encounter a }
	for !p.curTokenIs(tokens.RBRACE) {
		parsed, err := p.Parse()

		if err != nil {
			return nil, err
		}

		stmt.Then = append(stmt.Then, parsed)
	}

	// ELSE
	if !p.expectPeek(tokens.LBRACE) {
		return nil, ErrNoElse
	}
	p.nextToken()

	// Process ELSE until the parser encounter a }
	for !p.curTokenIs(tokens.RBRACE) {
		parsed, err := p.Parse()

		if err != nil {
			return nil, err
		}

		stmt.Else = append(stmt.Else, parsed)
	}

	return stmt, nil
}
