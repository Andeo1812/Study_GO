package actions

import (
	"Modules/project/calc/parser/configs"
	"fmt"
)

func actionSub(accum float64, expression string) (float64, int) {
	fmt.Println("Sub", accum, expression)
	var pos int = 0

	for pos < len(expression) {
		number, lengthNumber := GetNumber(expression[pos:])
		pos += lengthNumber

		if len(expression) == pos {
			return accum - number, pos
		}

		symbol := string(expression[pos])
		pos++

		var res float64
		var offset int

		switch symbol {
		case configs.Lex.Plus:
			res, offset = actionSum(-number, expression[pos:])
		case configs.Lex.Minus:
			accum -= number
		case configs.Lex.Multiply:
			res, offset = actionMul(-number, expression[pos:])
		case configs.Lex.Divide:
			res, offset = actionDiv(-number, expression[pos:])
		}

		accum += res
		pos += offset
	}

	return accum, pos
}
