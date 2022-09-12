package uniq

import "os"

func showLines(path string, outputRule outputRuleType, lines *[]string) {
	var stream *os.File = os.Stdout

	if path != "" {
		stream, errCreate := os.Create(path)
		if errCreate != nil {
			panic(errCreate)
		}

		defer stream.Close()

		outputRule(stream, lines)

		return
	}

	outputRule(stream, lines)
}
