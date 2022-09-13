package handlers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSkipSymbols(t *testing.T) {
	var input string = "1 22 333 4444!"

	// skip 10 symbols
	output := "444!"
	resSkip, res := skipSymbols(input, 10)
	require.Equal(t, nil, res, "they should be equal")
	require.Equal(t, output, resSkip, "they should be equal")

	// skip many symbols
	output = ""
	resSkip, res = skipSymbols(input, 100)
	require.NotNil(t, res)
	require.Equal(t, output, resSkip, "they should be equal")
}

func TestSkipWords(t *testing.T) {
	var input string = "1 22 333 4444!"

	// skip 1 word
	var output string = "22 333 4444!"
	resSkip, res := skipWords(input, 1)
	require.Equal(t, nil, res, "they should be equal")
	require.Equal(t, output, resSkip, "they should be equal")

	// skip many words
	output = ""
	resSkip, res = skipWords(input, 100)
	require.NotNil(t, res)
	require.Equal(t, output, resSkip, "they should be equal")

	// skip words in border
	output = ""
	input = "1 22 333 4444! "
	resSkip, res = skipWords(input, 4)
	require.NotNil(t, res)
	require.Equal(t, output, resSkip, "they should be equal")
}

func TestSkipCombination(t *testing.T) {
	var input string = "1 22 333 4444!"

	// combination skip 1 word skip 3 symbols
	output := "333 4444!"
	resSkip, res1 := skipWords(input, 1)
	resSkip, res2 := skipSymbols(resSkip, 3)
	require.Equal(t, nil, res1, "they should be equal")
	require.Equal(t, nil, res2, "they should be equal")
	require.Equal(t, output, resSkip, "they should be equal")

	// combination skip 2 word skip 1 symbols
	output = "33 4444!"
	resSkip, res1 = skipWords(input, 2)
	resSkip, res2 = skipSymbols(resSkip, 1)
	require.Equal(t, nil, res1, "they should be equal")
	require.Equal(t, nil, res2, "they should be equal")
	require.Equal(t, output, resSkip, "they should be equal")
}
