package actions

import (
	"Modules/project/calc/parser/configs"
	"fmt"
)

func actionSum(accum float64, expression string) (float64, int) {
	fmt.Println("Sum:", accum, expression)
	var pos int = 0

	for pos < len(expression) {
		if string(expression[pos]) == configs.Lex.CloseParen {
			return accum, pos + 1
		}

		addition, offset := actionOpenOperand(expression[pos:])
		pos += offset

		if len(expression) == pos {
			return accum + addition, pos
		}

		symbol := string(expression[pos])
		pos++

		fmt.Println("Sum:", symbol)

		switch symbol {
		case configs.Lex.Plus:
			accum += addition
		case configs.Lex.Minus:
			res, offset := actionSub(addition, expression[pos:])
			accum += res
			pos += offset
		case configs.Lex.Multiply:
			res, offset := actionMul(addition, expression[pos:])
			accum += res
			pos += offset
		case configs.Lex.Divide:
			res, offset := actionDiv(addition, expression[pos:])
			accum += res
			pos += offset
		case configs.Lex.CloseParen:
			fmt.Println("SumClose:", expression[pos:])
			if expression[pos:] == "" ||
				string(expression[pos]) == configs.Lex.Plus || string(expression[pos]) == configs.Lex.Minus || string(expression[pos]) == configs.Lex.Multiply || string(expression[pos]) == configs.Lex.Divide {
				fmt.Println("SumCloseBingo:", expression[pos:])
				return accum + addition, pos
			}
			return accum + addition, pos - 1
		}
	}

	return accum, pos
}
