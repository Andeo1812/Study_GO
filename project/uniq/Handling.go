package uniq

import "fmt"

type setup struct {
	options options
	data    []string
}

func createSetup(opt *options, data *[]string) setup {
	var s setup

	s.options = *opt
	s.data = *data

	return s
}

func Handling() {
	options := getOptions()

	lines := getLines(options.inputFile)

	var setup setup = createSetup(&options, lines)

	fmt.Println(setup.data)

}
