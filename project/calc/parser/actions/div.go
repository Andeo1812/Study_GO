package actions

import (
	"Modules/project/calc/parser/configs"
	"fmt"
)

func actionDiv(accum float64, expression string) (float64, int) {
	fmt.Println("Div:", accum, expression)
	var pos int = 0

	for pos < len(expression) {
		number, lengthNumber := GetNumber(expression[pos:])
		pos += lengthNumber

		if len(expression) == pos {
			return accum / number, pos
		}

		symbol := string(expression[pos])
		pos++

		switch symbol {
		case configs.Lex.Plus:
			accum /= number

			res, offset := actionSum(accum, expression[pos:])

			pos += offset

			return res, pos
		case configs.Lex.Minus:
			accum /= number

			res, offset := actionSub(accum, expression[pos:])

			pos += offset

			return res, pos
		case configs.Lex.Multiply:
			res, offset := actionMul(accum/number, expression[pos:])

			accum = res

			pos += offset
		case configs.Lex.Divide:
			accum /= number
		}
	}

	return accum, pos
}
