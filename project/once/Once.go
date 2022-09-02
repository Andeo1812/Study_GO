package once

import (
	"fmt"
	"sync"
)

func Init() {
	fmt.Println("Init once")
}

func Once() {
	const routinesNum = 10

	once := &sync.Once{}

	wg := &sync.WaitGroup{}

	wg.Add(routinesNum)
	for i := 0; i < routinesNum; i++ {
		go func() {
			defer wg.Done()

			once.Do(Init)
		}()
	}

	wg.Wait()
}
