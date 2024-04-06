package main

import (
	"fmt"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	array := []int{-7, -3, -1, 2, 3, 7, 9, 17}

	for key, value := range array {
		fmt.Println(key, "=>", value)
	}

	const target = 7

	if key, ok := search(array, target); ok {
		fmt.Println("Target:", target, "key:", key)
	} else {
		fmt.Println("Target:", target, "Not found.")
	}
}

func search(array []int, target int) (int, bool) {
	left := 0
	right := len(array) - 1

	for left <= right {
		middleIndex := (left + right) / 2
		middle := array[middleIndex]

		if target < middle {
			right = middleIndex - 1
		} else if target > middle {
			left = middleIndex + 1
		} else {
			return middleIndex, true
		}
	}

	return 0, false
}

/*
0 => -7
1 => -3
2 => -1
3 => 2
4 => 3
5 => 7
6 => 9
7 => 17
Target: 7 key: 5
*/
