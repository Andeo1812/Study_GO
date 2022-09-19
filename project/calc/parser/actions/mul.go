package actions

import (
	"Calc/project/calc/parser/configs"
	"errors"
)

func actionMul(accum float64, expression string) (float64, int, error) {
	var pos int = 0

	for pos < len(expression) {
		strings.s
		addition, offset, errGetOperand := actionGetOperand(expression[pos:])
		if errGetOperand != nil {
			return 0, 0, errGetOperand
		}

		pos += offset

		if len(expression) == pos {
			return accum * addition, pos, nil
		}

		symbol := string(expression[pos])
		pos++

		switch symbol {
		case configs.Lex.Plus:
			res, offset, errSum := actionSum(accum*addition, expression[pos:])
			if errSum != nil {
				return 0, 0, errSum
			}

			pos += offset

			return res, pos, nil
		case configs.Lex.Minus:
			res, offset, errSub := actionSub(accum*addition, expression[pos:])
			if errSub != nil {
				return 0, 0, errSub
			}

			pos += offset

			return res, pos, nil
		case configs.Lex.Multiply:
			accum *= addition
		case configs.Lex.Divide:
			res, offset, errDiv := actionDiv(accum*addition, expression[pos:])
			if errDiv != nil {
				return 0, 0, errDiv
			}

			accum = res
			pos += offset
		default:
			return 0, 0, errors.New(wrongSymbol(symbol))
		}
	}

	return accum, pos, nil
}
