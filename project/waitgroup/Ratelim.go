package waitgroup

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum_ = 6
	goroutinesNum_ = 5
	quotaLimit     = 2
)

func startWorker_(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	fmt.Println("start worker", in)

	quotaCh <- struct{}{} // ratelim.go, берём свободный слот

	// ratelim.go, возвращаем слот
	defer func() {
		<-quotaCh
	}()

	defer wg.Done()
	for j := 0; j < iterationsNum_; j++ {
		//  Целиком обработка 2 операций
		//if j % 2 == 0 {
		//	<-quotaCh
		//	quotaCh <- struct{}{}
		//}

		fmt.Printf(formatWork_(in, j))

		<-quotaCh
		quotaCh <- struct{}{}

		runtime.Gosched() // даём поработать другим горутинам явно с лимитом по числу горутин quotaLinit = 2
	}

}

func Ratelim() {
	wg := &sync.WaitGroup{}

	quotaCh := make(chan struct{}, quotaLimit) // ratelim.go

	for i := 0; i < goroutinesNum_; i++ {
		wg.Add(1)

		go startWorker_(i, wg, quotaCh)
	}

	time.Sleep(time.Millisecond)
	
	wg.Wait()
}

func formatWork_(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum_-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}
