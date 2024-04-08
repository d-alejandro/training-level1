package main

import (
	"fmt"
	"log"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

/*
Способ с применением перебора в цикле
*/

func main() {
	var inputString string

	fmt.Println("Введите строку:")

	if _, err := fmt.Scan(&inputString); err != nil {
		log.Fatal(err)
	}

	arrayRune := []rune(inputString)

	runeCount := len(arrayRune)
	lastIndex := runeCount - 1
	middleIndex := lastIndex / 2

	for key, value := range arrayRune {
		if key <= middleIndex {
			arrayRune[lastIndex-key], arrayRune[key] = value, arrayRune[lastIndex-key]
		} else {
			break
		}
	}

	fmt.Println(string(arrayRune))
}

/*
Введите строку:
❶❷❸ЫмаldJQ→↓♛♞♫▁▂▃▄▅▆▇█
█▇▆▅▄▃▂▁♫♞♛↓→QJdlамЫ❸❷❶
*/
