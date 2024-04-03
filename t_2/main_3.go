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
	array := [5]int{2, 4, 6, 8, 10}

	channelSquare := make(chan Square)

	for _, value := range array {
		go func(value int, channel chan<- Square) {
			channel <- Square{value, value * value}
		}(value, channelSquare)
	}

	var squares []Square

	counter := 0
	arrayLength := len(array)

	for {
		if counter == arrayLength {
			break
		}

		squares = append(squares, <-channelSquare)

		counter++
	}

	for _, square := range squares {
		fmt.Println(square.key, "=>", square.value)
	}
}

/*
10 => 100
2 => 4
4 => 16
6 => 36
8 => 64
*/
