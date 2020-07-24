// Интерфейсы

package main

import (
	"fmt"
	"math"
)

type Square struct {
	edge float64
}

func (s Square) Area() float64 {
	return s.edge * s.edge
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius
}

func SummAreas(shapes ...Shape) float64 {
	res := 0.0
	for _, shape := range shapes {
		if shape == nil {
			continue
		}
		res += shape.Area()
	}
	return res
}

type Shape interface {
	Area() float64
}

func main() {
	firstCircle := Circle{10}
	secondCircle := Circle{15}
	firstSquare := Square{10}
	total := SummAreas(firstCircle, secondCircle, firstSquare)

	fmt.Println(total)
}
