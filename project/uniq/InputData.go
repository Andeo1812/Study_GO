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

	defer stream.Close()

	getLines(stream)
}

func getLines(flow *os.File) {
	scanner := bufio.NewScanner(flow)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func getData(inputFile string) {
	if inputFile != "" {
		inputData(inputFile)
	} else {
		getLines(os.Stdin)
	}
}
