package main

import (
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

	timeout := time.After(sleepSeconds * time.Second)

	stopFlag := false

	for !stopFlag {
		select {
		case <-timeout:
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

	waitGroup.Wait()
}

/*
Sent: 270
Goroutine start
Got: 270
Sent: 893
Got: 893
Sent: 891
Got: 891
Goroutine stop
Start time: 2024-04-04 16:10:43.579829259 +0300 MSK m=+0.000037953
Stop time: 2024-04-04 16:10:46.580822658 +0300 MSK m=+3.001031362
*/
