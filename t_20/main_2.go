package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
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

	slices.Reverse(stringArray)

	outputString := strings.Join(stringArray, " ")

	fmt.Println(outputString)
}

/*
Введите строку:
▁▂▃▄▅▆▇█ snow dog sun
sun dog snow ▁▂▃▄▅▆▇█
*/
