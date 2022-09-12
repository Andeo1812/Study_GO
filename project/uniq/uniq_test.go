package uniq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkipSymbols(t *testing.T) {
	var input string = "1 22 333 4444!"

	// skip 10 symbols
	output := "444!"
	resSkip, res := skipSymbols(input, 10)
	assert.Equal(t, nil, res, "they should be equal")
	assert.Equal(t, output, resSkip, "they should be equal")
}

func TestSkipWords(t *testing.T) {
	var input string = "1 22 333 4444!"

	// skip 1 word
	var output string = "22 333 4444!"
	resSkip, res := skipWords(input, 1)

	assert.Equal(t, nil, res, "they should be equal")
	assert.Equal(t, output, resSkip, "they should be equal")
}

func TestSkipCombination(t *testing.T) {
	var input string = "1 22 333 4444!"

	// combination skip 1 word skip 3 symbols
	output := "333 4444!"
	resSkip, res1 := skipWords(input, 1)
	resSkip, res2 := skipSymbols(resSkip, 3)
	assert.Equal(t, nil, res1, "they should be equal")
	assert.Equal(t, nil, res2, "they should be equal")
	assert.Equal(t, output, resSkip, "they should be equal")

	// combination skip 2 word skip 1 symbols
	output = "33 4444!"
	resSkip, res1 = skipWords(input, 2)
	resSkip, res2 = skipSymbols(resSkip, 1)
	assert.Equal(t, nil, res1, "they should be equal")
	assert.Equal(t, nil, res2, "they should be equal")
	assert.Equal(t, output, resSkip, "they should be equal")
}

func TestGlobalDefaultSetup(t *testing.T) {
	var opt options

	opt.inputFile = "tests/Text.txt"
	opt.outputFile = "tmp.txt"

	expected := []string{"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."}
	output := getLines(opt.inputFile, opt.skipCountWords, opt.skipCountSymbols)
	assert.Equal(t, expected, output, "they should be equal")

	expected = []string{"I love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik."}
	showLines(output, opt)
	output = getLines(opt.outputFile, opt.skipCountWords, opt.skipCountSymbols)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestGlobalAllSetup(t *testing.T) {
	var opt options

	opt.inputFile = "tests/Text.txt"
	opt.outputFile = "tmp.txt"
	opt.showCountStr = true

	output := getLines(opt.inputFile, opt.skipCountWords, opt.skipCountSymbols)

	expected := []string{"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
		"2 I love music of Kartik."}
	showLines(output, opt)
	output = getLines(opt.outputFile, opt.skipCountWords, opt.skipCountSymbols)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestGlobalUniqSetup(t *testing.T) {
	var opt options

	opt.inputFile = "tests/Text.txt"
	opt.outputFile = "tmp.txt"
	opt.showUniqStr = true

	output := getLines(opt.inputFile, opt.skipCountWords, opt.skipCountSymbols)

	expected := []string{"",
		"Thanks."}
	showLines(output, opt)
	output = getLines(opt.outputFile, opt.skipCountWords, opt.skipCountSymbols)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestGlobalUnUniqSetup(t *testing.T) {
	var opt options

	opt.inputFile = "tests/Text.txt"
	opt.outputFile = "tmp.txt"
	opt.showUnUniqStr = true

	output := getLines(opt.inputFile, opt.skipCountWords, opt.skipCountSymbols)

	expected := []string{"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik."}
	showLines(output, opt)
	output = getLines(opt.outputFile, opt.skipCountWords, opt.skipCountSymbols)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestGlobalDefaultWithoutRegSetup(t *testing.T) {
	var opt options

	opt.inputFile = "tests/TextRegistr.txt"
	opt.outputFile = "tmp.txt"
	opt.registerNotImportant = true

	output := getLines(opt.inputFile, opt.skipCountWords, opt.skipCountSymbols)

	expected := []string{"I LOVE MUSIC.",
		"",
		"I love MuSIC of Kartik.",
		"Thanks.",
		"I love music of kartik."}
	showLines(output, opt)
	output = getLines(opt.outputFile, opt.skipCountWords, opt.skipCountSymbols)
	assert.Equal(t, expected, output, "they should be equal")
}
