package uniq

func analysisData(lines *[]string) map[string]int {
	resultTable := map[string]int{}

	for _, val := range *lines {
		if _, keyExist := resultTable[val]; !keyExist {
			resultTable[val] = 1
		} else {
			resultTable[val] += 1
		}
	}

	return resultTable
}

func Handling() {
	options := getOptions()

	lines := getLines(options.inputFile, options.skipCountWords, options.skipCountSymbols)

	if options.showCountStr {
		showLines(options.outputFile, showAll, lines, options.registerNotImportant)

		return
	}

	if options.showUniqStr {
		showLines(options.outputFile, showUniq, lines, options.registerNotImportant)

		return
	}

	if options.showUnUniqStr {
		showLines(options.outputFile, showUnUniq, lines, options.registerNotImportant)

		return
	}

	showLines(options.outputFile, showDefault, lines, options.registerNotImportant)
}
