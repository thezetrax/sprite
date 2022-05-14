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

	l.eatWhitespace()

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
	default:
		if isLetter(l.ch) {
			currentToken.Literal = l.readIdentifier()
			currentToken.Type = token.LookupIdent(currentToken.Literal)
			return currentToken
		} else if isDigit(l.ch) {
			currentToken.Type = token.INT
			currentToken.Literal = l.readNumber()
			return currentToken
		} else {
			currentToken = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return currentToken
}

/* Parses out the identifier */
func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

/* Parses out the number */
func (l *Lexer) readNumber() string {
	startPosition := l.position

	/* TODO: We can add things like '.' or something for floats */
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

/*
   Skip whitespace when encountered
*/
func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
