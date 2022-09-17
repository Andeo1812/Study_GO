package actions

import (
	"Modules/project/calc/parser/configs"
	"fmt"
)

func actionOpenOperand(expression string) (float64, int) {
	fmt.Println("OPERAND_BEGIN", expression)
	var pos int
	if string(expression[pos]) == configs.Lex.OpenParen {
		fmt.Println("OPERAND-START_ADD")
		pos++
		addition, offset := actionSum(0, expression[pos:])
		pos += offset

		fmt.Println("OPERAND-ADDDDDDDDDDDDDDDDDDDDD", addition, pos)

		return addition, pos
	}

	number, lengthNumber := GetNumber(expression[pos:])
	pos += lengthNumber

	fmt.Println("OPERAND-NUM", number, pos)

	return number, pos
}

func actionCloseOperand(accum float64, expression string) (float64, int) {
	if len(expression) == 0 {
		return accum, 0
	}

	var pos int

	symbol := string(expression[pos])
	pos++
	fmt.Println("CLOSEOP", symbol)

	switch symbol {
	case configs.Lex.Plus:
		res, offset := actionSum(accum, expression[pos:])
		pos += offset

		return res, pos
	case configs.Lex.Minus:
		res, offset := actionSub(accum, expression[pos:])
		pos += offset

		return res, pos
	case configs.Lex.Multiply:
		res, offset := actionMul(accum, expression[pos:])
		accum = res
		pos += offset
	case configs.Lex.Divide:
		res, offset := actionDiv(accum, expression[pos:])
		accum /= res
		pos += offset
	case configs.Lex.CloseParen:
		_, offset := actionCloseOperand(accum, expression[pos:])
		pos += offset

		return accum, pos
	}

	return accum, pos
}
