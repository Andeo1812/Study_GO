package main

import (
	"fmt"
	"modules/special_print/modules/basic_tools"
	"modules/special_print/modules/special_print"
)

func basic_demonstration() {
	basic_tools.Vars_1()

	basic_tools.Vars_2()

	basic_tools.Const()
}

func first_func() {
	special_print.Print()

	var name string = "Oleg"
	//fmt.Fscan(os.Stdin, &name)

	special_print.PrintName(name)
}

func run_first() {
	fmt.Println("Hello, World!")

}

func main() {
	run_first()

	first_func()

	basic_demonstration()
}
