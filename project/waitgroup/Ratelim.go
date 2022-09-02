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
	quotaCh <- struct{}{} // ratelim.go, берём свободный слот

	defer wg.Done()
	for j := 0; j < iterationsNum_; j++ {
		fmt.Printf(formatWork_(in, j))

		<-quotaCh
		quotaCh <- struct{}{}

		runtime.Gosched() // даём поработать другим горутинам
	}

	<-quotaCh // ratelim.go, возвращаем слот
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
