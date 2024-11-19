package main

import (
	"os"

	"github.com/AungKyawPhyo1142/go-compiler/src/lexer"
)

func main() {
	// read a file
	bytes, _ := os.ReadFile("./example/00.lang")

	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}

}
