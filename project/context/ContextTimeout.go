package _context

import (
	"context"
	"fmt"
	"time"
)

func ContextTimeout() {
	workTime := 50 * time.Millisecond

	ctx, _ := context.WithTimeout(context.Background(), workTime)

	result := make(chan int, 1)

	for i := 0; i <= 10; i++ {
		go студент(ctx, i, result)
	}

	totalFound := 0
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case foundBy := <-result:
			totalFound++

			fmt.Println("студент", foundBy, "задал вопрос")
		}
	}

	fmt.Println("всего вопросов", totalFound)

	time.Sleep(time.Second)
}
