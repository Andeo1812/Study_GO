package uniq

func Work() {
	options := getOptions()

	lines := getLines(options.inputFile, options.skipCountWords, options.skipCountSymbols)

	showLines(lines, options)
}
