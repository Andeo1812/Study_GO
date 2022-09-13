package uniq

import (
	"Modules/project/uniq/handlers"
	"Modules/project/uniq/iohandlers"
	"Modules/project/uniq/options"
	"os"
)

func showLines(lines []string, opt options.Options) {
	var stream *os.File = os.Stdout

	if opt.OutputFile != "" {
		stream, errCreate := os.Create(opt.OutputFile)
		if errCreate != nil {
			panic(errCreate)
		}

		defer stream.Close()

		result := handlers.ClassifierHandlers(lines, opt)
		iohandlers.WriteFlow(stream, result)

		return
	}

	result := handlers.ClassifierHandlers(lines, opt)
	iohandlers.WriteFlow(stream, result)
}
