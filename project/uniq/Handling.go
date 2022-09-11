package uniq

import (
	"os"
)

func getData(inputFile string) {
	if inputFile != "" {
		inputData(inputFile)
	} else {
		scanner(os.Stdin)
	}
}

func Handling() {
	options := getFlags()

	getData(options.inputFile)
}
