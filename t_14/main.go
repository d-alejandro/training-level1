package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	dataMap := map[string]interface{}{
		"int":      7,
		"string":   "string",
		"bool":     true,
		"channel":  make(chan int),
		"channel2": make(chan string),
	}

	for key, value := range dataMap {
		switch value.(type) {
		case int:
			fmt.Printf("%s: %d\n", key, value.(int))
		case string:
			fmt.Printf("%s: %s\n", key, value.(string))
		case bool:
			fmt.Printf("%s: %t\n", key, value.(bool))
		case chan int:
			fmt.Println(key+":", value.(chan int))
		default:
			fmt.Println(key+":", nil)
		}
	}
}

/*
bool: true
channel: 0xc0000281e0
channel2: <nil>
int: 7
string: string
*/
