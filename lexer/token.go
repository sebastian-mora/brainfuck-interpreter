package lexer

type TokenType int

const (
	INCREMENT TokenType = iota
	DECREMENT
	SHIFT_LEFT
	SHIFT_RIGHT
	OUTPUT
	INPUT
	OPEN_LOOP
	CLOSE_LOOP
	NEW_LINE
	EOF
)

type Token struct {
	TokenType TokenType
	Literal   string
}
