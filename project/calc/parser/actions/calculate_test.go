package actions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum(t *testing.T) {
	var input string = "10+10+10"
	var output float64 = 30
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestSub(t *testing.T) {
	var input string = "10-10"
	var output float64 = 0
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "10-10-1001"
	output = -1001
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "10-100-110-10"
	output = -210
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "10-10-10-10-10000"
	output = -10020
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestCombinationSubSum(t *testing.T) {
	var input string = "2+2-4-11+20"
	var output float64 = 9
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "1+1+1-1-1-1+2+2-4"
	output = 0
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestMultiply(t *testing.T) {
	var input string = "6*3*5"
	var output float64 = 90
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "6*3*5+4"
	output = 94
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}
func TestCombinationSubSumMul(t *testing.T) {
	var input string = "2+2*2"
	var output float64 = 6
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "2+2*2-2-2*2+2-2"
	output = 0
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestDivide(t *testing.T) {
	var input string = "300/25/4"
	var output float64 = 3
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "300/25/4+41"
	output = 44
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "300/25/4+41"
	output = 44
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "100-300/25/4-41"
	output = 56
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "300/25*5"
	output = 60
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "10/2*4/5"
	output = 4
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestCombination(t *testing.T) {
	var input string = "2+2*10"
	var output float64 = 22
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "2+2*2/2"
	output = 4
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "2/2-2"
	output = -1
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "2+2*2/2+2+2"
	output = 8
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "2/2+2"
	output = 3
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}
