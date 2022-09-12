package uniq

func Handling() {
	options := getOptions()

	lines := getLines(options.inputFile, options.skipCountWords, options.skipCountSymbols)

	showLines(lines, options)
}
