package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // 입력에서 현재 읽는 위치 (현재 문자)
	readPosition int  // 입력에서 현재 "읽는" 위치 (현재 문자 다음을 가르킴)
	ch           byte // 현재 조사하는 "문자"

	/*		position 과 readPosition 둘다 필요한 이유?					*
	 *		현재 문자를 "보존"하면서, 다음 문자를 미리 봐야 할 떄가 있기 때문이다.	*/
}

// New : 소스코드를 받고(string) Lexer 에 저장한다.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // l 을 초기화시킨다
	return l
}

// readChar : 읽는 문자를 다음으로 바꾸고, position, readPosition 을 다음으로 바꾸는 함수
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 파일의 끝(EOF) 혹은 아직 아무것도 읽지 않음 의 2가지 상태
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1 // 1 움직이는 방법은 ASCII 문자만 적용된다. UTF-8에는 적용 xxx!!

	// 과제 : 유니코드(와 이모티콘) 지원은 어떻게 할까?
	// unicode/utf8 을 활용하면, for 문을 통해서 십진수와 사이즈로 반환가능하다.
}

// NextToken : Lexer 가 보고있는 문자에 대응하는 Token 을 반환하고, 다음 Token 으로 Lexer 를 업데이트 하는 함수
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// 공백 무시하기
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		//tok = newToken(token.EOF, l.ch)
		tok.Literal = "" // 문자열 마지막이라 \x00 이 나온다. 우리는 "" 을 의도하고 있다.
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) { // identifier 와 keyword 식별하는 코드 test2 에 의해 추가한다.
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // readIdentifier 에 의해(내부의 readChar 에 의해) l 은 내부적으로 position 이 이동했기 때문에 조기종료해야 한다.
		} else if isDigit(l.ch) { // 숫자 읽기 (문자 읽기와 유사하다.)
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// newToken : Token 생성기
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier : identigier 혹은 keyword 를 인식하고 문자(character) 를 반환한다. cf) 문자와 글자(letter)는 다르다.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/*
isLetter : 언어가 인식하는 글자(letter) 를 인식한다. 해당 언어는 영어 소문자, 대문자와 언더바(_) 를 글자로 인정한다.
만약 다른 글자를 추가하고 싶다면 조건문에 추가하면 된다. ?, ! 를 추가하는 경우도 있다.
*/
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit : 해당 글자(letter) 가 숫자인지 판단
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// skipWhitespace : 언어가 지정한 공백문자를 Lexer 가 무시하게 해준다.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
