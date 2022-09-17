package actions

import "fmt"

func InitCalculate(expression string) (float64, error) {
	fmt.Println("----------------")
	res, _, err := actionSum(0, expression)
	return res, err
}
