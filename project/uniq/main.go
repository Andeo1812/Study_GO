package uniq

import (
	"Modules/project/uniq/iohandlers"
	"Modules/project/uniq/options"
)

func Work() {
	options := options.GetOptions()

	lines := iohandlers.GetLines(options.InputFile, options.SkipCountWords, options.SkipCountSymbols)

	showLines(lines, options)
}
