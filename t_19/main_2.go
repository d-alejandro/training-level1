package main

import (
	"fmt"
	"log"
	"slices"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

/*
Способ с применением slices.Reverse()
*/

func main() {
	var inputString string

	fmt.Println("Введите строку:")

	if _, err := fmt.Scan(&inputString); err != nil {
		log.Fatal(err)
	}

	arrayRune := []rune(inputString)
	slices.Reverse(arrayRune)

	fmt.Println(string(arrayRune))
}

/*
Введите строку:
❶❷❸ЫмаldJQ→↓♛♞♫▁▂▃▄▅▆▇█
█▇▆▅▄▃▂▁♫♞♛↓→QJdlамЫ❸❷❶
*/
