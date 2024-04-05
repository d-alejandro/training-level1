package main

import (
	"encoding/json"
	"fmt"
)

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

func main() {
	array := [5]string{"cat", "cat", "dog", "cat", "tree"}

	totalMap := make(map[string]struct{})

	for _, value := range array {
		_, isExist := totalMap[value]
		if isExist {
			continue
		}
		totalMap[value] = struct{}{}
	}

	str, err := json.MarshalIndent(totalMap, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("\n", string(str))
}

/*
 {
  "cat": {},
  "dog": {},
  "tree": {}
}
*/
