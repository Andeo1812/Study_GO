package basic_functional

import (
	"fmt"
)

func Loop() {
	fmt.Println("Loop")

	//  Цикл без условия, while(true) OR for(;;)
	for {
		fmt.Println("loop iteration")
		break
	}

	//  Цикл с условием, while(rule)
	rule := true
	for rule {
		fmt.Println("loop iteration")
		rule = false
	}

	//  Цикл с условием и блоком инициализации
	for i := 0; i < 2; i++ {
		fmt.Println("loop iteration", i)
		if i == 4 {
			continue
		}
	}

	//  Операции по slice
	sl := []int{1, 2, 3}
	idx := 0

	for idx < len(sl) {
		fmt.Println("while style loop, idx :", idx, "value:", sl[idx])
		idx++
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("c style loop, idx :", i, "value:", sl[i])
	}

	for idx := range sl {
		fmt.Println("range style loop, idx :", idx, "value:", sl[idx])
	}

	for idx, val := range sl {
		fmt.Println("range style loop, idx-value :", idx, "value:", val)
	}

	//  Операции по map
	profile := map[int]string{1: "Vasa", 2: "Oleg"}

	for key := range profile {
		fmt.Println("range map by key", key)
	}

	for key, val := range profile {
		fmt.Println("range map by key-val", key, "value:", val)
	}

	for _, val := range profile {
		fmt.Println("range map by val", val)
	}

	str := "Hello world"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}
