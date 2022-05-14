package lexer

import "sprite/token"

/* Returns a new token from provided token type and character byte */
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

/* Identify if the character-byte is equivalent to a letter */
func isLetter(ch byte) bool {
	return 'a' <= ch && 'z' >= ch || 'A' <= ch && 'Z' >= ch || ch == '_'
}

/* Identify if the current character-byte represents a number */
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
