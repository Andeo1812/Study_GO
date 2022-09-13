package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"uniq/project/uniq/options"
)

func TestDefaultSetup(t *testing.T) {
	var opt options.Options

	var input1 = []string{"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."}

	expected := []string{"I love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik."}
	output := ClassifierHandlers(input1, opt)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestAllSetup(t *testing.T) {
	var opt options.Options

	var input1 = []string{"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."}

	opt.ShowCountStr = true

	expected := []string{"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
		"2 I love music of Kartik."}
	output := ClassifierHandlers(input1, opt)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestUniqSetup(t *testing.T) {
	var opt options.Options

	var input1 = []string{"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."}

	opt.ShowUniqStr = true

	expected := []string{"",
		"Thanks."}
	output := ClassifierHandlers(input1, opt)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestUnUniqSetup(t *testing.T) {
	var opt options.Options

	var input1 = []string{"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."}

	opt.ShowUnUniqStr = true

	expected := []string{"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik."}
	output := ClassifierHandlers(input1, opt)
	assert.Equal(t, expected, output, "they should be equal")
}

func TestDefaultWithoutRegSetup(t *testing.T) {
	var opt options.Options

	var input2 = []string{"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik."}

	opt.RegisterNotImportant = true

	expected := []string{"I LOVE MUSIC.",
		"",
		"I love MuSIC of Kartik.",
		"Thanks.",
		"I love music of kartik."}
	output := ClassifierHandlers(input2, opt)
	assert.Equal(t, expected, output, "they should be equal")
}
