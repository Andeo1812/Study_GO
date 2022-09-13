package uniq

import "os"

func showLines(lines []string, opt options) {
	var stream *os.File = os.Stdout

	if opt.outputFile != "" {
		stream, errCreate := os.Create(opt.outputFile)
		if errCreate != nil {
			panic(errCreate)
		}

		defer stream.Close()

		result := classifierHandlers(lines, opt)
		writeFlow(stream, result)

		return
	}

	result := classifierHandlers(lines, opt)
	writeFlow(stream, result)
}

func writeFlow(stream *os.File, lines []string) {
	for _, val := range lines {
		stream.WriteString(val)
	}
}

func classifierHandlers(lines []string, opt options) (res []string) {
	switch {
	case opt.showCountStr:
		return showAll(lines, opt)
	case opt.showUniqStr:
		return showUniq(lines, opt)
	case opt.showUnUniqStr:
		return showUnUniq(lines, opt)
	default:
		return showDefault(lines, opt)
	}
}
