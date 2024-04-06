package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

type MapWriter struct {
	sync.Mutex
	data map[int]int
}

func NewMapWriter() *MapWriter {
	data := make(map[int]int)
	return &MapWriter{data: data}
}

func (mapWriter *MapWriter) Set(key int, value int) {
	mapWriter.Lock()
	defer mapWriter.Unlock()

	mapWriter.data[key] = value
}

func main() {
	var waitGroup sync.WaitGroup

	mapWriter := NewMapWriter()

	for goroutineNumber := 1; goroutineNumber <= 10; goroutineNumber++ {
		waitGroup.Add(1)

		go func(mapWriter *MapWriter, goroutineNumber int) {
			for y := 0; y < 10; y++ {
				key := goroutineNumber*10 + y
				mapWriter.Set(key, y)
				fmt.Println("Goroutine #", goroutineNumber, "write", key, ":", y)
			}

			waitGroup.Done()
		}(mapWriter, goroutineNumber)
	}

	waitGroup.Wait()
}

/*
Goroutine # 10 write 100 : 0
Goroutine # 10 write 101 : 1
Goroutine # 10 write 102 : 2
Goroutine # 10 write 103 : 3
Goroutine # 10 write 104 : 4
Goroutine # 10 write 105 : 5
Goroutine # 10 write 106 : 6
Goroutine # 10 write 107 : 7
Goroutine # 10 write 108 : 8
Goroutine # 10 write 109 : 9
Goroutine # 8 write 80 : 0
Goroutine # 7 write 70 : 0
Goroutine # 7 write 71 : 1
Goroutine # 7 write 72 : 2
Goroutine # 7 write 73 : 3
Goroutine # 7 write 74 : 4
Goroutine # 7 write 75 : 5
Goroutine # 7 write 76 : 6
Goroutine # 7 write 77 : 7
Goroutine # 7 write 78 : 8
Goroutine # 7 write 79 : 9
Goroutine # 6 write 60 : 0
Goroutine # 6 write 61 : 1
Goroutine # 6 write 62 : 2
Goroutine # 6 write 63 : 3
Goroutine # 6 write 64 : 4
Goroutine # 6 write 65 : 5
Goroutine # 6 write 66 : 6
Goroutine # 6 write 67 : 7
Goroutine # 6 write 68 : 8
Goroutine # 6 write 69 : 9
Goroutine # 5 write 50 : 0
Goroutine # 5 write 51 : 1
Goroutine # 5 write 52 : 2
Goroutine # 5 write 53 : 3
Goroutine # 5 write 54 : 4
Goroutine # 5 write 55 : 5
Goroutine # 1 write 10 : 0
Goroutine # 1 write 11 : 1
Goroutine # 1 write 12 : 2
Goroutine # 5 write 56 : 6
Goroutine # 5 write 57 : 7
Goroutine # 5 write 58 : 8
Goroutine # 5 write 59 : 9
Goroutine # 2 write 20 : 0
Goroutine # 2 write 21 : 1
Goroutine # 9 write 90 : 0
Goroutine # 1 write 13 : 3
Goroutine # 9 write 91 : 1
Goroutine # 2 write 22 : 2
Goroutine # 2 write 23 : 3
Goroutine # 2 write 24 : 4
Goroutine # 2 write 25 : 5
Goroutine # 2 write 26 : 6
Goroutine # 2 write 27 : 7
Goroutine # 2 write 28 : 8
Goroutine # 2 write 29 : 9
Goroutine # 8 write 81 : 1
Goroutine # 8 write 82 : 2
Goroutine # 8 write 83 : 3
Goroutine # 8 write 84 : 4
Goroutine # 8 write 85 : 5
Goroutine # 8 write 86 : 6
Goroutine # 8 write 87 : 7
Goroutine # 8 write 88 : 8
Goroutine # 8 write 89 : 9
Goroutine # 9 write 92 : 2
Goroutine # 9 write 93 : 3
Goroutine # 9 write 94 : 4
Goroutine # 9 write 95 : 5
Goroutine # 9 write 96 : 6
Goroutine # 9 write 97 : 7
Goroutine # 9 write 98 : 8
Goroutine # 9 write 99 : 9
Goroutine # 1 write 14 : 4
Goroutine # 1 write 15 : 5
Goroutine # 1 write 16 : 6
Goroutine # 1 write 17 : 7
Goroutine # 1 write 18 : 8
Goroutine # 1 write 19 : 9
Goroutine # 3 write 30 : 0
Goroutine # 3 write 31 : 1
Goroutine # 4 write 40 : 0
Goroutine # 4 write 41 : 1
Goroutine # 4 write 42 : 2
Goroutine # 4 write 43 : 3
Goroutine # 4 write 44 : 4
Goroutine # 4 write 45 : 5
Goroutine # 4 write 46 : 6
Goroutine # 4 write 47 : 7
Goroutine # 4 write 48 : 8
Goroutine # 4 write 49 : 9
Goroutine # 3 write 32 : 2
Goroutine # 3 write 33 : 3
Goroutine # 3 write 34 : 4
Goroutine # 3 write 35 : 5
Goroutine # 3 write 36 : 6
Goroutine # 3 write 37 : 7
Goroutine # 3 write 38 : 8
Goroutine # 3 write 39 : 9
*/
