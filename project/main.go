package main

import (
	"Modules/project/basic_functional"
	"Modules/project/functions"
	"Modules/project/interface"
	"Modules/project/special_print"
	"Modules/project/structs"
	"fmt"
)

func basicDemo() {
	basic_functional.Vars1()

	basic_functional.Vars2()

	basic_functional.Const()

	basic_functional.Array()

	basic_functional.Slice1()

	basic_functional.Slice2()

	basic_functional.String()

	basic_functional.Map()

	basic_functional.Types()

	basic_functional.Pointers()

	basic_functional.Control()

	basic_functional.Loop()
}

func functionsDemo() {
	functions.Functions()

	functions.FirstClass()

	functions.Defer()

	functions.Recover()
}

func structsDemo() {
	structs.Structs()

	structs.Methods()
}

func interfaceDemo() {
	_interface.Basic()

	_interface.Embed()

	_interface.Many()

	_interface.Cast()
}

func demo() {
	basicDemo()

	functionsDemo()

	structsDemo()

	interfaceDemo()
}

func firstFunc() {
	special_print.Print()

	var name string = "Oleg"
	//fmt.Fscan(os.Stdin, &name)

	special_print.PrintName(name)
}

func runFirst() {
	fmt.Println("Hello, World!")
}

func main() {
	runFirst()

	firstFunc()

	demo()
}
