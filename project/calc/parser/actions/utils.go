package actions

import (
	"Calc/project/calc/parser/configs"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func isDigit(value string) bool {
	if strings.Contains(configs.Lex.Digit, value) || value == "." {
		return true
	}

	return false
}

func GetNumber(expression string) (float64, int, error) {
	var iter = 0

	for ; iter < len(expression); iter++ {
		if !isDigit(string(expression[iter])) {
			res, resConversation := strconv.ParseFloat(expression[:iter], 64)
			return res, iter, resConversation
		}

	}

	if iter == len(expression) {
		res, resConversation := strconv.ParseFloat(expression[:iter], 64)
		return res, iter, resConversation
	}

	return 0, 0, errors.New("wrong expression")
}

func GetEndExpression(expression string) (int, error) {
	var countOpen int = 1

	posClose := strings.Index(expression, configs.Lex.CloseParen)
	if posClose == -1 {
		return 0, errors.New("wrong expression")
	}

	posOpen := strings.Index(expression, configs.Lex.OpenParen)
	if posOpen == -1 {
		return posClose, nil
	}

	for ind, val := range expression {
		if string(val) == configs.Lex.OpenParen {
			countOpen++
		}

		if string(val) == configs.Lex.CloseParen {
			countOpen--

			if countOpen == 0 {
				return ind, nil
			}
		}

	}

	return 0, errors.New("wrong expression")
}

func wrongSymbol(symbol string) string {
	massage := "No such symbol in functional: %s\n"

	return fmt.Sprintf(massage, symbol)
}
