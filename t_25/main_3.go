package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

/*
Способ через бесконечный цикл
*/

func main() {
	const sleepSeconds = 3

	fmt.Println("\nSleep seconds:", sleepSeconds)

	fmt.Println("Start time:", time.Now())

	sleep3(sleepSeconds * time.Second)

	fmt.Println("End time:", time.Now())
}

func sleep3(duration time.Duration) {
	timeNow := time.Now()
	stopTime := timeNow.Add(duration)

	for time.Now().Before(stopTime) {
	}
}

/*
Sleep seconds: 3
Start time: 2024-04-07 22:19:16.420607309 +0300 MSK m=+0.000065045
End time: 2024-04-07 22:19:19.420700945 +0300 MSK m=+3.000158681
*/
