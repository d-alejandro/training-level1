package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
}

func (human *Human) GetName() string {
	return human.name
}

func (human *Human) GetAge() int {
	return human.age
}

type Action struct {
	Human
}

func (action *Action) GetName() string {
	return action.name + " (from Action)"
}

func main() {
	human := &Human{
		name: "Ivan",
		age:  30,
	}

	action := &Action{*human}

	fmt.Println("Name:", action.GetName())
	fmt.Println("Age:", (*Action).GetAge(action))

	action.age = 33
	fmt.Println("New Age:", action.GetAge())
}

/*
Name: Ivan (from Action)
Age: 30
New Age: 33
*/
