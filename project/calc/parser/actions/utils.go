package actions

import (
	"Modules/project/calc/parser/configs"
	"errors"
	"strconv"
	"strings"
)

func isDigit(value string) bool {
	if strings.Contains(configs.Lex.Digit, value) || value == "." {
		return true
	}

	return false
}

func GetNumber(expression string) (float64, int) {
	//  fmt.Println(expression)
	var iter = 0

	if string(expression[0]) == configs.Lex.Minus {
		iter = 1
	}

	for ; iter < len(expression); iter++ {
		if !isDigit(string(expression[iter])) {
			res, resConversation := strconv.ParseFloat(expression[:iter], 64)
			if resConversation != nil {
				panic(resConversation)
			}

			return res, iter
		}

	}

	if iter == len(expression) {
		res, resConversation := strconv.ParseFloat(expression[:iter], 64)
		if resConversation != nil {
			panic(resConversation)
		}
		return res, iter
	}

	return 0, 0
}

func getEndExpression(expression string) (int, error) {
	for {
		posClose := strings.Index(expression, configs.Lex.CloseParen)
		if posClose == -1 {
			return 0, errors.New("bad expression: no close paren")
		}

		posOpen := strings.Index(expression, configs.Lex.OpenParen)

		if posClose < posOpen || posOpen == -1 {
			return posClose, nil
		}
	}
}
