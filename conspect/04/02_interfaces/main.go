// Интерфейсы

package main

import (
	"GeekBrains-1/conspect/04/02_interfaces/interfaces"
	"fmt"
)

func main() {

	// s1 := Square{5}
	// c1 := Circle{3}
	// s1.Area()
	// c1.Area()
	// fmt.Println(s1)
	// fmt.Println(c1)

	firstCircle := interfaces.Circle{10}
	secondCircle := interfaces.Circle{15}
	firstSquare := interfaces.Square{10}

	fmt.Println(firstCircle, secondCircle, firstSquare)
	fmt.Println(firstCircle.Area(), firstCircle.Perim())

	totalAreas := interfaces.SummAreas(firstCircle, secondCircle, firstSquare)
	fmt.Println(totalAreas)
	totalPerims := interfaces.SummPerims(firstCircle, secondCircle, firstSquare)
	fmt.Println(totalPerims)
}
