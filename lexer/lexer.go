package lexer

type Lexer struct {
	input        string
	position     int  // 입력에서 현재 읽는 위치 (현재 문자)
	readPosition int  // 입력에서 현재 "읽는" 위치 (현재 문자 다음을 가르킴)
	ch           byte // 현재 조사하는 "문자"

	/*		position 과 readPosition 둘다 필요한 이유?
			현재 문자를 "보존"하면서, 다음 문자를 미리 봐야 할 떄가 있기 때문이다.
	*/
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() { // 읽는 문자를 다음으로 바꾸고, position 들을 다음으로 바꾸는 함수
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
