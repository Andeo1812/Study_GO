package parser

import "fmt"

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
		case lex.plus:
			accum /= number

			res, offset := actionSum(accum, expression[pos:])

			pos += offset

			return res, pos
		case lex.minus:
			accum /= number

			res, offset := actionSub(accum, expression[pos:])

			pos += offset

			return res, pos
		case lex.multiply:
			res, offset := actionMul(accum/number, expression[pos:])

			accum = res

			pos += offset
		case lex.divide:
			accum /= number
		default:
			return accum / number, pos
		}
	}

	return accum, pos
}

func actionMul(accum float64, expression string) (float64, int) {
	fmt.Println("Mul:", accum, expression)
	var pos int = 0

	for pos < len(expression) {
		number, lengthNumber := GetNumber(expression[pos:])
		pos += lengthNumber

		if len(expression) == pos {
			return accum * number, pos
		}

		symbol := string(expression[pos])
		pos++

		switch symbol {
		case lex.plus:
			accum *= number

			res, offset := actionSum(accum, expression[pos:])

			pos += offset

			return res, pos
		case lex.minus:
			accum *= number

			res, offset := actionSub(accum, expression[pos:])

			pos += offset

			return res, pos
		case lex.multiply:
			accum *= number
		case lex.divide:
			res, offset := actionDiv(accum*number, expression[pos:])

			accum = res

			pos += offset
		default:
			return accum * number, pos
		}
	}

	return accum, pos
}

func actionSum(accum float64, expression string) (float64, int) {
	fmt.Println("Sum:", accum, expression)
	var pos int = 0

	for pos < len(expression) {
		number, lengthNumber := GetNumber(expression[pos:])
		pos += lengthNumber

		if len(expression) == pos {
			return accum + number, pos
		}

		symbol := string(expression[pos])
		pos++

		switch symbol {
		case lex.plus:
			accum += number
		case lex.minus:
			res, offset := actionSub(number, expression[pos:])

			accum += res

			pos += offset

		case lex.multiply:
			res, offset := actionMul(number, expression[pos:])

			accum += res

			pos += offset
		case lex.divide:
			res, offset := actionDiv(number, expression[pos:])

			accum += res

			pos += offset
		}
	}

	return accum, pos
}

func actionSub(accum float64, expression string) (float64, int) {
	fmt.Println("Sub", accum, expression)
	var pos int = 0

	for pos < len(expression) {
		number, lengthNumber := GetNumber(expression[pos:])
		pos += lengthNumber

		if len(expression) == pos {
			return accum - number, pos
		}

		symbol := string(expression[pos])
		pos++

		switch symbol {
		case lex.plus:
			res, offset := actionSum(-number, expression[pos:])

			accum += res

			pos += offset
		case lex.minus:
			accum -= number
		case lex.multiply:
			res, offset := actionMul(-number, expression[pos:])

			accum += res

			pos += offset
		case lex.divide:
			res, offset := actionDiv(-number, expression[pos:])

			accum += res

			pos += offset
		}
	}

	return accum, pos
}

func InitCalculate(expression string) float64 {
	res, _ := actionSum(0, expression)
	return res
}
