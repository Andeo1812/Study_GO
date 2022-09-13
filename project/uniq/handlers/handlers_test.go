package handlers

import (
	"github.com/stretchr/testify/require"
	"testing"
	"uniq/project/uniq/options"
)

func TestDefaultSetup(t *testing.T) {
	var opt options.Options

	var initData = []string{"I love music.",
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
	output, errHandler := ClassifierHandlers(initData, opt)
	require.Nil(t, errHandler)
	require.Equal(t, expected, output, "they should be equal")
}

func TestAllSetup(t *testing.T) {
	var opt options.Options
	opt.ShowCountStr = true

	var initData = []string{"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."}

	expected := []string{"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
		"2 I love music of Kartik."}
	output, errHandler := ClassifierHandlers(initData, opt)
	require.Nil(t, errHandler)
	require.Equal(t, expected, output, "they should be equal")
}

func TestUniqSetup(t *testing.T) {
	var opt options.Options
	opt.ShowUniqStr = true

	var initData = []string{"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."}

	expected := []string{"",
		"Thanks."}
	output, errHandler := ClassifierHandlers(initData, opt)
	require.Nil(t, errHandler)
	require.Equal(t, expected, output, "they should be equal")
}

func TestUnUniqSetup(t *testing.T) {
	var opt options.Options
	opt.ShowDuplicateStr = true

	var initData = []string{"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"I love music of Kartik.",
		"I love music of Kartik."}

	expected := []string{"I love music.",
		"I love music of Kartik.",
		"I love music of Kartik."}
	output, errHandler := ClassifierHandlers(initData, opt)
	require.Nil(t, errHandler)
	require.Equal(t, expected, output, "they should be equal")
}

func TestDefaultWithoutRegSetup(t *testing.T) {
	var opt options.Options
	opt.RegisterNotImportant = true

	var initData = []string{"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik."}

	expected := []string{"I LOVE MUSIC.",
		"",
		"I love MuSIC of Kartik.",
		"Thanks.",
		"I love music of kartik."}
	output, errHandler := ClassifierHandlers(initData, opt)
	require.Nil(t, errHandler)
	require.Equal(t, expected, output, "they should be equal")
}

func TestCombinationSetup(t *testing.T) {
	var opt options.Options
	opt.RegisterNotImportant = true
	opt.ShowCountStr = true

	var initData = []string{"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
		"I love music of kartik.",
		"I love MuSIC of Kartik."}

	expected := []string{"3 I LOVE MUSIC.",
		"1 ",
		"2 I love MuSIC of Kartik.",
		"1 Thanks.",
		"2 I love music of kartik."}
	output, errHandler := ClassifierHandlers(initData, opt)
	require.Nil(t, errHandler)
	require.Equal(t, expected, output, "they should be equal")

	opt.ShowCountStr = false
	opt.ShowUniqStr = true

	expected = []string{"",
		"Thanks."}
	output, errHandler = ClassifierHandlers(initData, opt)
	require.Nil(t, errHandler)
	require.Equal(t, expected, output, "they should be equal")

	opt.ShowUniqStr = false
	opt.ShowDuplicateStr = true

	expected = []string{"I LOVE MUSIC.",
		"I love MuSIC of Kartik.",
		"I love music of kartik."}
	output, errHandler = ClassifierHandlers(initData, opt)
	require.Nil(t, errHandler)
	require.Equal(t, expected, output, "they should be equal")
}

func TestBadSetup(t *testing.T) {
	var opt options.Options

	var initData = make([]string, 0)

	_, errHandler := ClassifierHandlers(initData, opt)
	require.NotNil(t, errHandler)

	opt.ShowCountStr = true

	_, errHandler = ClassifierHandlers(initData, opt)
	require.NotNil(t, errHandler)

	opt.ShowCountStr = false
	opt.ShowUniqStr = true

	_, errHandler = ClassifierHandlers(initData, opt)
	require.NotNil(t, errHandler)

	opt.ShowUniqStr = false
	opt.ShowDuplicateStr = true

	_, errHandler = ClassifierHandlers(initData, opt)
	require.NotNil(t, errHandler)
}
