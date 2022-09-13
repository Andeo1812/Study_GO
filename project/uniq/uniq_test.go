package uniq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var input1 = []string{"I love music.",
	"I love music.",
	"I love music.",
	"",
	"I love music of Kartik.",
	"I love music of Kartik.",
	"Thanks.",
	"I love music of Kartik.",
	"I love music of Kartik."}

var input2 = []string{"I LOVE MUSIC.",
	"I love music.",
	"I LoVe MuSiC.",
	"",
	"I love MuSIC of Kartik.",
	"I love music of kartik.",
	"Thanks.",
	"I love music of kartik.",
	"I love MuSIC of Kartik."}

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
		"",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik."}
	output := classifierHandlers(input1, opt)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestGlobalAllSetup(t *testing.T) {
	var opt options

	opt.inputFile = "tests/Text.txt"
	opt.outputFile = "tmp.txt"
	opt.showCountStr = true

	expected := []string{"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
		"2 I love music of Kartik."}
	output := classifierHandlers(input1, opt)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestGlobalUniqSetup(t *testing.T) {
	var opt options

	opt.inputFile = "tests/Text.txt"
	opt.outputFile = "tmp.txt"
	opt.showUniqStr = true

	expected := []string{"",
		"Thanks."}
	output := classifierHandlers(input1, opt)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestGlobalUnUniqSetup(t *testing.T) {
	var opt options

	opt.inputFile = "tests/Text.txt"
	opt.outputFile = "tmp.txt"
	opt.showUnUniqStr = true

	expected := []string{"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik."}
	output := classifierHandlers(input1, opt)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestGlobalDefaultWithoutRegSetup(t *testing.T) {
	var opt options

	opt.inputFile = "tests/TextRegistr.txt"
	opt.outputFile = "tmp.txt"
	opt.registerNotImportant = true

	expected := []string{"I LOVE MUSIC.",
		"",
		"I love MuSIC of Kartik.",
		"Thanks.",
		"I love music of kartik."}
	output := classifierHandlers(input2, opt)
	assert.Equal(t, expected, output, "they should be equal")
}
