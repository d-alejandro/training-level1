package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

type Operation int

const (
	Mul Operation = iota
	Quo
	Add
	Sub
)

func main() {
	a := big.NewFloat(math.Pow(5, 20) + 3)
	b := big.NewFloat(math.Pow(5, 20))

	fmt.Println("\na =", a)
	fmt.Println("b =", b)

	fmt.Println("\na * b", "=", execute(a, b, Mul))
	fmt.Println("a / b", "=", execute(a, b, Quo))
	fmt.Println("a + b", "=", execute(a, b, Add))
	fmt.Println("a - b", "=", execute(a, b, Sub))
}

func execute(a, b *big.Float, operation Operation) *big.Float {
	total := new(big.Float)

	switch operation {
	case Mul:
		total.Mul(a, b)
	case Quo:
		total.Quo(a, b)
	case Add:
		total.Add(a, b)
	case Sub:
		total.Sub(a, b)
	default:
		log.Fatal(errors.New("operation not recognized"))
	}

	return total
}

/*
a = 9.5367431640628e+13
b = 9.5367431640625e+13

a * b = 9.094947017729569e+27
a / b = 1.0000000000000315
a + b = 1.90734863281253e+14
a - b = 3
*/
