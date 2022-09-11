package uniq

import "fmt"

func Handling() {
	options := getOptions()

	lines := getData(options.inputFile)

	fmt.Println(lines)

}
