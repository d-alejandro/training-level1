package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

func main() {
	var waitGroup sync.WaitGroup

	array := []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}

	firstChannel := make(chan int)
	secondChannel := make(chan int)

	waitGroup.Add(2)

	go func(firstChannel <-chan int, secondChannel chan<- int) {
		for value := range firstChannel {
			secondChannel <- value * 2
		}
		close(secondChannel)
		waitGroup.Done()
	}(firstChannel, secondChannel)

	go func(secondChannel <-chan int) {
		for value := range secondChannel {
			fmt.Println(value)
		}
		waitGroup.Done()
	}(secondChannel)

	for _, value := range array {
		firstChannel <- value
	}

	close(firstChannel)

	waitGroup.Wait()
}
