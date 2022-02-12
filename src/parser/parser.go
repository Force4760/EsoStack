package parser

import (
	"github.com/Force4760/pipes/src/ast"
	"github.com/Force4760/pipes/src/lexer"
	"github.com/Force4760/pipes/src/tokens"
)

type Parser struct {
	tokens    []tokens.Token
	curToken  tokens.Token
	peekToken tokens.Token
	index     int
	length    int
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{
		tokens: lex.Tokens,
		length: len(lex.Tokens),
		index:  -1,
	}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() (*ast.FunctionFile, error) {
	program := ast.FunctionFile{}

	for p.curToken.Type != tokens.EOF {
		parsed, err := p.Parse()

		if err != nil {
			return nil, err
		}

		program = append(program, parsed)
	}

	return &program, nil
}

func (p *Parser) Parse() (ast.Expression, error) {
	var exp ast.Expression
	var err error = nil

	switch p.curToken.Type {
	case tokens.IF:
		exp, err = p.parseIfExpression()

		if err != nil {
			return nil, err
		}

	case tokens.FOR:
		exp, err = p.parseForExpression()

		if err != nil {
			return nil, err
		}

	case tokens.LBRACE:
		// Misplaced left brace {
		return nil, ExtraLBraceError

	case tokens.RBRACE:
		// Misplaced right brace }
		return nil, ExtraRBraceError

	default:
		exp = &ast.ValExpression{Token: p.curToken}
	}

	p.nextToken()

	return exp, err
}
