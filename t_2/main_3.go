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

		go func(value int, channel chan<- Square) {
			channel <- Square{value, value * value}
			close(channel)
		}(value, channelArray[key])
	}

	var squares []Square

	counter := 0

	for counter < arrayLength {
		select {
		case square, opened := <-channelArray[0]:
			if !opened {
				continue
			}
			squares = append(squares, square)
		case square, opened := <-channelArray[1]:
			if !opened {
				continue
			}
			squares = append(squares, square)
		case square, opened := <-channelArray[2]:
			if !opened {
				continue
			}
			squares = append(squares, square)
		case square, opened := <-channelArray[3]:
			if !opened {
				continue
			}
			squares = append(squares, square)
		case square, opened := <-channelArray[4]:
			if !opened {
				continue
			}
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
