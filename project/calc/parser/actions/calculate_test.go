package actions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSumParens(t *testing.T) {
	var input string = "1+(2+(3+4)+5)"
	var output float64 = 15
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
}

//
//func TestMultiplyParens(t *testing.T) {
//	var input string = "6*3*5"
//	var output float64 = 90
//	resExpression := InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//
//	input = "6*3*5+4"
//	output = 94
//	resExpression = InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//}
//
//
//func TestDivideParens(t *testing.T) {
//	var input string = "300/25/4"
//	var output float64 = 3
//	resExpression := InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//
//}
//
//func TestCombinationParens(t *testing.T) {
//	var input string = "2+2*10"
//	var output float64 = 22
//	resExpression := InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//}
