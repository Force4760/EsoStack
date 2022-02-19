package parser

import (
	"github.com/force4760/esostack/src/ast"
	"github.com/force4760/esostack/src/lexer"
	"github.com/force4760/esostack/src/tokens"
)

//////////////////////////////////////////////////////////////
// PARSER                                                   //
//////////////////////////////////////////////////////////////

// Parser struct
type Parser struct {
	tokens    []tokens.Token
	curToken  tokens.Token
	peekToken tokens.Token
	index     int
	length    int
}

// Create a new Parser object from a lexer
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

// Parse a program into an AST
func (p *Parser) ParseProgram() (*ast.FunctionFile, error) {
	program := ast.FunctionFile{}

	// while the token is not EOF
	for p.curToken.Type != tokens.EOF {
		parsed, err := p.Parse()

		if err != nil {
			return nil, err
		}

		program = append(program, parsed)
	}

	return &program, nil
}

// Parse a single Token into an AST expression
func (p *Parser) Parse() (ast.Expression, error) {
	var exp ast.Expression
	var err error = nil

	switch p.curToken.Type {
	case tokens.IF:
		exp, err = p.parseIfExpression()

		// Found an error
		if err != nil {
			return nil, err
		}

	case tokens.FOR:
		exp, err = p.parseForExpression()

		// Found an error
		if err != nil {
			return nil, err
		}

	case tokens.LBRACE:
		// Misplaced left brace { Error
		return nil, ErrExtraLBrace

	case tokens.RBRACE:
		// Misplaced right brace } Error
		return nil, ErrExtraRBrace

	default:
		exp = &ast.ValExpression{Token: p.curToken}
	}

	p.nextToken()

	return exp, err
}
