package main

import (
	"Calc/project/calc/iohandlers"
	"Calc/project/calc/parser/actions"
	"github.com/wonderivan/logger"
)

func main() {
	expression := iohandlers.Input()

	res, err := actions.Calculate(expression)
	if err != nil {
		logger.Error(err)
		return
	}

	iohandlers.Output(res)
}
