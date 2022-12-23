package interpreter

import (
	"fmt"

	"github.com/sebastian-mora/brainfuck-interpreter/lexer"
	"github.com/sebastian-mora/brainfuck-interpreter/parser"
)

type Brainfuck struct {
	Instructions []parser.Operation
	Memory       [30000]int
	dp           int
	ip           int
}

func NewBrainfuck(code string) (*Brainfuck, error) {
	lex, _ := lexer.Lex(code)
	parseTree, err := parser.Parse(lex)

	if err != nil {
		return nil, err
	}
	return &Brainfuck{Instructions: parseTree}, nil
}

func (bf *Brainfuck) Interpret() {
	bf.ip = 0                          // set instruction pointer to 0
	for bf.ip < len(bf.Instructions) { // this might cause a bug if the loop is at the end of the program
		op := bf.Instructions[bf.ip] // get the current instruction
		bf.ip++
		switch op.Token.TokenType {
		case lexer.INCREMENT:
			bf.Memory[bf.dp]++
		case lexer.DECREMENT:
			bf.Memory[bf.dp]--
		case lexer.SHIFT_LEFT:
			bf.dp--
		case lexer.SHIFT_RIGHT:
			bf.dp++
		case lexer.OUTPUT:
			fmt.Print(bf.Memory[bf.dp])
		case lexer.INPUT:
			var input int
			_, _ = fmt.Scan(&input)
			bf.Memory[bf.dp] = input
		case lexer.OPEN_LOOP:
			if bf.Memory[bf.dp] <= 0 {
				bf.ip = op.Jump + 1
			}
		case lexer.CLOSE_LOOP:
			bf.ip = op.Jump

		}
	}
}
