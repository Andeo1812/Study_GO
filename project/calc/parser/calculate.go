package parser

import "fmt"

type node struct {
	number float64
	length int
}

func action(accum float64, expression string, operation string) (float64, int) {
	lengthNumber, number, resGet := GetNumber(expression)
	if resGet != nil {
		panic(resGet)
	}

}

func Calculate(accum float64, expression string, operation string) float64 {
	fmt.Println(accum)

	if len(expression) == 0 {
		return accum
	}

	symbol := string(expression[0])

	//  fmt.Println(symbol)

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
	case symbol == lex.multiply:
		fmt.Println("multiply", operation)
		return accum * Calculate(accum, expression[1:], symbol)
	case symbol == lex.divide:
		fmt.Println("divide", operation)
		return accum / accum / Calculate(accum, expression[1:], symbol)
	default:
		panic("No such symbol")
	}
}

func InitCalculate(expression string) float64 {
	return Calculate(0, expression, "")
}
