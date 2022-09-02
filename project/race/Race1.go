package race

//  Ошибка множественной записи в мапу

import (
	"fmt"
	"runtime"
)

func Race1() {
	runtime.GOMAXPROCS(0)

	var counters = map[int]int{}

	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int) {
			for j := 0; j < 5; j++ {
				counters[th*10+j]++
			}
		}(counters, i)
	}

	fmt.Scanln()

	fmt.Println("counters result", counters)
}
