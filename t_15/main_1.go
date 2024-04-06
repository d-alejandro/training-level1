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
Проблема данного кода в том, что мы не можем достоверно оценить его работу из-за отсутствия реализации
функции createHugeString(int) string.
Предположим, что функция createHugeString возвращает строку, состоящую из 99 японских иероглифов и 11 латинских букв.
Следовательно, в результате копирования мы получим 34 символа вместо 100-а, что не может являться ожидаемым результатом.
Причина заключается в копировании строки с использованием кода v[:100], который работает с типом string как
со срезом байт и при копировании берёт первые 100 байт. Из-за наличия в строке символов, размер которых превышает один
байт, 100 байт не будут равны 100 символов.
В данном коде глобальная переменная justString используется только в одной функции, что является нецелесообразным
такое объявление.
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]

	fmt.Println("\nИсходная строка:\n" + v)
	fmt.Println("Полученный результат:\n" + justString)
}

func createHugeString(x int) string {
	_ = x
	return "本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本" +
		"本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本Fxxxxxxxxxx"
}

func main() {
	someFunc()
}

/*
Исходная строка:
本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本Fxxxxxxxxxx
Полученный результат:
本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本本�
*/