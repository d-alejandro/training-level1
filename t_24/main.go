package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (point *Point) GetX() float64 {
	return point.x
}

func (point *Point) GetY() float64 {
	return point.y
}

func main() {
	point1 := NewPoint(9.5, 10.5)
	point2 := NewPoint(12.5, 6.5)

	sideSquare1 := math.Pow(point1.GetX()-point2.GetX(), 2)
	sideSquare2 := math.Pow(point1.GetY()-point2.GetY(), 2)

	result := math.Pow(sideSquare1+sideSquare2, 0.5)

	fmt.Printf("\nPoint #1:  %+v\n", point1)
	fmt.Printf("Point #2:  %+v\n", point2)

	fmt.Println("Расстояние между точками =", result)
}

/*
Point #1:  &{x:9.5 y:10.5}
Point #2:  &{x:12.5 y:6.5}
Расстояние между точками = 5
*/
