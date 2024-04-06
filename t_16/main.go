package main

import (
	"fmt"
	"slices"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	array := []int{9, 7, 3, 2, 1, 0}
	fmt.Println(array)

	sortedArray := slices.Clone(array)
	sort(sortedArray)
	fmt.Println(sortedArray)
}

func sort(array []int) {
	runSorting(array, 0, len(array)-1)
}

func runSorting(array []int, startIndex, endIndex int) {
	if startIndex < endIndex {
		left := startIndex
		right := endIndex

		middleIndex := (left + right) / 2
		middle := array[middleIndex]

		for left < right {
			for array[left] < middle {
				left++
			}

			for array[right] > middle {
				right--
			}

			if left <= right {
				array[left], array[right] = array[right], array[left]
				left++
				right--
			}
		}

		runSorting(array, startIndex, right)
		runSorting(array, left, endIndex)
	}
}

/*
[9 7 3 2 1 0]
[0 1 2 3 7 9]
*/
