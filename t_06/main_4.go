package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

/*
Через context.WithCancel.
*/

func main() {
	var waitGroup sync.WaitGroup

	contextWithCancel, cancel := context.WithCancel(context.Background())

	waitGroup.Add(1)

	go func(contextWithCancel context.Context) {
		stopFlag := false

		for !stopFlag {
			select {
			case <-contextWithCancel.Done():
				fmt.Println("Goroutine stop")
				stopFlag = true
			default:
				fmt.Println("Goroutine run")
				time.Sleep(time.Second)
			}
		}

		waitGroup.Done()
	}(contextWithCancel)

	time.Sleep(3 * time.Second)

	cancel()

	waitGroup.Wait()
}

/*
Goroutine run
Goroutine run
Goroutine run
Goroutine run
Goroutine stop
*/
