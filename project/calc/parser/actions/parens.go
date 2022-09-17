package actions

import (
	"Modules/project/calc/parser/configs"
)

func actionOpenOperand(expression string) (float64, int, error) {
	var pos int
	if string(expression[pos]) == configs.Lex.OpenParen {
		pos++
		posEnd := GetEndExpression(expression[pos:])

		addition, offset, err := actionSum(0, expression[pos:posEnd+1])
		pos += offset
		pos++

		return addition, pos, err
	}

	number, lengthNumber, err := GetNumber(expression[pos:])
	pos += lengthNumber

	return number, pos, err
}
