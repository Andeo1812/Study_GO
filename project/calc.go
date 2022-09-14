package main

import (
	"Modules/project/calc/iohandlers"
	"Modules/project/calc/parser"
)

func main() {
	expression := iohandlers.Input()

	res, resGet := parser.GetNumber(expression)
	if resGet != nil {
		panic("Bad input - error conversation")
	}

	iohandlers.Output(res)
}
