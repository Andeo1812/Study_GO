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

func Atomic() {
	for i := 0; i < 1000; i++ {
		go inc_()
	}

	time.Sleep(2 * time.Millisecond)

	//  race found
	//  fmt.Println("total operation = ", totalOperations_)

	fmt.Println("total operation = ", atomic.LoadInt32(&totalOperations_))
}
