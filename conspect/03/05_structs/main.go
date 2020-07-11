package main

import (
	"GeekBrains-1/conspect/03/05_structs/circle"
	"fmt"
)

func main() {
	firstCircle := circle.Circle{
		Point: circle.Position{
			X: 1,
			Y: 2,
		},
		Radius: 42,
	}

	fmt.Println(firstCircle)
	fmt.Println(firstCircle.Point.X)
	fmt.Println(firstCircle.Point.Y)
	fmt.Println(firstCircle.Radius)
}
