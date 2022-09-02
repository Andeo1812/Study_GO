package _select

import (
	"fmt"
	"sync"
)

func Select4() {
	// runtime.GOMAXPROCS(1)
	cancelCh := make(chan bool)

	dataCh := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(cancelCh chan bool, dataCh chan int) {
			defer wg.Done()

			val := 0
			for {
				select {
				case <-cancelCh:
					fmt.Println("cancelled")
					return
				case dataCh <- val:
					val++
				}
			}
		}(cancelCh, dataCh)
	}

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	for curVal := range dataCh {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelCh <- true // будет ждать отработки в select
			// break
		}
	}

}
