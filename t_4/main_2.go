package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
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

func main() {
	var (
		waitGroup     sync.WaitGroup
		channelNumber uint
		stopFlag      bool
	)

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	workerCount := getWorkerCount()
	workerChannels := make([]chan int, workerCount)
	makeWorkers(workerCount, workerChannels, &waitGroup)

	for !stopFlag {
		select {
		case systemSignal := <-channel:
			fmt.Println("\nGot signal:", systemSignal)

			for channelNumber = 0; channelNumber < workerCount; channelNumber++ {
				close(workerChannels[channelNumber])
			}

			fmt.Println("Publisher stop")

			stopFlag = true
		default:
			value := rand.IntN(1000)

			for channelNumber = 0; channelNumber < workerCount; channelNumber++ {
				workerChannels[channelNumber] <- value
				fmt.Println("Sent", value, "to Worker #", channelNumber+1)
			}
		}
	}

	waitGroup.Wait()
}

func getWorkerCount() uint {
	fmt.Println("\nВведите кол-во воркеров:")

	var workerCount uint

	if _, err := fmt.Scanln(&workerCount); err != nil {
		log.Fatal(err)
	}

	fmt.Println("N =", workerCount)

	return workerCount
}

func makeWorkers(workerCount uint, workerChannels []chan int, waitGroup *sync.WaitGroup) {
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

/*
Введите кол-во воркеров:
3
N = 3
Sent 146 to Worker # 1
Sent 146 to Worker # 2
Sent 146 to Worker # 3
Worker # 2 got 146
Worker # 3 got 146
Worker # 1 got 146
Sent 755 to Worker # 1
Sent 755 to Worker # 2
Sent 755 to Worker # 3
^C
Worker # 3 got 755
Worker # 2 got 755
Worker # 1 got 755
Sent 454 to Worker # 1
Sent 454 to Worker # 2
Sent 454 to Worker # 3

Got signal: interrupt
Publisher stop
Worker # 1 got 454
Worker # 2 got 454
Worker # 3 got 454
Worker # 3 stop
Worker # 1 stop
Worker # 2 stop
*/
