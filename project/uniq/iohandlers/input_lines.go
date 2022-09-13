package iohandlers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

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

func scanLines(flow *os.File, skipWordsCount int, skipSymbolsCount int) (res []string) {
	scanner := bufio.NewScanner(flow)

	for scanner.Scan() {
		line, errSkipWords := skipWords(scanner.Text(), skipWordsCount)
		line, errSkipSymbols := skipWords(line, skipWordsCount)

		if errSkipWords == nil && errSkipSymbols == nil {
			res = append(res, line)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return res
}

func GetLines(path string, skipWords int, skipSymbols int) (res []string) {
	var stream *os.File = os.Stdin

	if path != "" {
		stream, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		defer stream.Close()

		return scanLines(stream, skipWords, skipSymbols)
	}

	return scanLines(stream, skipWords, skipSymbols)
}
