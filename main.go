package main

import (
	"fmt"
)

func main() {
	lexer := NewLexer(`package main
	func main() {
	    var a int = 10;
	    var b float = 3.14;
		var msg string = "hello";
		a++;
		b--;
		if a != 10 && b >= 2.0 {
			msg = "ok";
		}
	}`)

	fmt.Println("Лексема:     | Тип: | Значение: | Строка: | Позиция:")
	for {
		tok := lexer.getToken()
		fmt.Printf("%-12s | %-4d | %-9d | %-7d | %-8d\n", tok.Literal, tok.Type, tok.Value, tok.Line, tok.Col)
		if tok.Type == TokenEOF {
			break
		}
	}
}
