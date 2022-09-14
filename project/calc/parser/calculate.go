package parser

type node struct {
	number float64
	length int
}

func Calculate(expression string) float64 {
	var accum float64 = 0

	var i int = 0

	for i < len(expression) {
		symbol := string(expression[i])

		//  fmt.Println(symbol)

		switch {
		case isDigit(symbol) || symbol == lex.minus:
			lengthNumber, number, resGet := GetNumber(expression[i:])
			if resGet != nil {
				panic(resGet)
			}

			accum += number

			i += lengthNumber

			//  fmt.Println("params number", number, lengthNumber, expression)
		case symbol == lex.plus:
			accum += Calculate(expression[i+1:])

			return accum
		case symbol == lex.multiply:
			accum *= Calculate(expression[i+1:])

			return accum
		case symbol == lex.divide:
			accum /= Calculate(expression[i+1:])

			return accum
		}
	}

	return accum
}
