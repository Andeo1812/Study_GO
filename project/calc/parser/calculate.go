package parser

type node struct {
	number float64
	length int
}

func Calculate(expression string) float64 {
	var accum float64 = 0

	for i := 0; i < len(expression); {
		symbol := string(expression[i])

		switch {
		case isDigit(symbol):
			lengthNumber, number, resGet := GetNumber(expression)
			if resGet != nil {
				panic("Bad input - error conversation")
			}

			accum = number

			i += lengthNumber
		case symbol == "+":
			accum += Calculate(expression[i+1:])

			return accum
		case symbol == "*":
			accum *= Calculate(expression[i+1:])

			return accum
		}
	}

	return accum
}
