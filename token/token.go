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
	STRING     = "STRING"

	ASSIGNMENT  = "="
	PLUS        = "+"
	MINUS       = "-"
	BANG        = "!"
	ASTERISK    = "*"
	SLASH       = "/"
	LESSTHAN    = "<"
	GREATERTHAN = ">"

	COMMA            = ","
	COLON            = ":"
	SEMICOLON        = ";"
	OPENPARENTHESIS  = "("
	CLOSEPARENTHESIS = ")"
	OPENBRACE        = "{"
	CLOSEBRACE       = "}"
	OPENBRACKET      = "["
	CLOSEBRACKET     = "]"

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
