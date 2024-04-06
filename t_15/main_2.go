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
Корректный пример реализации.

Исходя из того, что у нас возвращается строка только из символов санскрита, то скорректируем код с учётом
одного символа равного 12 байт. Копирование выполним в байтовом представлении без помещения в отдельную
переменную и с последующим помещением на вход функции Clone().
*/

func main() {
	_someFunc()
}

func _someFunc() {
	v := _createHugeString(1 << 10)
	_justString := strings.Clone(v[:100*12])

	fmt.Println("\nКол-во символов в строке из функции createHugeString(int) string:", strings.Count(v, "क्षि"))
	fmt.Println("Кол-во символов в переменной justString после копирования:", strings.Count(_justString, "क्षि"))
}

func _createHugeString(symbolCount int) string {
	return strings.Repeat("क्षि", symbolCount)
}

/*
Кол-во символов в строке из функции createHugeString(int) string: 1024
Кол-во символов в переменной justString после копирования: 100
*/
