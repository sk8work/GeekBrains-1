package main

import (
	"GeekBrains-1/conspect/03/05_structs/circle"
	"fmt"
	"reflect"
)

func main() {
	c1 := circle.Circle{
		Point: circle.Position{
			X: 1,
			Y: 2,
		},
		Radius: 42,
	}

	fmt.Println(c1)
	fmt.Println(c1.Point.X)
	fmt.Println(c1.Point.Y)
	fmt.Println(c1.Radius)

	c2 := circle.Circle{}
	c2.Point = circle.Position{X: 5, Y: 3}
	c2.Radius = 15
	fmt.Println(c2)

	c3 := circle.Circle{circle.Position{2, 13}, 43}
	fmt.Println(c3)
	fmt.Printf("%v\n", c3)
	fmt.Println(reflect.TypeOf(c3.Radius))
	// str1 := reflect.TypeOf(c3.Radius)
	// str2 := reflect.TypeOf(true)
	// fmt.Println(str1, str2)

	if reflect.TypeOf(c3.Radius) == reflect.TypeOf(true) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
