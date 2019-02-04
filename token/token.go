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
	INTEGER          = "INT"
	ASSIGNMENT       = "="
	PLUS             = "+"
	COMMA            = ","
	SEMICOLON        = ";"
	OPENPARENTHESIS  = "("
	CLOSEPARENTHESIS = ")"
	OPENBRACE        = "{"
	CLOSEBRACE       = "}"
	FUNCTION         = "FUNCTION"
	LET              = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	return IDENTIFIER
}
