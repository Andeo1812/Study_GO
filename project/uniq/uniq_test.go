package uniq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkipInput(t *testing.T) {
	var input string = "1 22 333 4444!"

	// skip 1 word
	var output string = "22 333 4444!"
	resSkip, res := skipWords(input, 1)

	assert.Equal(t, Success, res, "they should be equal")
	assert.Equal(t, output, resSkip, "they should be equal")

	// skip 10 symbols
	output = "444!"
	resSkip, res = skipSymbols(input, 10)
	assert.Equal(t, Success, res, "they should be equal")
	assert.Equal(t, output, resSkip, "they should be equal")

	// combinate
	output = "333 4444!"
	resSkip, res_1 := skipWords(input, 1)
	resSkip, res_2 := skipSymbols(resSkip, 3)
	assert.Equal(t, Success, res_1, "they should be equal")
	assert.Equal(t, Success, res_2, "they should be equal")
	assert.Equal(t, output, resSkip, "they should be equal")
}
