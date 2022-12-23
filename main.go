package main

import (
	"github.com/sebastian-mora/brainfuck-interpreter/interpreter"
)

func main() {

	bf, _ := interpreter.NewBrainfuck("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")

	// for _, v := range bf.Instructions {
	// 	fmt.Print(v.Token.Literal, v.Jump, "\n")
	// }

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(2)
	// }

	bf.Interpret()

}
