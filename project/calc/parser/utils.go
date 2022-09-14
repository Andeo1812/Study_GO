package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func isDigit(value string) bool {
	if strings.Contains(lex.digit, value) || value == "." {
		return true
	}

	return false
}

func GetNumber(expression string) (int, float64, error) {
	fmt.Println(expression)
	var iter = 0

	if string(expression[0]) == lex.minus {
		iter = 1
	}

	for ; iter < len(expression); iter++ {
		if !isDigit(string(expression[iter])) {
			res, resConversation := strconv.ParseFloat(expression[:iter], 64)

			return iter, res, resConversation
		}

	}

	if iter == len(expression) {
		res, resConversation := strconv.ParseFloat(expression[:iter], 64)
		return iter, res, resConversation
	}

	return 0, 0, errors.New("error convert")
}
