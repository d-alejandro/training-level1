package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

/*
Способ через time.After()
*/

func main() {
	const sleepSeconds = 3

	fmt.Println("\nSleep seconds:", sleepSeconds)

	fmt.Println("Start time:", time.Now())

	sleep(sleepSeconds * time.Second)

	fmt.Println("End time:", time.Now())
}

func sleep(duration time.Duration) {
	timeout := time.After(duration)
	<-timeout
}

/*
Sleep seconds: 3
Start time: 2024-04-07 21:16:48.776913816 +0300 MSK m=+0.000042061
End time: 2024-04-07 21:16:51.778046203 +0300 MSK m=+3.001174448
*/
