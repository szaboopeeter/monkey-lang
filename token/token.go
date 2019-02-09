package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INT"

	ASSIGNMENT  = "="
	PLUS        = "+"
	MINUS       = "-"
	BANG        = "!"
	ASTERISK    = "*"
	SLASH       = "/"
	LESSTHAN    = "<"
	GREATERTHAN = ">"

	COMMA            = ","
	SEMICOLON        = ";"
	OPENPARENTHESIS  = "("
	CLOSEPARENTHESIS = ")"
	OPENBRACE        = "{"
	CLOSEBRACE       = "}"

	EQUAL    = "=="
	NOTEQUAL = "!="

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return IDENTIFIER
}
