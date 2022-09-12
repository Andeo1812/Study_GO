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

	lines := getLines(options.inputFile)

	//  fmt.Println(lines)

	linesTable := analysisData(lines)

	showLines(options.outputFile, simple, &linesTable)
}
