package uniq

import (
	"bufio"
	"fmt"
	"os"
)

func inputData(path string) {
	stream, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner(stream)
}

func scanner(flow *os.File) {
	scanner := bufio.NewScanner(flow)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func Handling() {
	options := getFlags()

	if options.inputFile != "" {
		inputData(options.inputFile)
	} else {
		scanner(os.Stdin)
	}
}
