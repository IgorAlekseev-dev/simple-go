package main

type TokenType int

const (
	TokenEOF     TokenType = iota // конец файла
	TokenIllegal                  // некорректный символ

	TokenIdent     // Идентификатор
	TokenIntLit    // 9
	TokenFloatLit  // 3.14
	TokenStringLit // "строка"

	TokenPackage // package
	TokenMain    // main
	TokenFunc    // func
	TokenVar     // var
	TokenConst   // const
	TokenInt     // int
	TokenFloat   // float
	TokenString  // string
	TokenBool    // bool
	TokenTrue    // true
	TokenFalse   // false
	TokenIf      // if
	TokenElse    // else
	TokenFor     // for
	TokenScan    // scan
	TokenPrint   // print
	TokenPrintln // println

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

func evaluate(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return TokenIdent
}
