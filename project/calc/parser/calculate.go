package parser

import "fmt"

type node struct {
	number float64
	length int
}

func action(accum float64, expression string, operation string) (float64, int) {
	fmt.Println("Init", accum, expression)

	if len(expression) == 0 {
		return accum, 0
	}

	lengthNumber, number, resGet := GetNumber(expression)
	if resGet != nil {
		panic(resGet)
	}

	switch {
	case operation == lex.multiply:
		accum *= number
	case operation == lex.divide:
		accum /= number
	}

	if len(expression[lengthNumber:]) == 0 {
		return accum, 0
	}

	symbol := string(expression[lengthNumber])

	fmt.Println("res", accum)

	fmt.Println(symbol)

	switch {
	case symbol == lex.multiply || symbol == lex.divide:
		return action(accum, expression[lengthNumber+1:], symbol)
	default:
		return accum, lengthNumber
	}
}

func Calculate(accum float64, expression string, operation string) float64 {
	fmt.Println(accum)

	if len(expression) == 0 {
		return accum
	}

	symbol := string(expression[0])

	fmt.Println(symbol)

	switch {
	case isDigit(symbol):
		lengthNumber, number, resGet := GetNumber(expression)
		if resGet != nil {
			panic(resGet)
		}

		fmt.Println("params number", number, lengthNumber, expression)

		if operation == lex.minus {
			return Calculate(-number, expression[lengthNumber:], symbol)
		}

		return Calculate(number, expression[lengthNumber:], symbol)
	case symbol == lex.plus || symbol == lex.minus:
		fmt.Println("plus-minus", operation)
		return accum + Calculate(accum, expression[1:], symbol)
	case symbol == lex.multiply || symbol == lex.divide:
		fmt.Println("M_D", operation)
		res, offset := action(accum, expression[1:], symbol)

		fmt.Println("M_D", operation)

		if len(expression) == offset {
			return accum + res
		}

		return accum + res + Calculate(accum, expression[offset:], symbol)
	default:
		panic("No such symbol")
	}
}

func InitCalculate(expression string) float64 {
	return Calculate(0, expression, "")
}

//  	case symbol == lex.multiply:
//			return accum * Calculate(accum, expression[1:], symbol)
//		case symbol == lex.divide:
//			return accum / accum / Calculate(accum, expression[1:], symbol)
