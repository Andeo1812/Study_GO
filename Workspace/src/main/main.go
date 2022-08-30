package main

import (
	"fmt"
	"modules/special_print/modules/special_print"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	special_print.Print()

	var name string
	fmt.Fscan(os.Stdin, &name)

	special_print.PrintName(name)
}
