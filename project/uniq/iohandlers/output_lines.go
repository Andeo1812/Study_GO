package iohandlers

import (
	"os"
	"uniq/project/uniq/handlers"
	"uniq/project/uniq/options"
)

func writeFlow(stream *os.File, lines []string) {
	for _, val := range lines {
		stream.WriteString(val)
	}
}

func ShowLines(lines []string, opt options.Options) error {
	var stream *os.File = os.Stdout

	if opt.OutputFile != "" {
		stream, errCreate := os.Create(opt.OutputFile)
		if errCreate != nil {
			panic(errCreate)
		}

		defer stream.Close()

		result, errHandler := handlers.ClassifierHandlers(lines, opt)
		writeFlow(stream, result)

		return errHandler
	}

	result, errHandler := handlers.ClassifierHandlers(lines, opt)
	writeFlow(stream, result)

	return errHandler
}
