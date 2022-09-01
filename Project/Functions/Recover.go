package Functions

import "fmt"

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend SECOND:", err)
		}

		fmt.Println("Second defer")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend FIRST:", err)
			panic("second panic")
		}
	}()

	fmt.Println("useful")
	panic("something bad happend")
}

func Recover() {
	fmt.Println("Recover")

	deferTest()

	fmt.Println("some some")
}
