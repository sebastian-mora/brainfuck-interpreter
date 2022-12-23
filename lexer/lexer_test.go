package lexer

import (
	"reflect"
	"testing"
)

func TestValidLex(t *testing.T) {
	program := "><"

	got, err := Lex(program)

	want := []Token{{TokenType: SHIFT_RIGHT, Literal: ">"}, {TokenType: SHIFT_LEFT, Literal: "<"}}

	if err != nil {
		t.Errorf("%s ", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func TestInvalidLex(t *testing.T) {
	program := ">>>>>><<<<>><<>>...4.."

	_, err := Lex(program)

	if err == nil {
		t.Errorf("Failed to detect invalid character")
	}

}
