package parser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//func TestSum(t *testing.T) {
//	var input string = "10+10+10"
//
//	// skip many symbols
//	var output float64 = 30
//	resExpression := Calculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//}

func TestSub(t *testing.T) {
	var input string = "10-10"

	// simple
	var output float64 = 0
	resExpression := Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "10-10-1000"

	// simple
	output = -1000
	resExpression = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "10-100-110-10"

	// simple
	output = -210
	resExpression = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "10-10-10-10-10000"

	// simple
	output = -10020
	resExpression = Calculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

//func TestMultiply(t *testing.T) {
//	var input string = "10*10*10"
//
//	// skip many symbols
//	var output float64 = 1000
//	resExpression := Calculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//}

//func TestDivide(t *testing.T) {
//	var input string = "1000/10/10"
//
//	// skip many symbols
//	var output float64 = 10
//	resExpression := Calculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//}
