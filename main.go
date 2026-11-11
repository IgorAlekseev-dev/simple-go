package main

import (
	"fmt"
)

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenIllegal

	TokenIdent
	TokenIntLit
	TokenFloatLit
	TokenStringLit

	TokenPackage
	TokenMain
	TokenFunc
	TokenVar
	TokenConst
	TokenInt
	TokenFloat
	TokenString
	TokenBool
	TokenTrue
	TokenFalse
	TokenIf
	TokenElse
	TokenFor
	TokenScan
	TokenPrint
	TokenPrintln

	TokenPlus    // +
	TokenMinus   // -
	TokenStar    // *
	TokenSlash   // /
	TokenPercent // %
	TokenInc     // ++
	TokenDec     // --
	TokenAssign  // =

	TokenEqual    // ==
	TokenNotEqual // !=
	TokenLess     // <
	TokenLessEq   // <=
	TokenGreat    // >
	TokenGreatEq  // >=

	TokenAnd // &&
	TokenOr  // ||
	TokenNot // !

	TokenLParen    // (
	TokenRParen    // )
	TokenLBrace    // {
	TokenRBrace    // }
	TokenComma     // ,
	TokenSemicolon // ;
)

var keywords = map[string]TokenType{
	"package": TokenPackage,
	"main":    TokenMain,
	"func":    TokenFunc,
	"var":     TokenVar,
	"const":   TokenConst,
	"int":     TokenInt,
	"float":   TokenFloat,
	"string":  TokenString,
	"bool":    TokenBool,
	"true":    TokenTrue,
	"false":   TokenFalse,
	"if":      TokenIf,
	"else":    TokenElse,
	"for":     TokenFor,
	"scan":    TokenScan,
	"print":   TokenPrint,
	"println": TokenPrintln,
}

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
	Value   int
}

var comment bool = false

type Lexer struct {
	SourceProgram string // Исходный текст программы
	Index         int
	RowNumber     int // Номер строки
	RowIndex      int // Номер текущего символа
	CurrentSymbol byte
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) getChar() byte {
	if l.Index >= len(l.SourceProgram) {
		return 0
	}

	ch := l.SourceProgram[l.Index]
	l.Index++

	if ch == '\n' {
		l.RowNumber++
		l.RowIndex = 0
	} else {
		l.RowIndex++
	}

	l.CurrentSymbol = ch

	return l.CurrentSymbol
}

func (l *Lexer) peekChar() byte {
	if l.Index >= len(l.SourceProgram) {
		return 0
	}

	ch := l.SourceProgram[l.Index+1]

	if ch == '\n' {
		l.RowNumber++
		l.RowIndex = 0
	} else {
		l.RowIndex++
	}

	l.CurrentSymbol = ch

	return l.CurrentSymbol
}

func evaluate(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return TokenIdent
}

func main() {
	lexer := Lexer{
		SourceProgram: "package main\nfunc main(){}",
		RowNumber:     1,
	}

	for {
		ch := lexer.getChar()

		if ch == 0 {
			break
		}
		fmt.Printf("Символ %c (строка %d, позиция %d)\n", ch, lexer.RowNumber, lexer.RowIndex)
	}
}
