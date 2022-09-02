package channels

import (
	"fmt"
	"runtime"
)

func Сhannels2() {
	runtime.GOMAXPROCS(1)

	in := make(chan int, 1)

	//  Канал только для записи
	go func(out chan<- int) {
		for i := 0; i <= 10; i++ {
			fmt.Println("before", i)
			out <- i

			fmt.Println("after", i)
		}
		// out <- 12

		//  Если не закрыть, то произойдет дедлок, main будет ждать на чтение
		close(out)

		fmt.Println("generator finish")
	}(in)

	for i := range in {
		fmt.Println("\tget", i)
	}

	//  Без range
	//for {
	//	i, isOpen := <-in
	//	if !isOpen {
	//		break
	//	}
	//	fmt.Println("\tget", i, isOpen)
	//}

	// fmt.Scanln()
}
