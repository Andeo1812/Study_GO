package actions

import (
	"Modules/project/calc/parser/configs"
	"errors"
)

func actionSum(accum float64, expression string) (float64, int, error) {
	var pos int = 0

	for pos < len(expression) {
		addition, offset, errGetOperand := actionOpenOperand(expression[pos:])
		if errGetOperand != nil {
			return 0, 0, errGetOperand
		}

		pos += offset

		if len(expression) == pos {
			return accum + addition, pos, nil
		}

		symbol := string(expression[pos])
		pos++

		switch symbol {
		case configs.Lex.Plus:
			accum += addition
		case configs.Lex.Minus:
			res, offset, errSub := actionSub(addition, expression[pos:])
			if errSub != nil {
				return 0, 0, errSub
			}

			accum += res
			pos += offset
		case configs.Lex.Multiply:
			res, offset, errMul := actionMul(addition, expression[pos:])
			if errMul != nil {
				return 0, 0, errMul
			}

			accum += res
			pos += offset
		case configs.Lex.Divide:
			res, offset, errDiv := actionDiv(addition, expression[pos:])
			if errDiv != nil {
				return 0, 0, errDiv
			}

			accum += res
			pos += offset
		default:
			return 0, 0, errors.New(wrongSymbol(symbol))
		}
	}

	return accum, pos, nil
}
