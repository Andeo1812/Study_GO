package actions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetNumber(t *testing.T) {
	var input string = "-10basd"

	// skip many symbols
	var output float64 = -10
	number, length := GetNumber(input)
	require.Equal(t, length, 3, "they should be equal", input)
	require.Equal(t, output, number, "they should be equal", input)

	input = "10basd"

	// skip many symbols
	output = 10
	number, length = GetNumber(input)
	require.Equal(t, length, 2, "they should be equal", input)
	require.Equal(t, output, number, "they should be equal", input)
}
