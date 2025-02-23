package lexer

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal = "Illegal"
	Eof     = "EOF"

	Assign     = "="
	ParamStart = "?"

	LeftParenthesis  = "("
	RightParenthesis = ")"
	Comma            = ","
	Quote            = "'"

	Equals             = "equals"
	LessThan           = "lessThan"
	LessOrEqual        = "lessOrEqual"
	GreaterThan        = "greaterThan"
	GreaterThanOrEqual = "greaterOrEqual"
	Contains           = "contains"
	StartWith          = "startsWith"
	EndsWith           = "endsWith"
)

type Lexer struct {
	input        string // input for lexer
	position     int    //current position in the input
	readPosition int    // current reading position in the input
	ch           byte   //current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() Token {

}
