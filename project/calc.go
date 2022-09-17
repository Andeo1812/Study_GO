package main

import (
	"Modules/project/calc/iohandlers"
	"Modules/project/calc/parser/actions"
	"fmt"
)

func main() {
	expression := iohandlers.Input()

	res, errorCalc := actions.InitCalculate(expression)
	if errorCalc != nil {
		fmt.Println(errorCalc)
	}

	iohandlers.Output(res)
}
