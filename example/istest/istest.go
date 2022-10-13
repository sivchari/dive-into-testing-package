package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	rune, _ := utf8.DecodeRuneInString("Testa"[len("Test"):])
	fmt.Println(string(rune))

	rune, _ = utf8.DecodeRuneInString("TestA"[len("Test"):])
	fmt.Println(string(rune))

	rune, _ = utf8.DecodeRuneInString("Test_"[len("Test"):])
	fmt.Println(string(rune))
}
