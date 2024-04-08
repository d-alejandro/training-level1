package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
*/

/*
Способ с применением структуры sync.Mutex
*/

type Counter struct {
	sync.Mutex
	total uint64
}

func NewCounter() *Counter {
	return &Counter{}
}

func (counter *Counter) increment() {
	counter.Lock()
	defer counter.Unlock()
	counter.total++
}

func (counter *Counter) getTotal() uint64 {
	return counter.total
}

func main() {
	var waitGroup sync.WaitGroup

	counter := NewCounter()

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go runCounter(counter, &waitGroup)
	}

	waitGroup.Wait()

	fmt.Println("Total:", counter.getTotal())
}

func runCounter(counter *Counter, waitGroup *sync.WaitGroup) {
	for x := 0; x < 100; x++ {
		counter.increment()
	}
	waitGroup.Done()
}

/*
Total: 1000
*/
