package main

import (
	"Modules/Modules/BasicFunctional"
	"Modules/Modules/SpecialPrint"
	"fmt"
)

func basicDemonstration() {
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

	basicDemonstration()
}
