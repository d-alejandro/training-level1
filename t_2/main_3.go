package main

import "fmt"

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
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

		go func(value int, channel chan<- Square) {
			channel <- Square{value, value * value}
		}(value, channelArray[key])
	}

	var squares []Square

	counter := 0

	for counter < arrayLength {
		select {
		case square := <-channelArray[0]:
			squares = append(squares, square)
		case square := <-channelArray[1]:
			squares = append(squares, square)
		case square := <-channelArray[2]:
			squares = append(squares, square)
		case square := <-channelArray[3]:
			squares = append(squares, square)
		case square := <-channelArray[4]:
			squares = append(squares, square)
		}

		counter++
	}

	for _, square := range squares {
		fmt.Println(square.key, "=>", square.value)
	}
}

/*
10 => 100
2 => 4
6 => 36
8 => 64
4 => 16
*/
