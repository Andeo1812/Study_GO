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

		classifierHandlers(stream, lines, opt)

		return
	}

	classifierHandlers(stream, lines, opt)
}

func classifierHandlers(stream *os.File, lines []string, opt options) {
	switch {
	case opt.showCountStr:
		showAll(stream, lines, opt)
	case opt.showUniqStr:
		showUniq(stream, lines, opt)
	case opt.showUnUniqStr:
		showUnUniq(stream, lines, opt)
	default:
		showDefault(stream, lines, opt)
	}
}
