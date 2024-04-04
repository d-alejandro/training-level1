package main

import (
	"fmt"
)

/*
Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	var (
		variable   int64 = 77777777777777
		bitNumber1       = 3
		bitNumber2       = 4
	)

	fmt.Printf("\nVariable: %d\n\n", variable)

	bitOffset1 := getBitOffset(bitNumber1)
	bitOffset2 := getBitOffset(bitNumber2)

	format := "Variable to binary / Bit offset (Bit number = %d) / " +
		"Bit number to one / Bit number to zero\n%064b\n%064b\n%064b\n%064b\n"

	fmt.Printf(
		format,
		bitNumber1,
		variable,
		bitOffset1,
		makeBitNumberToOne(variable, bitOffset1),
		makeBitNumberToZero(variable, bitOffset1),
	)

	fmt.Printf(
		"\n"+format,
		bitNumber2,
		variable,
		bitOffset2,
		makeBitNumberToOne(variable, bitOffset2),
		makeBitNumberToZero(variable, bitOffset2),
	)
}

func makeBitNumberToOne(variable int64, bitOffset int64) int64 {
	return variable | bitOffset
}

func makeBitNumberToZero(variable int64, bitOffset int64) int64 {
	return variable &^ bitOffset
}

func getBitOffset(bitNumber int) int64 {
	return int64(1 << bitNumber)
}

/*
Variable: 77777777777777

Variable to binary / Bit offset (Bit number = 3) / Bit number to one / Bit number to zero
0000000000000000010001101011110100001100110100001101110001110001
0000000000000000000000000000000000000000000000000000000000001000
0000000000000000010001101011110100001100110100001101110001111001
0000000000000000010001101011110100001100110100001101110001110001

Variable to binary / Bit offset (Bit number = 4) / Bit number to one / Bit number to zero
0000000000000000010001101011110100001100110100001101110001110001
0000000000000000000000000000000000000000000000000000000000010000
0000000000000000010001101011110100001100110100001101110001110001
0000000000000000010001101011110100001100110100001101110001100001
*/
