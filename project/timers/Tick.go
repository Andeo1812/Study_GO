package timers

import (
	"fmt"
	"time"
)

func Tick() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			// надо останавливать, иначе потечет
			ticker.Stop()
			break
		}
	}
	fmt.Println("total", i)

	// return

	// не может быть остановлен и собран сборщиком мусора
	// используйте если должен работать вечно
	c := time.Tick(time.Second)
	i = 0
	for tickTime := range c {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			break
		}
	}
}