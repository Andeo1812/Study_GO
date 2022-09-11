package uniq

import (
	"bufio"
	"fmt"
	"os"
)

func inputLines(path string) (res *[]string) {
	stream, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer stream.Close()

	res = scanLines(stream)

	return res
}

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

func getLines(inputFile string) (res *[]string) {
	if inputFile != "" {
		res = inputLines(inputFile)
		return res
	}

	res = scanLines(os.Stdin)

	return res
}
