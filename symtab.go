package main

import "fmt"

type SymbolKind string

const (
	KindVar   SymbolKind = "Переменная"
	KindConst SymbolKind = "Константа"
)

type Symbol struct {
	Index int        // Номер в таблице
	Name  string     // Имя переменной или литерал константы
	Kind  SymbolKind // Переменная или константа
	Type  string     // int, float, string, bool или "-"
	Value string     // Значение константы или "-"
}

type SymbolTable struct {
	Symbols []Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		Symbols: make([]Symbol, 0),
	}
}

func (st *SymbolTable) Evaluate(tok Token) int {
	for _, sym := range st.Symbols {
		if sym.Name == tok.Literal {
			return sym.Index
		}
	}

	newIndex := len(st.Symbols)
	newSym := Symbol{
		Index: newIndex,
		Name:  tok.Literal,
		Value: "-",
		Type:  "-",
	}

	switch tok.Type {
	case TokenIdent:
		newSym.Kind = KindVar

	case TokenIntLit:
		newSym.Kind = KindConst
		newSym.Type = "int"
		newSym.Value = tok.Literal

	case TokenFloatLit:
		newSym.Kind = KindConst
		newSym.Type = "float"
		newSym.Value = tok.Literal

	case TokenStringLit:
		newSym.Kind = KindConst
		newSym.Type = "string"
		newSym.Value = tok.Literal

	case TokenTrue, TokenFalse:
		newSym.Kind = KindConst
		newSym.Type = "bool"
		newSym.Value = tok.Literal

	default:
		return 0
	}

	st.Symbols = append(st.Symbols, newSym)
	return newIndex
}

func (st *SymbolTable) Print() {
	fmt.Println()
	fmt.Println("Таблица символов (переменных и констант):")
	fmt.Println()
	fmt.Println("Индекс | Имя            | Вид          | Тип     | Значение")
	fmt.Println("-----------------------------------------------------------------")
	for _, sym := range st.Symbols {
		fmt.Printf("%-6d | %-14s | %-12s | %-7s | %-10s\n",
			sym.Index, sym.Name, sym.Kind, sym.Type, sym.Value)
	}
	fmt.Println()
}
