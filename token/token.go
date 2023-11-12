package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
} // 타입에는 자료형과 그 값이 저장되어야 한다.
