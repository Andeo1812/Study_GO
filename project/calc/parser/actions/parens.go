package actions

import (
	"Modules/project/calc/parser/configs"
)

func actionOpenOperand(expression string) (float64, int, error) {
	var pos int
	if string(expression[pos]) == configs.Lex.OpenParen {
		pos++
		posEnd, errGetEnd := GetEndExpression(expression[pos:])
		if errGetEnd != nil {
			return 0, 0, errGetEnd
		}

		addition, offset, errResExpression := actionSum(0, expression[pos:posEnd+1])
		if errResExpression != nil {
			return 0, 0, errResExpression
		}

		pos += offset
		pos++

		return addition, pos, nil
	}

	number, lengthNumber, errGet := GetNumber(expression[pos:])
	if errGet != nil {
		return 0, 0, errGet
	}

	pos += lengthNumber

	return number, pos, nil
}
