package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

/*
Способ через context.WithTimeout()
*/

func main() {
	const sleepSeconds = 3

	fmt.Println("\nSleep seconds:", sleepSeconds)

	fmt.Println("Start time:", time.Now())

	sleep2(sleepSeconds * time.Second)

	fmt.Println("End time:", time.Now())
}

func sleep2(duration time.Duration) {
	contextTimeout, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	<-contextTimeout.Done()
}

/*
Sleep seconds: 3
Start time: 2024-04-07 21:53:30.967586959 +0300 MSK m=+0.000049124
End time: 2024-04-07 21:53:33.968107855 +0300 MSK m=+3.000570060
*/
