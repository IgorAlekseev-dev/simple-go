package main

import (
	"fmt"
	"os"
)

func main() {
	symTable := NewSymbolTable()

	source, err := os.ReadFile("program.sgo")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	lexer := NewLexer(string(source), symTable)

	fmt.Println("Список лексем:")
	fmt.Println("Лексема:     | Тип: | Значение: | Строка: | Позиция:")
	fmt.Println("-----------------------------------------------------")
	for {
		tok := lexer.getToken()
		fmt.Printf("%-12s | %-4d | %-9d | %-7d | %-8d\n", tok.Literal, tok.Type, tok.Value, tok.Line, tok.Col)
		if tok.Type == TokenEOF {
			break
		}
	}
	symTable.Print()
}
