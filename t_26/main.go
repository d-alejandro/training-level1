package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	const (
		string1 = "abcd"
		string2 = "abCdefAaf"
		string3 = "aabcd"
	)

	fmt.Println(string1, "=", check(string1))
	fmt.Println(string2, "=", check(string2))
	fmt.Println(string3, "=", check(string3))
}

func check(text string) bool {
	textToLower := strings.ToLower(text)

	arrayRune := []rune(textToLower)

	slices.Sort(arrayRune)

	arrayRuneCount := len(arrayRune) - 1

	for key := 0; key < arrayRuneCount; key++ {
		if arrayRune[key] == arrayRune[key+1] {
			return false
		}
	}
	return true
}
