package uniq

import (
	"flag"
	"fmt"
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

func getFlags() options {
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

	return o
}

func Init() {
	options := getFlags()

	fmt.Println(options.showCountStr)

	fmt.Println(options.showUnUniqStr)

	fmt.Println(options.showUniqStr)

	fmt.Println(options.skipCountWords)

	fmt.Println(options.skipCountChars)

	fmt.Println(options.inputFile)

	fmt.Println(options.outputFile)
}
