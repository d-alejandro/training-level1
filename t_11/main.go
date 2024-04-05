package main

import (
	"fmt"
	"slices"
)

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	slice1 := []int{3, 9, 3, 0, 3, 7, 3, 8, 7, 5}
	slice2 := []int{0, 7, 3, 7, 3, 4, 1, 5, 2, 6, 7}

	var totalSlice []int

	for _, valueInSlice1 := range slice1 {
		if slices.Contains(slice2, valueInSlice1) {
			quantityValuesInTotalSlice := countValuesInSlice(totalSlice, valueInSlice1)
			quantityValuesInSlice2 := countValuesInSlice(slice2, valueInSlice1)

			if quantityValuesInTotalSlice < quantityValuesInSlice2 {
				totalSlice = append(totalSlice, valueInSlice1)
			}
		}
	}

	fmt.Printf("\nSlice# 1:\n%v\n", slice1)
	fmt.Printf("Slice# 2:\n%v\n", slice2)
	fmt.Printf("Total slice:\n%v\n", totalSlice)
}

func countValuesInSlice(slice []int, inputValue int) int {
	counter := 0

	for _, value := range slice {
		if value == inputValue {
			counter++
		}
	}

	return counter
}

/*
Slice# 1:
[3 9 3 0 3 7 3 8 7 5]
Slice# 2:
[0 7 3 7 3 4 1 5 2 6 7]
Total slice:
[3 3 0 7 7 5]
*/
