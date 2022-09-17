package actions

import "fmt"

func InitCalculate(expression string) float64 {
	fmt.Println("----------------")
	res, _ := actionSum(0, expression)
	return res
}
