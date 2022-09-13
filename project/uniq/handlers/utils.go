package handlers

import (
	"errors"
	"strings"
)

type node struct {
	key   string
	value int
}

func skipWords(line string, countSkip int) (string, error) {
	for i := 0; i < countSkip; i++ {
		indexSpace := strings.Index(line, " ")
		if indexSpace == -1 {
			return "", errors.New("no result")
		}

		if indexSpace+1 >= len(line)-indexSpace {
			return "", errors.New("no result")
		}

		line = line[indexSpace+1:]
	}

	return line, nil
}

func skipSymbols(line string, countSkip int) (string, error) {
	if countSkip <= len(line) {
		return line[countSkip:], nil
	}

	return "", errors.New("no result")
}

func getTypeComparator(RegisterNotImportant bool) func(string, string) int {
	compareRule := func(left string, right string) int {
		if strings.EqualFold(left, right) {
			return 0
		}

		return -1
	}

	if !RegisterNotImportant {
		return strings.Compare
	}

	return compareRule
}
