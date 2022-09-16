package actions

func InitCalculate(expression string) float64 {
	res, _ := actionSum(0, expression)
	return res
}
