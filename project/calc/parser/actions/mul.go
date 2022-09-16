package actions

import (
	"Modules/project/calc/parser/configs"
	"fmt"
)

func actionMul(accum float64, expression string) (float64, int) {
	fmt.Println("Mul:", accum, expression)
	var pos int = 0

	for pos < len(expression) {
		number, lengthNumber := GetNumber(expression[pos:])
		pos += lengthNumber

		if len(expression) == pos {
			return accum * number, pos
		}

		symbol := string(expression[pos])
		pos++

		switch symbol {
		case configs.Lex.Plus:
			res, offset := actionSum(accum*number, expression[pos:])
			pos += offset

			return res, pos
		case configs.Lex.Minus:
			res, offset := actionSub(accum*number, expression[pos:])
			pos += offset

			return res, pos
		case configs.Lex.Multiply:
			accum *= number
		case configs.Lex.Divide:
			res, offset := actionDiv(accum*number, expression[pos:])
			accum = res
			pos += offset
		}
	}

	return accum, pos
}
