package actions

import (
	"Modules/project/calc/parser/configs"
	"errors"
)

func actionDiv(accum float64, expression string) (float64, int, error) {
	var pos int = 0

	for pos < len(expression) {
		addition, offset, errGetOperand := actionOpenOperand(expression[pos:])
		if errGetOperand != nil {
			return 0, 0, errGetOperand
		}

		if addition == 0 {
			return 0, 0, errors.New("divide on zero")
		}

		pos += offset

		if len(expression) == pos {
			return accum / addition, pos, nil
		}

		symbol := string(expression[pos])
		pos++

		switch symbol {
		case configs.Lex.Plus:
			res, offset, errSum := actionSum(accum/addition, expression[pos:])
			if errSum != nil {
				return 0, 0, errSum
			}

			pos += offset

			return res, pos, nil
		case configs.Lex.Minus:
			res, offset, errSub := actionSub(accum/addition, expression[pos:])
			if errSub != nil {
				return 0, 0, errSub
			}

			pos += offset

			return res, pos, nil
		case configs.Lex.Multiply:
			res, offset, errMul := actionMul(accum/addition, expression[pos:])
			if errMul != nil {
				return 0, 0, errMul
			}

			accum = res
			pos += offset
		case configs.Lex.Divide:
			accum /= addition
		default:
			return 0, 0, errors.New(wrongSymbol(symbol))
		}
	}

	return accum, pos, nil
}
