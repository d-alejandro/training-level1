package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

/*
Запись результата в map c использованием sync.Mutex.
*/

type SquareCalculation struct {
	sync.Mutex
	Squares map[int]int
}

func NewSquareCalculation() *SquareCalculation {
	squares := make(map[int]int)
	return &SquareCalculation{Squares: squares}
}

func (receiver *SquareCalculation) Set(key int, square int) {
	receiver.Lock()
	defer receiver.Unlock()

	receiver.Squares[key] = square
}

func main() {
	const waitGroupDelta = 1

	var waitGroup sync.WaitGroup

	array := [5]int{2, 4, 6, 8, 10}

	squareCalculation := NewSquareCalculation()

	for _, value := range array {
		waitGroup.Add(waitGroupDelta)

		go func(value int) {
			squareCalculation.Set(value, value*value)
			waitGroup.Done()
		}(value)
	}

	waitGroup.Wait()

	for key, squares := range squareCalculation.Squares {
		fmt.Println(key, "=>", squares)
	}
}

/*
10 => 100
2 => 4
4 => 16
6 => 36
8 => 64
*/
