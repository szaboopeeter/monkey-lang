package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL          = "ILLEGAL"
	EOF              = "EOF"
	IDENTIFIER       = "IDENTIFIER"
	INT              = "INT"
	ASSIGNMENT       = "="
	PLUS             = "+"
	COMMA            = ","
	SEMICOLON        = ";"
	OPENPARENTHESIS  = "("
	CLOSEPARENTHESIS = ")"
	OPENBRACE        = "{"
	CLOSEBRACE       = "}"
	FUNCTION         = "FUNCTION"
	VARIABLE         = "LET"
)
