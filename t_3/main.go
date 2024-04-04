package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

/*
Запись результата выполняется в atomic.Uint32.
*/

func main() {
	const waitGroupDelta = 1

	var (
		waitGroup     sync.WaitGroup
		atomicCounter atomic.Uint32
	)

	array := [5]uint8{2, 4, 6, 8, 10}

	for _, value := range array {
		waitGroup.Add(waitGroupDelta)

		go func(value uint8) {
			atomicCounter.Add(uint32(value * value))
			waitGroup.Done()
		}(value)
	}

	waitGroup.Wait()

	fmt.Println(atomicCounter.Load())
}

/*
220
*/
