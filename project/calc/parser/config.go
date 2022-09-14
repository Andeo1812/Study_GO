package parser

type lexemes struct {
	plus       string
	minus      string
	multiply   string
	divide     string
	openParen  string
	closeParen string
	digit      string
}

var lex lexemes = lexemes{"+", "-", "*", "/", "(", ")", "123456789"}
