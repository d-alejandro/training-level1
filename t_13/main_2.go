package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

/*
Арифметический способ
*/

func main() {
	variable1 := 3
	variable2 := 7

	fmt.Println("\nVariable #1\n", variable1, "\nVariable #2\n", variable2)

	variable1 = variable1 + variable2
	variable2 = variable1 - variable2
	variable1 = variable1 - variable2

	fmt.Println("\nVariable #1\n", variable1, "\nVariable #2\n", variable2)
}

/*
Variable #1
 3
Variable #2
 7

Variable #1
 7
Variable #2
 3
*/
