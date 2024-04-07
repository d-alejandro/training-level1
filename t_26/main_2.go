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

/*
Способ с применением slices.Compact()
*/

func main() {
	const (
		string1 = "abcd"
		string2 = "abCdefAaf"
		string3 = "aabcd"
	)

	fmt.Println(string1, "=", check2(string1))
	fmt.Println(string2, "=", check2(string2))
	fmt.Println(string3, "=", check2(string3))
}

func check2(text string) bool {
	textToLower := strings.ToLower(text)
	arrayRune := []rune(textToLower)

	slices.Sort(arrayRune)
	arrayRuneCompact := slices.Compact(arrayRune)

	if len(arrayRune) != len(arrayRuneCompact) {
		return false
	}
	return true
}

/*
abcd = true
abCdefAaf = false
aabcd = false
*/
