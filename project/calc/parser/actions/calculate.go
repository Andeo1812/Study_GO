package actions

func Calculate(expression string) (float64, error) {
	res, _, err := actionSum(0, expression)
	return res, err
}
