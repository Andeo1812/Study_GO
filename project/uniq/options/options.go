package options

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	ShowCountStr         bool
	ShowUnUniqStr        bool
	ShowUniqStr          bool
	RegisterNotImportant bool
	SkipCountWords       int
	SkipCountSymbols     int
	InputFile            string
	OutputFile           string
}

func GetOptions() Options {
	var o Options

	flag.BoolVar(&o.ShowCountStr, showCountStrFlag, o.ShowCountStr, showCountStrUsage)
	flag.BoolVar(&o.ShowUnUniqStr, showUnUniqStrFlag, o.ShowUnUniqStr, showUnUniqStrUsage)
	flag.BoolVar(&o.ShowUniqStr, showUniqStrFlag, o.ShowUniqStr, showUniqStrUsage)

	flag.BoolVar(&o.RegisterNotImportant, registerNotImportantFlag, o.RegisterNotImportant, registerNotImportantUsage)

	flag.IntVar(&o.SkipCountWords, skipCountWordsFlag, o.SkipCountWords, skipCountWordsUsage)
	flag.IntVar(&o.SkipCountSymbols, skipCountCharsFlag, o.SkipCountSymbols, skipCountCharsUsage)

	flag.Parse()

	o.InputFile = flag.Arg(0)
	o.OutputFile = flag.Arg(1)

	o.checkOptions()

	return o
}

func (o *Options) checkOptions() {
	var countKeys int

	checkField := func(flag bool, count int) int {
		if flag {
			return count + 1
		}

		return count
	}

	countKeys = checkField(o.ShowCountStr, countKeys)
	countKeys = checkField(o.ShowUnUniqStr, countKeys)
	countKeys = checkField(o.ShowUniqStr, countKeys)

	if countKeys > 1 {
		format := "Only one of the keys: -%s, -%s, -%s is possible\n"

		fmt.Printf(format, showCountStrFlag, showUnUniqStrFlag, showUniqStrFlag)

		os.Exit(4)
	}
}
