package main

import "fmt"

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

Большая строка представлена как срез рун. Копирование происходит с использованием кода hugeString[:100:100].
Указание ёмкости (capacity) позволит выделить для нового среза отдельную память.
После отработки функции _createHugeString память под срез hugeString будет освобождена.
*/

func main() {
	_someFunc()
}

func _someFunc() {
	_justString := _createHugeString(100)

	fmt.Println("Полученный результат:\n" + _justString)
}

func _createHugeString(length int) string {
	hugeString := []rune("本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本" +
		"本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本Fxxxxxxxxxx")

	fmt.Println("\nИсходная строка:\n" + string(hugeString))

	return string(hugeString[:length:length])
}

/*
Исходная строка:
本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本Fxxxxxxxxxx
Полученный результат:
本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本F
*/
