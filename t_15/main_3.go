package main

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

var justString string

func someFunc() {
    v := createHugeString(1 << 10)
    justString = v[:100]
}

func main() {
    someFunc()
}
*/

/*
Пример для текста с символами unicode
*/

func main() {
	someFunc2()
}

func someFunc2() {
	v := createHugeString2(1 << 10)

	_justString := string(v[:100:100])

	fmt.Println("\nКол-во символов в строке из функции createHugeString(int):", strings.Count(string(v), "⓴"))
	fmt.Println("Кол-во символов в переменной justString после копирования:", strings.Count(_justString, "⓴"))
}

func createHugeString2(symbolCount int) []rune {
	symbol := []rune("⓴")

	runes := make([]rune, symbolCount)

	for key := range runes {
		runes[key] = symbol[0]
	}

	return runes
}

/*
Кол-во символов в строке из функции createHugeString(int): 1024
Кол-во символов в переменной justString после копирования: 100
*/
