package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers & Int
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "+"
	PLUS   = "-"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Function
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
