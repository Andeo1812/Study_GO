package uniq

import (
	"bufio"
	"fmt"
	"os"
)

func inputData(path string) (res []string) {
	stream, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer stream.Close()

	return getLines(stream)
}

func getLines(flow *os.File) (res []string) {
	scanner := bufio.NewScanner(flow)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return res
}

func getData(inputFile string) (res []string) {
	if inputFile != "" {
		res = inputData(inputFile)
		return res
	}

	res = getLines(os.Stdin)

	return res
}
