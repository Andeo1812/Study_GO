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

	linesTable := analysisData(lines)

	if options.showCountStr {
		showLines(options.outputFile, showAll, &linesTable)

		return
	}

	if options.showUniqStr {
		showLines(options.outputFile, showUniq, &linesTable)

		return
	}

	if options.showUnUniqStr {
		showLines(options.outputFile, showUnUniq, &linesTable)

		return
	}
}
