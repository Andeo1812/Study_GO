package atomic_

import (
	"fmt"
	"sync/atomic" // atomic_2.go
	"time"
)

var totalOperations_ int32

func inc_() {
	atomic.AddInt32(&totalOperations_, 1) // атомарно
}

func Atomic2() {
	for i := 0; i < 1000; i++ {
		go inc_()
	}

	time.Sleep(2 * time.Millisecond)

	fmt.Println("total operation = ", totalOperations_)
}
