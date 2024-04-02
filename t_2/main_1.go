package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	const waitGroupDelta = 1

	var (
		waitGroup sync.WaitGroup
		squares   sync.Map
	)

	array := [5]int{2, 4, 6, 8, 10}

	for _, value := range array {
		waitGroup.Add(waitGroupDelta)

		go func(value int) {
			squares.Store(value, value*value)
			waitGroup.Done()
		}(value)
	}

	waitGroup.Wait()

	squares.Range(func(keyMap, valueMap any) bool {
		fmt.Println(keyMap, "=>", valueMap)
		return true
	})
}

/*
6 => 36
8 => 64
10 => 100
2 => 4
4 => 16
*/
