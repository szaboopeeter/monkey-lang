package lexer

import (
	"monkey-lang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, z) {
		x + y;
	};
	
	let result = add(five, ten);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGNMENT, "="},
		{token.INTEGER, "int"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGNMENT, "="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGNMENT, "="},
		{token.FUNCTION, "fn"},
		{token.OPENPARENTHESIS, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.CLOSEPARENTHESIS, ")"},
		{token.OPENBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.CLOSEBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGNMENT, "="},
		{token.IDENTIFIER, "add"},
		{token.OPENPARENTHESIS, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.CLOSEPARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}
