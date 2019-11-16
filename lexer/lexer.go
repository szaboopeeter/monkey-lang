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

	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			currentChar := l.currentChar
			l.readChar()
			literal := string(currentChar) + string(l.currentChar)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			tok = newToken(token.ASSIGNMENT, l.currentChar)
		}
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
	case '-':
		tok = newToken(token.MINUS, l.currentChar)
	case '!':
		if l.peekChar() == '=' {
			currentChar := l.currentChar
			l.readChar()
			literal := string(currentChar) + string(l.currentChar)
			tok = token.Token{Type: token.NOTEQUAL, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.currentChar)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.currentChar)
	case '/':
		tok = newToken(token.SLASH, l.currentChar)
	case '<':
		tok = newToken(token.LESSTHAN, l.currentChar)
	case '>':
		tok = newToken(token.GREATERTHAN, l.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case '(':
		tok = newToken(token.OPENPARENTHESIS, l.currentChar)
	case ')':
		tok = newToken(token.CLOSEPARENTHESIS, l.currentChar)
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
	case '{':
		tok = newToken(token.OPENBRACE, l.currentChar)
	case '}':
		tok = newToken(token.CLOSEBRACE, l.currentChar)
	case '[':
		tok = newToken(token.OPENBRACKET, l.currentChar)
	case ']':
		tok = newToken(token.CLOSEBRACKET, l.currentChar)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case ':':
		tok = newToken(token.COLON, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.currentChar) {
			tok.Type = token.INTEGER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currentChar)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' ||
		l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.currentChar == '"' || l.currentChar == 0 {
			break
		}
	}

	return l.input[position:l.position]
}
