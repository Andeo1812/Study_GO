package parser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//func TestSum(t *testing.T) {
//	var input string = "10+10+10"
//	var output float64 = 30
//	resExpression := InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//}
//
//func TestSub(t *testing.T) {
//	var input string = "10-10"
//	var output float64 = 0
//	resExpression := InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//
//	input = "10-10-1001"
//	output = -1001
//	resExpression = InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//
//	input = "10-100-110-10"
//	output = -210
//	resExpression = InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//
//	input = "10-10-10-10-10000"
//	output = -10020
//	resExpression = InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//}

func TestMultiply(t *testing.T) {
	var input string = "6*3*5"
	var output float64 = 15
	//resExpression := InitCalculate(input)
	//require.Equal(t, output, resExpression, "they should be equal", input)

	output = 90
	resExpression, _ := action(6, "3*5", "*")
	require.Equal(t, output, resExpression, "they should be equal", input)
}

func TestDivide(t *testing.T) {
	var input string = "300/25/4"
	var output float64 = 4
	resExpression := InitCalculate(input)
	require.Equal(t, output, resExpression, "they should be equal", input)

	//input = "6+300/25/4"
	//output = 10
	//resExpression = InitCalculate(input)
	//require.Equal(t, output, resExpression, "they should be equal", input)
}

//
//func TestCombination(t *testing.T) {
//	var input string = "2+2*10"
//	var output float64 = 22
//	resExpression := InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//
//	input = "2+2*2/2"
//	output = 4
//	resExpression = InitCalculate(input)
//	require.Equal(t, output, resExpression, "they should be equal", input)
//}
