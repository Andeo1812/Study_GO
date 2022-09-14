package parser

import (
	"errors"
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
	var iterator = 0
	for _, val := range expression {
		if !isDigit(string(val)) {
			res, resConversation := strconv.ParseFloat(expression[:iterator], 64)
			return iterator, res, resConversation
		}

		iterator++
	}

	if iterator == len(expression) {
		res, resConversation := strconv.ParseFloat(expression[:iterator], 64)
		return iterator, res, resConversation
	}

	return 0, 0, errors.New("error convert")
}
