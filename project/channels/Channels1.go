package channels

import "fmt"

//  Каналы нужны для коммуникации между горутинами

func Сhannels1() {
	//  Буферизированный канал с емкостью 1
	ch1 := make(chan int, 1)

	go func(in chan int) {
		val := <-in //  ждем записи в канал
		fmt.Println("Go: get from chan", val)
		fmt.Println("Go: after read from chan")
	}(ch1)

	//  Отправка в канал
	ch1 <- 32
	ch1 <- 100500

	fmt.Println("Сhannels: after read from chan")

	//  fmt.Scanln()
}
