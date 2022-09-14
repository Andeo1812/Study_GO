package parser

import (
	"errors"
	"strconv"
	"strings"
)

func isDigit(digits string, value string) bool {
	if strings.Contains(digits, value) || value == "." {
		return true
	}

	return false
}

func GetNumber(expression string) (float64, error) {
	var iterator = 0
	for _, val := range expression {
		if !isDigit(lex.digit, string(val)) {
			return strconv.ParseFloat(expression[:iterator], 64)
		}

		iterator++
	}

	if iterator == len(expression) {
		return strconv.ParseFloat(expression, 64)
	}

	return 0, errors.New("error convert")
}
