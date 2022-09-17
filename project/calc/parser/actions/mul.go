package actions

import (
	"Modules/project/calc/parser/configs"
	"fmt"
)

func actionMul(accum float64, expression string) (float64, int) {
	fmt.Println("Mul:", accum, expression)
	var pos int = 0

	for pos < len(expression) {
		fmt.Println("MulBegin", string(expression[pos:]))
		fmt.Println("MulBegin", string(expression[pos]))
		if string(expression[pos]) == configs.Lex.CloseParen {
			return accum, pos + 1
		}

		addition, offset := actionOpenOperand(expression[pos:])
		pos += offset

		if len(expression) == pos {
			return accum * addition, pos
		}

		symbol := string(expression[pos])
		pos++

		fmt.Println("Mul:", symbol)

		switch symbol {
		case configs.Lex.Plus:
			res, offset := actionSum(accum*addition, expression[pos:])
			pos += offset

			return res, pos
		case configs.Lex.Minus:
			res, offset := actionSub(accum*addition, expression[pos:])
			pos += offset

			return res, pos
		case configs.Lex.Multiply:
			accum *= addition
		case configs.Lex.Divide:
			res, offset := actionDiv(accum*addition, expression[pos:])
			accum = res
			pos += offset
		case configs.Lex.CloseParen:
			pos--
			fmt.Println("MULClose:", string(expression[pos]))
			return accum * addition, pos
		}
	}

	return accum, pos
}
