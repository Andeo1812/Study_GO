package calc

import (
	"errors"
	"strconv"
	"strings"
)

func getDigit(expression string) (int, error) {
	for ind, val := range expression {
		if !strings.Contains(lex.digit, string(val)) {
			return strconv.Atoi(expression[:ind])
		}
	}

	return 0, errors.New("error convert")
}

func Work() {
	expression := input()

	res, _ := getDigit(expression)

	//  fmt.Println(res)

	output(strconv.Itoa(res))
}
