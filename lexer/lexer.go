package lexer

import "sprite/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var currentToken token.Token

	switch l.ch {
	case '=':
		currentToken = newToken(token.ASSIGN, l.ch)
	case ';':
		currentToken = newToken(token.SEMICOLON, l.ch)
	case '(':
		currentToken = newToken(token.LPAREN, l.ch)
	case ')':
		currentToken = newToken(token.RPAREN, l.ch)
	case ',':
		currentToken = newToken(token.COMMA, l.ch)
	case '+':
		currentToken = newToken(token.PLUS, l.ch)
	case '{':
		currentToken = newToken(token.LBRACE, l.ch)
	case '}':
		currentToken = newToken(token.RBRACE, l.ch)
	case 0:
		currentToken.Literal = ""
		currentToken.Type = token.EOF
	}

	l.readChar()
	return currentToken
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
