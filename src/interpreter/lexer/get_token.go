package lexer

import tok "github.com/force4760/esostack/src/interpreter/tokens"

// Get the token at a specified index
// Return an invalidError if it finds an INVALID Token
func (le *Lexer) GetToken(index int) (tok.Token, error) {
	ch := le.chrs[index]

	var token tok.TokenType

	switch ch {
	case "+":
		token = tok.ADD
	case "-":
		token = tok.SUB
	case "*":
		token = tok.MUL
	case "/":
		token = tok.DIV
	case "%":
		token = tok.MOD
	case "^":
		token = tok.POW

	case ">":
		token = tok.GREAT
	case "<":
		token = tok.LESS
	case ">=":
		token = tok.GREATEQ
	case "<=":
		token = tok.LESSEQ
	case "==":
		token = tok.EQUAL
	case "!=":
		token = tok.DIFF

	case "&&", "and":
		token = tok.AND
	case "||", "or":
		token = tok.OR
	case "XX", "xor":
		token = tok.XOR
	case "!!", "not":
		token = tok.NOT

	case "√x", "sqrt":
		token = tok.SQRT
	case "x!", "fact":
		token = tok.FACT
	case "exp":
		token = tok.EXP
	case "log":
		token = tok.LOG
	case "sin":
		token = tok.SIN
	case "cos":
		token = tok.COS
	case "tan":
		token = tok.TAN
	case "↑", "max":
		token = tok.MAX
	case "↓", "min":
		token = tok.MIN
	case "-x", "sim":
		token = tok.SIM
	case "⌊x⌋", "floor":
		token = tok.FLOOR
	case "⌈x⌉", "ceil":
		token = tok.CEIL
	case "|x|", "abs":
		token = tok.ABS

	case "swap":
		token = tok.SWAP
	case "drop":
		token = tok.DROP
	case "dup":
		token = tok.DUP
	case "dup2":
		token = tok.DUP2
	case "duplicate":
		token = tok.DUPLICATE
	case "rot":
		token = tok.ROT
	case "over":
		token = tok.OVER
	case "size":
		token = tok.SIZE

	case "true", "false":
		token = tok.BOOL

	case "if":
		token = tok.IF
	case "for":
		token = tok.FOR

	case "{":
		token = tok.LBRACE
	case "}":
		token = tok.RBRACE

	case "e":
		return tok.NewToken(tok.FLOAT, NEPER), nil
	case "π", "pi":
		return tok.NewToken(tok.FLOAT, PI), nil
	case "τ", "2π", "tau":
		return tok.NewToken(tok.FLOAT, TAU), nil
	case "ϕ", "phi":
		return tok.NewToken(tok.FLOAT, PHI), nil

	default:
		if isInt(ch) {
			// INT
			token = tok.INT
		} else if isFloat(ch) {
			// FLOAT
			token = tok.FLOAT
		} else {
			// Invalid token, returns an error
			return tok.NewToken(tok.INVALID, ch), ErrInvalid(ch)
		}
	}

	return tok.NewToken(token, ch), nil
}
