package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
} // 타입에는 자료형과 그 값이 저장되어야 한다.

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// keywords : 예약어들을 리터럴은 key 로, TokenType 은 value 로 저장한다.
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent : 예약어 혹은 식별자인지 판단하고 그 타입을 반환
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
