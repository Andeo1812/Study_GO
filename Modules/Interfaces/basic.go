package Functions

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars")
	return "getSomeVars return"
}

func Defer() {
	fmt.Println("Defer")

	defer fmt.Println("After work")

	defer func() {
		fmt.Println(getSomeVars())
	}()

	fmt.Println("useful")
}
