package main

import (
	"Modules/Modules/BasicTools"
	"Modules/Modules/SpecialPrint"
	"fmt"
)

func basicDemonstration() {
	BasicTools.Vars1()

	BasicTools.Vars2()

	BasicTools.Const()

	BasicTools.Array()

	BasicTools.Slice1()

	BasicTools.Slice2()

	BasicTools.String()
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
