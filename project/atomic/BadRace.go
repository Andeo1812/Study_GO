package atomic_

import (
	"fmt"
	"time"
)

var totalOperations int32 = 0

func inc() {
	totalOperations++
}

func BadRace() {
	for i := 0; i < 1000; i++ {
		go inc()
	}

	time.Sleep(1000 * time.Millisecond)

	// ождается 1000, но по факту будет меньше
	fmt.Println("total operation = ", totalOperations)
}
