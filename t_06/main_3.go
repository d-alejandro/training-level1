package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

/*
Через закрытие канала.
*/

func main() {
	var waitGroup sync.WaitGroup

	stopChannel := make(chan struct{})

	waitGroup.Add(1)

	go func(channel <-chan struct{}) {
		stopFlag := false

		for !stopFlag {
			select {
			case _, opened := <-channel:
				if !opened {
					fmt.Println("Goroutine stop")
					stopFlag = true
					continue
				}
			default:
				fmt.Println("Goroutine run")
				time.Sleep(time.Second)
			}
		}

		waitGroup.Done()
	}(stopChannel)

	time.Sleep(3 * time.Second)

	close(stopChannel)

	waitGroup.Wait()
}

/*
Goroutine run
Goroutine run
Goroutine run
Goroutine stop
*/
