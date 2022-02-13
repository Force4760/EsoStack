package tokens

// TokenType Enum
//go:generate stringer -type TokenType .
type TokenType int

const (
	INVALID TokenType = iota + 1
	EOF

	// Math Ops
	ADD
	SUB
	MUL
	DIV
	MOD
	POW

	// Logical Ops
	AND
	OR
	XOR
	NOT

	// Comparisson Ops
	LESS
	LESSEQ
	GREAT
	GREATEQ
	EQUAL
	DIFF

	// Math Funcs
	SIN
	COS
	TAN
	SQRT
	FACT
	EXP
	LOG
	CEIL
	FLOOR
	ABS
	MAX
	MIN
	SIM

	// Stack Funcs
	SWAP
	DUP
	DUP2
	DROP
	OVER
	ROT

	// KeyWords & DELIMITERS
	IF
	FOR
	LBRACE
	RBRACE

	// VALUES
	BOOL
	INT
	FLOAT
)
