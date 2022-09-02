package gorutines

//  Goroutine - легковесный поток, не системный тред

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	iterationsNum = 6
	goroutinesNum = 6
)

func doWork(th int) {
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(th, j))

		//  Данный sleep - выполнение 1 итерации на нескольких горутинах
		//  time.Sleep(time.Millisecond)

		//  Самостоятельный вызов планировщика - обычно не применяют
		//  runtime.Gosched()
	}

	go func() {
		fmt.Println("kek4")
	}()
}

func Goroutines() {
	//  runtime.GOMAXPROCS(0)
	//  Планировщик сам может остановить бесконечный цикл
	for i := 0; i < goroutinesNum; i++ {
		//  go - запустить на горутинах (в асинхронном варианте)
		go doWork(i)

		go func() {
			fmt.Println("kek")
		}()
	}

	//  Если main отработает до окончания работы горутин, то работа остановится и горутины не доработают
	time.Sleep(2 * time.Second)
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}

func imports() {
	fmt.Println(time.Millisecond, runtime.NumCPU())
}
