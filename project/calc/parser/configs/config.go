package configs

type Lexemes struct {
	Plus       string
	Minus      string
	Multiply   string
	Divide     string
	OpenParen  string
	CloseParen string
	Digit      string
}

var Lex Lexemes = Lexemes{"+", "-", "*", "/", "(", ")", "0123456789"}
