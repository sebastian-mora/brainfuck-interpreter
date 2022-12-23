package lexer

import (
	"errors"
	"fmt"
)

func Lex(program string) ([]Token, error) {
	tokens := make([]Token, len(program))

	for i, char := range program {

		switch char {
		case '+':
			tokens[i] = Token{TokenType: INCREMENT, Literal: string(char)}
		case '-':
			tokens[i] = Token{TokenType: DECREMENT, Literal: string(char)}
		case '>':
			tokens[i] = Token{TokenType: SHIFT_RIGHT, Literal: string(char)}
		case '<':
			tokens[i] = Token{TokenType: SHIFT_LEFT, Literal: string(char)}
		case '.':
			tokens[i] = Token{TokenType: OUTPUT, Literal: string(char)}
		case ',':
			tokens[i] = Token{TokenType: INPUT, Literal: string(char)}
		case '[':
			tokens[i] = Token{TokenType: OPEN_LOOP, Literal: string(char)}
		case ']':
			tokens[i] = Token{TokenType: CLOSE_LOOP, Literal: string(char)}
		default:
			err := fmt.Sprintf("Error at %d: Invalid character \"%s\" ", i, string(char))
			return nil, errors.New(err)
		}
	}

	tokens = append(tokens, Token{TokenType: EOF})
	return tokens, nil
}
