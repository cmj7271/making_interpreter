package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
} // 타입에는 자료형과 그 값이 저장되어야 한다.

const (
	// error
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// indentifier
	IDENT = "IDENT"
	INT   = "INT"

	// unary operator (단항 연산자)
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// binary operator(이항 연산자)
	EQ     = "=="
	NOT_EQ = "!="
	/*
		생각해볼 거리 : 현재까지는 이항연산자가 2개뿐이기 때문에
		Lexer 에서 case 에서 일일이 구분해도 충분하다.
		만약 이항 연산자가 늘어난다면?
		함수를 만들어 추상화하고 싶어진다. 예컨데 makeTwoCharToken 같은 이름으로
	*/

	LT = "<"
	GT = ">"

	// end of statement
	COMMA     = ","
	SEMICOLON = ";"

	// code block
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// keywords : 예약어들을 리터럴은 key 로, TokenType 은 value 로 저장한다.
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent : 예약어 혹은 식별자인지 판단하고 그 타입을 반환
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
