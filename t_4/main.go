package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

var (
	atomicStopFlag atomic.Bool
	waitGroup      sync.WaitGroup
)

func main() {
	workerCount := getWorkerCountValue()

	workerChannels := make([]chan int, workerCount)

	runWorkers(workerCount, workerChannels)

	runPublisher(workerCount, workerChannels)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	systemSignal := <-channel
	fmt.Println("\nGot signal:", systemSignal)

	atomicStopFlag.Store(true)

	waitGroup.Wait()
}

func getWorkerCountValue() uint {
	fmt.Println("\nВведите кол-во воркеров:")

	var workerCount uint

	if _, err := fmt.Scanln(&workerCount); err != nil {
		log.Fatal(err)
	}

	fmt.Println("N =", workerCount)

	return workerCount
}

func runWorkers(workerCount uint, workerChannels []chan int) {
	var channelNumber uint

	for channelNumber = 0; channelNumber < workerCount; channelNumber++ {
		workerChannels[channelNumber] = make(chan int)

		waitGroup.Add(1)

		go func(workerNumber uint, channel <-chan int) {
			for number := range channel {
				time.Sleep(2 * time.Second)
				fmt.Println("Worker #", workerNumber, "got", number)
			}

			fmt.Println("Worker #", workerNumber, "stop")

			waitGroup.Done()
		}(channelNumber+1, workerChannels[channelNumber])
	}
}

func runPublisher(workerCount uint, workerChannels []chan int) {
	var channelNumber uint

	waitGroup.Add(1)

	go func() {
		for !atomicStopFlag.Load() {
			value := rand.IntN(1000)

			for channelNumber = 0; channelNumber < workerCount; channelNumber++ {
				workerChannels[channelNumber] <- value
				fmt.Println("Sent", value, "to Worker #", channelNumber+1)
			}
		}

		for channelNumber = 0; channelNumber < workerCount; channelNumber++ {
			close(workerChannels[channelNumber])
		}

		fmt.Println("Publisher stop")

		waitGroup.Done()
	}()
}

/*
Введите кол-во воркеров:
3
N = 3
Sent 470 to Worker # 1
Sent 470 to Worker # 2
Sent 470 to Worker # 3
Worker # 1 got 470
Worker # 3 got 470
Sent 168 to Worker # 1
Worker # 2 got 470
Sent 168 to Worker # 2
Sent 168 to Worker # 3
^C
Got signal: interrupt
Worker # 1 got 168
Sent 550 to Worker # 1
Worker # 2 got 168
Sent 550 to Worker # 2
Worker # 3 got 168
Sent 550 to Worker # 3
Publisher stop
Worker # 2 got 550
Worker # 1 got 550
Worker # 3 got 550
Worker # 3 stop
Worker # 1 stop
Worker # 2 stop
*/
