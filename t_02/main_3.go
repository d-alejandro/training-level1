package main

import "fmt"

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

/*
Запись результата происходит в главном потоке в slice структур Square.
Получение результатов из каналов реализовано через select.
*/

type Square struct {
	key   int
	value int
}

func main() {
	const arrayLength = 5

	array := [arrayLength]int{2, 4, 6, 8, 10}

	channelArray := make([]chan Square, arrayLength)

	for key, value := range array {
		channelArray[key] = make(chan Square)

		go calculateAndSendResult(value, channelArray[key])
	}

	var (
		squares []Square
		counter int
	)

	checkChannel := func(opened bool, square Square, channelNumber int) {

		if opened {
			channelArray[channelNumber] = nil
			squares = append(squares, square)
			counter++
		}
	}

	for counter < arrayLength {
		select {
		case square, opened := <-channelArray[0]:
			checkChannel(opened, square, 0)
		case square, opened := <-channelArray[1]:
			checkChannel(opened, square, 1)
		case square, opened := <-channelArray[2]:
			checkChannel(opened, square, 2)
		case square, opened := <-channelArray[3]:
			checkChannel(opened, square, 3)
		case square, opened := <-channelArray[4]:
			checkChannel(opened, square, 4)
		default:
		}
	}

	for _, square := range squares {
		fmt.Println(square.key, "=>", square.value)
	}
}

func calculateAndSendResult(value int, channel chan<- Square) {
	channel <- Square{value, value * value}
	close(channel)
}

/*
10 => 100
2 => 4
6 => 36
8 => 64
4 => 16
*/
