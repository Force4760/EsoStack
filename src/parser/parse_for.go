package parser

import (
	"github.com/Force4760/pipes/src/ast"
	"github.com/Force4760/pipes/src/tokens"
)

// Parse expression of the type
// for {body}
// Where body is a list of expressions
func (p *Parser) parseForExpression() (*ast.ForExpression, error) {
	stmt := &ast.ForExpression{Token: p.curToken}

	// BODY
	if !p.expectPeek(tokens.LBRACE) {
		return nil, ErrNoBraceAfterFor
	}
	p.nextToken()

	// Process BODY until the parser encounter a }
	for !p.curTokenIs(tokens.RBRACE) {
		parsed, err := p.Parse()

		if err != nil {
			return nil, err
		}

		stmt.Body = append(stmt.Body, parsed)
	}

	return stmt, nil
}
