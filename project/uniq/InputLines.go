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

func skipInput(line string, skipWords int, skipSymbols int) (string, int) {
	for i := 0; i < skipWords; i++ {
		indexSpace := strings.Index(line, " ")
		if indexSpace == -1 {
			return "", noResultCode
		}

		curLength := len(line)

		if indexSpace+1 <= curLength-indexSpace+1 {
			line = line[indexSpace+1 : curLength-indexSpace+1]
		} else {
			return "", noResultCode
		}
	}

	if skipSymbols < len(line) {
		line = line[skipSymbols : len(line)-skipWords]
	} else {
		return "", noResultCode
	}

	return line, Success
}

func scanLines(flow *os.File, skipWords int, skipSymbols int) (res *[]string) {
	scanner := bufio.NewScanner(flow)

	var buf []string

	for scanner.Scan() {
		line, res := skipInput(scanner.Text(), skipWords, skipSymbols)
		if res == Success {
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
