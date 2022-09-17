package actions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSumParens(t *testing.T) {
	var input string = "1+(2+(3+4)+5)+(10+10)"
	var output float64 = 35
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestMulParens(t *testing.T) {
	var input string = "(1*4)*5"
	var output float64 = 20
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "(1+2*3*4)"
	output = 25
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "(1+2*3*4)*5"
	output = 125
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "25*(3+1)*(3+3)"
	output = 600
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestDivideParens(t *testing.T) {
	var input string = "(1/4)/5"
	var output float64 = 0.05
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "(1+1/4/2)"
	output = 1.125
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "(1+1/4/2)/5"
	output = 0.225
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	input = "100/(3+1)/(3+2)"
	output = 5
	resExpression = InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestCombinationParens(t *testing.T) {
	var input string = "(25/1+1)*(4+2)"
	var output float64 = 12.4
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)
}
