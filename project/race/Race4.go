package race

import (
	"fmt"
	"sync"
	"time"
)

// look sync.Map

func Race4() {
	var counters = map[int]int{}

	mu := &sync.RWMutex{}

	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int) {
			for j := 0; j < 5; j++ {
				if (th*10+j)%7 == 0 {
					mu.Lock()
					counters[th*10+j]++
					mu.Unlock()
				} else {
					mu.RLock()
					fmt.Printf("[%d,%d] result %v\n", th, j, counters)
					mu.RUnlock()
				}
			}
		}(counters, i)
	}

	time.Sleep(100 * time.Millisecond)

	mu.RLock()
	fmt.Println("counters result", counters)
	mu.RUnlock()
}
