package lexer

import "monkey-lang/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.currentChar {
	case '=':
		tok = newToken(token.ASSIGNMENT, l.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case '(':
		tok = newToken(token.OPENPARENTHESIS, l.currentChar)
	case ')':
		tok = newToken(token.CLOSEPARENTHESIS, l.currentChar)
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
	case '{':
		tok = newToken(token.OPENBRACE, l.currentChar)
	case '}':
		tok = newToken(token.CLOSEBRACE, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}
