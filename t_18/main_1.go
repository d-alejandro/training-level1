package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/

/*
Способ с применением структуры atomic.Uint64
*/

type AtomicCounter struct {
	atomic atomic.Uint64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

func (atomicCounter *AtomicCounter) increment() {
	atomicCounter.atomic.Add(1)
}

func (atomicCounter *AtomicCounter) getTotal() uint64 {
	return atomicCounter.atomic.Load()
}

func main() {
	var waitGroup sync.WaitGroup

	atomicCounter := NewAtomicCounter()

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)

		go func(atomicCounter *AtomicCounter) {
			for x := 0; x < 100; x++ {
				atomicCounter.increment()
			}
			waitGroup.Done()
		}(atomicCounter)
	}

	waitGroup.Wait()

	fmt.Println("Total:", atomicCounter.getTotal())
}

/*
Total: 1000
*/
