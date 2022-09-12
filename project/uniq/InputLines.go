package uniq

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Success      = 1
	noResultCode = -3 //  Code : No result
)

func skipWords(line string, countSkip int) (string, int) {
	for i := 0; i < countSkip; i++ {
		indexSpace := strings.Index(line, " ")
		if indexSpace == -1 {
			return "", noResultCode
		}

		curLength := len(line)

		if indexSpace+1 < curLength-indexSpace {
			line = line[indexSpace+1:]
		} else {
			return "", noResultCode
		}
	}

	return line, Success
}

func skipSymbols(line string, countSkip int) (string, int) {
	if countSkip <= len(line) {
		line = line[countSkip:]
	} else {
		return "", noResultCode
	}

	return line, Success
}

func scanLines(flow *os.File, skipWordsCount int, skipSymbolsCount int) (res *[]string) {
	scanner := bufio.NewScanner(flow)

	var buf []string

	for scanner.Scan() {
		line, resSkipWords := skipWords(scanner.Text(), skipWordsCount)
		line, resSkipSumbols := skipWords(line, skipWordsCount)

		if resSkipWords == Success && resSkipSumbols == Success {
			buf = append(buf, line)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return &buf
}

func getLines(path string, skipWords int, skipSymbols int) (res *[]string) {
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
