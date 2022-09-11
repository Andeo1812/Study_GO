package uniq

import (
	"flag"
	"fmt"
	"os"
)

type options struct {
	showCountStr         bool
	showUnUniqStr        bool
	showUniqStr          bool
	registerNotImportant bool
	skipCountWords       uint
	skipCountChars       uint
	inputFile            string
	outputFile           string
}

func getOptions() options {
	var o options

	flag.BoolVar(&o.showCountStr, showCountStrFlag, o.showCountStr, showCountStrUsage)
	flag.BoolVar(&o.showUnUniqStr, showUnUniqStrFlag, o.showUnUniqStr, showUnUniqStrUsage)
	flag.BoolVar(&o.showUniqStr, showUniqStrFlag, o.showUniqStr, showUniqStrUsage)

	flag.BoolVar(&o.registerNotImportant, registerNotImportantFlag, o.registerNotImportant, registerNotImportantUsage)

	flag.UintVar(&o.skipCountWords, skipCountWordsFlag, o.skipCountWords, skipCountWordsUsage)
	flag.UintVar(&o.skipCountChars, skipCountCharsFlag, o.skipCountChars, skipCountCharsUsage)

	flag.Parse()

	o.inputFile = flag.Arg(0)
	o.outputFile = flag.Arg(1)

	o.checkOptions()

	return o
}

func (o *options) checkOptions() {
	var countKeys int

	if o.showCountStr {
		countKeys++
	}

	if o.showUnUniqStr {
		countKeys++
	}

	if o.showUniqStr {
		countKeys++
	}

	if countKeys > 1 {
		format := "Only one of the keys: -%s, -%s, -%s is possible\n"

		fmt.Printf(format, showCountStrFlag, showUnUniqStrFlag, showUniqStrFlag)

		os.Exit(wrongCombinationFlags)
	}
}
