package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

/*
Программа завершается после чтения из канала данных от context.WithTimeout.
*/

func main() {
	const sleepSeconds = 3

	var waitGroup sync.WaitGroup

	channel := make(chan int)

	waitGroup.Add(1)

	go func(readChan <-chan int) {
		fmt.Println("Goroutine start")
		for value := range readChan {
			fmt.Println("Got:", value)
		}
		fmt.Println("Goroutine stop")
		waitGroup.Done()
	}(channel)

	timeStart := time.Now()

	contextTimeout, cancel := context.WithTimeout(context.Background(), sleepSeconds*time.Second)

	stopFlag := false

	for !stopFlag {
		select {
		case <-contextTimeout.Done():
			close(channel)
			stopFlag = true
			fmt.Println("Start time:", timeStart)
			fmt.Println("Stop time:", time.Now())
		default:
			value := rand.IntN(1000)
			fmt.Println("Sent:", value)
			channel <- value

			time.Sleep(time.Second)
		}
	}

	cancel()

	waitGroup.Wait()
}

/*
Sent: 496
Goroutine start
Got: 496
Sent: 14
Got: 14
Sent: 398
Got: 398
Goroutine stop
Start time: 2024-04-04 16:26:38.251008887 +0300 MSK m=+0.000039175
Stop time: 2024-04-04 16:26:41.252839103 +0300 MSK m=+3.001869441
*/
