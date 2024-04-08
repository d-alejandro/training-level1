package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

/*
Способ с применением перебора в цикле
*/

func main() {
	fmt.Println("Введите строку:")

	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	inputString = strings.TrimSuffix(inputString, "\n")
	stringArray := strings.Split(inputString, " ")

	count := len(stringArray)
	lastIndex := count - 1
	middleIndex := lastIndex / 2

	for key, value := range stringArray {
		if key <= middleIndex {
			stringArray[lastIndex-key], stringArray[key] = value, stringArray[lastIndex-key]
		} else {
			break
		}
	}

	outputString := strings.Join(stringArray, " ")

	fmt.Println(outputString)
}

/*
Введите строку:
▁▂▃▄▅▆▇█ snow dog sun
sun dog snow ▁▂▃▄▅▆▇█
*/
