package main

import (
	"fmt"
	"sync"
	atomic2 "sync/atomic"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.

Через глобальную переменную
*/

func main() {
	var (
		waitGroup  sync.WaitGroup
		atomicBool atomic2.Bool
	)

	waitGroup.Add(1)

	go func() {
		for atomicBool.Load() == false {
			fmt.Println("Goroutine run")
			time.Sleep(time.Second)
		}
		fmt.Println("Goroutine stop")
		waitGroup.Done()
	}()

	time.Sleep(3 * time.Second)

	atomicBool.Store(true)

	waitGroup.Wait()
}

/*
Goroutine run
Goroutine run
Goroutine run
Goroutine run
Goroutine stop
*/
