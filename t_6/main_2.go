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
Через отдельный канал, получение данных из которого сигнализирует об остановке горутины.
*/

func main() {
	var waitGroup sync.WaitGroup

	stopChannel := make(chan struct{})

	waitGroup.Add(1)

	go func(channel <-chan struct{}) {
		stopFlag := false

		for !stopFlag {
			select {
			case <-channel:
				fmt.Println("Goroutine stop")
				stopFlag = true
			default:
				fmt.Println("Goroutine run")
				time.Sleep(time.Second)
			}
		}

		waitGroup.Done()
	}(stopChannel)

	time.Sleep(3 * time.Second)

	stopChannel <- struct{}{}
	close(stopChannel)

	waitGroup.Wait()
}

/*
Goroutine run
Goroutine run
Goroutine run
Goroutine run
Goroutine stop
*/
