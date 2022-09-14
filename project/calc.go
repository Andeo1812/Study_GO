package main

import (
	"Modules/project/calc/iohandlers"
	"Modules/project/calc/parser"
)

func main() {
	expression := iohandlers.Input()

	res := parser.InitCalculate(expression)

	iohandlers.Output(res)
}
