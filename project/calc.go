package main

import (
	"Modules/project/calc/iohandlers"
	"Modules/project/calc/parser/actions"
)

func main() {
	expression := iohandlers.Input()

	res := actions.InitCalculate(expression)

	iohandlers.Output(res)
}
