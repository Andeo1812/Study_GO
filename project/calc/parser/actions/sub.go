package actions

import (
	"Modules/project/calc/parser/configs"
	"fmt"
)

func actionSub(accum float64, expression string) (float64, int) {
	fmt.Println("Sub", accum, expression)
	var pos int = 0

	for pos < len(expression) {
		if string(expression[pos]) == configs.Lex.CloseParen {
			return accum, pos + 1
		}

		addition, offset := actionOpenOperand(expression[pos:])
		pos += offset

		if len(expression) == pos {
			return accum - addition, pos
		}

		symbol := string(expression[pos])
		pos++

		switch symbol {
		case configs.Lex.Plus:
			res, offset := actionSum(-addition, expression[pos:])
			accum += res
			pos += offset
		case configs.Lex.Minus:
			accum -= addition
		case configs.Lex.Multiply:
			res, offset := actionMul(-addition, expression[pos:])
			accum += res
			pos += offset
		case configs.Lex.Divide:
			res, offset := actionDiv(-addition, expression[pos:])
			accum += res
			pos += offset
		case configs.Lex.CloseParen:
			if expression[pos:] == "" ||
				string(expression[pos]) == configs.Lex.Plus || string(expression[pos]) == configs.Lex.Minus || string(expression[pos]) == configs.Lex.Multiply || string(expression[pos]) == configs.Lex.Divide {
				return accum + addition, pos
			}
			return accum + addition, pos - 1

		}
	}

	return accum, pos
}
