package uniq

import "os"

func showLines(path string, outputRule outputRuleType, lines *[]string, registerNotImportant bool) {
	var stream *os.File = os.Stdout

	if path != "" {
		stream, errCreate := os.Create(path)
		if errCreate != nil {
			panic(errCreate)
		}

		defer stream.Close()

		outputRule(stream, lines, registerNotImportant)

		return
	}

	outputRule(stream, lines, registerNotImportant)
}
