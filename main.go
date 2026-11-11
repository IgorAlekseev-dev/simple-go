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

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func evaluate(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return TokenIdent
}

type Lexer struct {
	SourceProgram string // Исходный текст программы
	Index         int    // Индекс текущего символа в коде
	RowNumber     int    // Номер строки
	RowIndex      int    // Номер текущего символа в строке
	CurrentSymbol byte   // Текущий символ
}

func NewLexer(source string) *Lexer {
	l := &Lexer{
		SourceProgram: source,
		Index:         0,
		RowNumber:     1,
		RowIndex:      0,
	}
	l.getChar()
	return l
}

func (l *Lexer) getChar() byte {
	if l.Index >= len(l.SourceProgram) {
		l.CurrentSymbol = 0
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

	return l.SourceProgram[l.Index]
}

func (l *Lexer) skipWhitespace() {
	for l.CurrentSymbol == ' ' ||
		l.CurrentSymbol == '\t' ||
		l.CurrentSymbol == '\n' ||
		l.CurrentSymbol == '\r' {
		l.getChar()
	}
}

func (l *Lexer) readIdentifier() string {
	var word string

	for isLetter(l.CurrentSymbol) || isDigit(l.CurrentSymbol) {
		word += string(l.CurrentSymbol)
		l.getChar()
	}

	return word
}

func (l *Lexer) readNumber() (string, TokenType) {
	tokType := TokenIntLit
	number := ""

	for isDigit(l.CurrentSymbol) {
		number += string(l.CurrentSymbol)
		l.getChar()
	}
	if l.CurrentSymbol == '.' && isDigit(l.peekChar()) {
		tokType = TokenFloatLit
		number += "."
		l.getChar()
		for isDigit(l.CurrentSymbol) {
			number += string(l.CurrentSymbol)
			l.getChar()
		}
	}

	return number, tokType
}

func (l *Lexer) getToken() Token {
	l.skipWhitespace()
	line := l.RowNumber
	col := l.RowIndex

	if isLetter(l.CurrentSymbol) {
		literal := l.readIdentifier()
		tokType := evaluate(literal)
		return Token{
			Type:    tokType,
			Literal: literal,
			Line:    line,
			Col:     col,
		}
	} else if isDigit(l.CurrentSymbol) {
		number, tokType := l.readNumber()
		return Token{
			Type:    tokType,
			Literal: number,
			Line:    line,
			Col:     col,
		}
	}

	switch l.CurrentSymbol {
	case '(':
		tokType := TokenLParen
		l.getChar()
		return Token{
			Type:    tokType,
			Literal: "(",
			Line:    line,
			Col:     col,
		}
	case ')':
		tokType := TokenRParen
		l.getChar()
		return Token{
			Type:    tokType,
			Literal: ")",
			Line:    line,
			Col:     col,
		}
	case '{':
		tokType := TokenLBrace
		l.getChar()
		return Token{
			Type:    tokType,
			Literal: "{",
			Line:    line,
			Col:     col,
		}
	case '}':
		tokType := TokenRBrace
		l.getChar()
		return Token{
			Type:    tokType,
			Literal: "}",
			Line:    line,
			Col:     col,
		}
	case ';':
		tokType := TokenSemicolon
		l.getChar()
		return Token{
			Type:    tokType,
			Literal: ";",
			Line:    line,
			Col:     col,
		}
	case '=':
		if l.peekChar() == '=' {
			l.getChar()
			l.getChar()
			return Token{
				Type:    TokenEqual,
				Literal: "==",
				Line:    line,
				Col:     col,
			}
		}
		l.getChar()
		return Token{
			Type:    TokenAssign,
			Literal: "=",
			Line:    line,
			Col:     col,
		}
	case 0:
		tokType := TokenEOF
		return Token{
			Type:    tokType,
			Literal: "EOF",
			Line:    line,
			Col:     col,
		}
	default:
		badChar := string(l.CurrentSymbol)
		l.getChar()
		return Token{
			Type:    TokenIllegal,
			Literal: badChar,
			Line:    line,
			Col:     col,
		}
	}
}

func main() {
	lexer := NewLexer(`package main
	func main() {
	    var a int = 10;
	    var b float = 3.14;
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
