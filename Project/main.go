package main

import (
	"Modules/Project/BasicFunctional"
	"Modules/Project/Functions"
	"Modules/Project/Interface"
	"Modules/Project/SpecialPrint"
	"Modules/Project/Structs"
	"fmt"
)

func basicDemo() {
	BasicFunctional.Vars1()

	BasicFunctional.Vars2()

	BasicFunctional.Const()

	BasicFunctional.Array()

	BasicFunctional.Slice1()

	BasicFunctional.Slice2()

	BasicFunctional.String()

	BasicFunctional.Map()

	BasicFunctional.Types()

	BasicFunctional.Pointers()

	BasicFunctional.Control()

	BasicFunctional.Loop()
}

func functionsDemo() {
	Functions.Functions()

	Functions.FirstClass()

	Functions.Defer()

	Functions.Recover()
}

func structsDemo() {
	Structs.Structs()

	Structs.Methods()
}

func interfaceDemo() {
	Interface.Basic()

	Interface.Embed()

	Interface.Many()

	Interface.Cast()
}

func demo() {
	basicDemo()

	functionsDemo()

	structsDemo()

	interfaceDemo()
}

func firstFunc() {
	SpecialPrint.Print()

	var name string = "Oleg"
	//fmt.Fscan(os.Stdin, &name)

	SpecialPrint.PrintName(name)
}

func runFirst() {
	fmt.Println("Hello, World!")
}

func main() {
	runFirst()

	firstFunc()

	demo()
}
