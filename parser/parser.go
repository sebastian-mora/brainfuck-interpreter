package parser

import (
	"errors"

	"github.com/sebastian-mora/brainfuck-interpreter/lexer"
)

/*
Context Free Grammar: BrainFuck

S -> t | L
L -> [S]

Where "t" is a terminal character

>>
----
-
	\
		-
		\
			-
			\
				-

+[-]
----
+
	[
		-
			]
*/

type Operation struct {
	Token lexer.Token
	Count int
	Jump  int // set only if loop
}

var stack Stack

func Parse(tokens []lexer.Token) ([]Operation, error) {
	operations := make([]Operation, len(tokens))
	for i, token := range tokens {
		switch token.TokenType {
		case lexer.DECREMENT,
			lexer.INCREMENT,
			lexer.SHIFT_LEFT,
			lexer.SHIFT_RIGHT,
			lexer.INPUT,
			lexer.OUTPUT:
			operations[i] = createOperation(i, tokens)

		case lexer.OPEN_LOOP:
			stack.Push(i)
			operations[i] = createLoop(-1, lexer.OPEN_LOOP) // jump position is unknown for now

		case lexer.CLOSE_LOOP:
			if stack.IsEmpty() {
				return nil, errors.New("closing bracket does not match any opening bracket")
			}
			open_loop, _ := stack.Pop()
			operations[i] = createLoop(open_loop, lexer.CLOSE_LOOP) // send end loop jump
			operations[open_loop].Jump = i                          // set the open loop jmp to end

		case lexer.EOF:
			if !stack.IsEmpty() {
				return nil, errors.New("unmatched opening bracket")
			}
			continue

		}
	}

	return operations, nil
}

func createOperation(index int, tokens []lexer.Token) Operation {
	return Operation{Token: tokens[index]}
}

func createLoop(jump int, token lexer.TokenType) Operation {
	return Operation{Token: lexer.Token{TokenType: token}, Jump: jump}
}

func eatOperations(start_index int, tokens []lexer.Token) int {
	start_token := tokens[start_index]
	i := start_index + 1
	count := 1
	for i < len(tokens) && start_token == tokens[i] {
		count++
		i++
	}

	return count
}
