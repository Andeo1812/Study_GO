package main

import (
	"Calc/project/calc/iohandlers"
	"Calc/project/calc/parser/actions"
	"github.com/wonderivan/logger"
	"strings"
)

func main() {
	expression := iohandlers.Input()

	noSpacesExpr := strings.ReplaceAll(expression, " ", "")

	res, err := actions.Calculate(noSpacesExpr)
	if err != nil {
		logger.Error(err)
		return
	}

	iohandlers.Output(res)
}
