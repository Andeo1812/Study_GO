package parser

import "fmt"

func actionDiv(accum float64, expression string) float64 {
	fmt.Println("Div:", accum, expression)
	var i int = 0

	for i < len(expression) {
		number, lengthNumber := GetNumber(expression[i:])
		i += lengthNumber

		if len(expression) == i {
			return accum / number
		}

		symbol := string(expression[i])
		i++

		switch symbol {
		case lex.plus:
			return actionSum(number+accum, expression[i:])
		case lex.minus:
			return actionSub(number+accum, expression[i:])
		case lex.multiply:
			return actionMul(number*accum, expression[i:])
		case lex.divide:
			accum /= number
		}
	}

	return accum
}

func actionMul(accum float64, expression string) float64 {
	fmt.Println("Mul:", accum, expression)
	var i int = 0

	for i < len(expression) {
		number, lengthNumber := GetNumber(expression[i:])
		i += lengthNumber

		if len(expression) == i {
			return accum * number
		}

		symbol := string(expression[i])
		i++

		switch symbol {
		case lex.plus:
			return actionSum(number+accum, expression[i:])
		case lex.minus:
			return actionSub(number+accum, expression[i:])
		case lex.multiply:
			accum *= number
		case lex.divide:
			return accum + actionDiv(number, expression[i:])
		}
	}

	return accum
}

func actionSum(accum float64, expression string) float64 {
	fmt.Println("Sum:", accum, expression)
	var i int = 0

	for i < len(expression) {
		number, lengthNumber := GetNumber(expression[i:])
		i += lengthNumber

		if len(expression) == i {
			return accum + number
		}

		symbol := string(expression[i])
		i++

		switch symbol {
		case lex.plus:
			accum += number
		case lex.minus:
			return actionSub(number+accum, expression[i:])
		case lex.multiply:
			return accum + actionMul(number, expression[i:])
		case lex.divide:
			return accum + actionDiv(number, expression[i:])
		}
	}

	return accum
}

func actionSub(accum float64, expression string) float64 {
	fmt.Println("SubAction:", accum, expression)
	var i int = 0

	for i < len(expression) {
		number, lengthNumber := GetNumber(expression[i:])
		i += lengthNumber

		if len(expression) == i {
			return accum - number
		}

		symbol := string(expression[i])
		i++

		switch symbol {
		case lex.plus:
			return actionSum(-number+accum, expression[i:])
		case lex.minus:
			accum -= number
		case lex.multiply:
			return actionMul(-number*accum, expression[i:])
		case lex.divide:
			return actionDiv(-number*accum, expression[i:])
		}
	}

	return accum
}

func InitCalculate(expression string) float64 {
	return actionSum(0, expression)
}
