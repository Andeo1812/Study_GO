package main

import (
	"uniq/project/uniq/iohandlers"
	"uniq/project/uniq/options"
)

func main() {
	options := options.GetOptions()

	lines := iohandlers.GetLines(options.InputFile, options.SkipCountWords, options.SkipCountSymbols)

	resShow := iohandlers.ShowLines(lines, options)
	if resShow != nil {
		panic("Bad data")
	}
}
