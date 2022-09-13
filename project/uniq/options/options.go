package options

import (
	"errors"
	"flag"
	"fmt"
)

type Options struct {
	ShowCountStr         bool
	ShowDuplicateStr     bool
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
	flag.BoolVar(&o.ShowDuplicateStr, showDuplicateStrFlag, o.ShowDuplicateStr, showDuplicateStrUsage)
	flag.BoolVar(&o.ShowUniqStr, showUniqStrFlag, o.ShowUniqStr, showUniqStrUsage)

	flag.BoolVar(&o.RegisterNotImportant, registerNotImportantFlag, o.RegisterNotImportant, registerNotImportantUsage)

	flag.IntVar(&o.SkipCountWords, skipCountWordsFlag, o.SkipCountWords, skipCountWordsUsage)
	flag.IntVar(&o.SkipCountSymbols, skipCountCharsFlag, o.SkipCountSymbols, skipCountCharsUsage)

	flag.Parse()

	o.InputFile = flag.Arg(0)
	o.OutputFile = flag.Arg(1)

	resCheck := o.checkOptions()
	if resCheck != nil {
		panic(resCheck)
	}

	return o
}

func (o *Options) checkOptions() error {
	var countKeys int

	checkField := func(flag bool, count int) int {
		if flag {
			return count + 1
		}

		return count
	}

	countKeys = checkField(o.ShowCountStr, countKeys)
	countKeys = checkField(o.ShowDuplicateStr, countKeys)
	countKeys = checkField(o.ShowUniqStr, countKeys)

	if countKeys > 1 {
		format := "Only one of the keys: -%s, -%s, -%s is possible\n"

		fmt.Printf(format, showCountStrFlag, showDuplicateStrFlag, showUniqStrFlag)

		return errors.New(format)
	}

	return nil
}
