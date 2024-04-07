package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func main() {
	const i = 4

	array := []rune("Программа")

	fmt.Println("\nИсходный текст:", string(array))

	fmt.Println("Удалить элемент: i =", i)

	array = append(array[:i], array[i+1:]...)

	fmt.Println("Результат:", string(array))
}

/*
Исходный текст: Программа
Удалить элемент: i = 4
Результат: Прогамма
*/
