package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	// test1

	//input := `=+(){},;`
	//
	//tests := []struct {
	//	expectedType    token.TokenType
	//	expectedLiteral string
	//}{
	//	{token.ASSIGN, "="},
	//	{token.PLUS, "+"},
	//	{token.LPAREN, "("},
	//	{token.RPAREN, ")"},
	//	{token.LBRACE, "{"},
	//	{token.RBRACE, "}"},
	//	{token.COMMA, ","},
	//	{token.SEMICOLON, ";"},
	//	{token.EOF, ""},
	//}

	// test 2
	//	input := `let five = 5;
	//let ten = 10;
	//
	//let add = fn(x, y) {
	//	x + y;
	//};
	//
	//let result = add(five, ten);
	//`
	//
	//	tests := []struct {
	//		expectedType    token.TokenType
	//		expectedLiteral string
	//	}{
	//		{token.LET, "let"},
	//		{token.IDENT, "five"},
	//		{token.ASSIGN, "="},
	//		{token.INT, "5"},
	//		{token.SEMICOLON, ";"},
	//		{token.LET, "let"},
	//		{token.IDENT, "ten"},
	//		{token.ASSIGN, "="},
	//		{token.INT, "10"},
	//		{token.SEMICOLON, ";"},
	//		{token.LET, "let"},
	//		{token.IDENT, "add"},
	//		{token.ASSIGN, "="},
	//		{token.FUNCTION, "fn"},
	//		{token.LPAREN, "("},
	//		{token.IDENT, "x"},
	//		{token.COMMA, ","},
	//		{token.IDENT, "y"},
	//		{token.RPAREN, ")"},
	//		{token.LBRACE, "{"},
	//		{token.IDENT, "x"},
	//		{token.PLUS, "+"},
	//		{token.IDENT, "y"},
	//		{token.SEMICOLON, ";"},
	//		{token.RBRACE, "}"},
	//		{token.SEMICOLON, ";"},
	//		{token.LET, "let"},
	//		{token.IDENT, "result"},
	//		{token.ASSIGN, "="},
	//		{token.IDENT, "add"},
	//		{token.LPAREN, "("},
	//		{token.IDENT, "five"},
	//		{token.COMMA, ","},
	//		{token.IDENT, "ten"},
	//		{token.RPAREN, ")"},
	//		{token.SEMICOLON, ";"},
	//	}

	// test3 : -, /, *, <, > 추가하기
	input := `-/*<>!
true false if else return 
== !=`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.BANG, "!"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.RETURN, "return"},
		{token.EQ, "=="},
		{token.NOT_EQ, "!="},
	}

	l := New(input) // golang 의 고유 문법이 아니라 새로운 함수임!!

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}

	}

}
