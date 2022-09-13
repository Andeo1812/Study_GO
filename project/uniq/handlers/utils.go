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

		curLength := len(line)

		if indexSpace+1 < curLength-indexSpace {
			line = line[indexSpace+1:]
		} else {
			return "", errors.New("no result")
		}
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
	compareRule := strings.Compare

	if RegisterNotImportant {
		compareRule = func(left string, right string) int {
			if strings.EqualFold(left, right) {
				return 0
			}

			return -1
		}
	}

	return compareRule
}
