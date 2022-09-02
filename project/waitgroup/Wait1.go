package waitgroup

//  Записываем результат в канал, проблема если горутин много

import (
	"log"
	"time"
)

func Waitgroup1() {
	result := make(chan string)
	go func(out chan<- string) {
		time.Sleep(1 * time.Second)
		log.Println("async operation ready, return result")
		out <- "success"
	}(result)

	time.Sleep(2 * time.Second)
	log.Println("some userful work")

	opStatus := <-result
	log.Println("main goroutine:", opStatus)
}
