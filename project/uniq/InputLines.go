package uniq

import (
	"bufio"
	"fmt"
	"os"
)

func scanLines(flow *os.File) (res *[]string) {
	scanner := bufio.NewScanner(flow)

	var buf []string

	for scanner.Scan() {
		buf = append(buf, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return &buf
}

func getLines(path string) (res *[]string) {
	var stream *os.File = os.Stdin

	if path != "" {
		stream, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		defer stream.Close()

		return scanLines(stream)
	}

	return scanLines(stream)
}
