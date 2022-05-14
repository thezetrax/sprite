package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

/* Tokens */
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers & Int
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

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

/* Keywords */
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

/* Checks weather the identifier is a keyword, or user defined */
func LookupIdent(ident string) TokenType {
	if keywordToken, ok := keywords[ident]; ok {
		return keywordToken
	}

	return IDENT
}
