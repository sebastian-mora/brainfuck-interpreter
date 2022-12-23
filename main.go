package main

import (
	"fmt"
	"os"

	"github.com/sebastian-mora/brainfuck-interpreter/interpreter"
)

func main() {

	bf, err := interpreter.NewBrainfuck("++++++++[>++++[>++>[+++>+\n++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")

	// for _, v := range bf.Instructions {
	// 	fmt.Print(v.Token.Literal, v.Jump, "\n")
	// }

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	bf.Interpret()

}
