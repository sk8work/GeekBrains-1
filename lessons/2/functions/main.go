package main

import (
	"fmt"
	"math"
)

func funcName(b bool) (bool, int, string) {
	return b, 42, "sorok dva"
}

func calcParams(a, b float64, cb func()) (float64, float64) {
	c := math.Sqrt(a*a + b*b)
	hep := a + b + c
	cb()
	return c, hep
}

func main() {
	condition, _, str := funcName(true)
	fmt.Println(condition)
	// fmt.Println(number)
	fmt.Println(str)

	fmt.Println(calcParams(1, 2, func() {}))
	fmt.Println(calcParams(3, 4, func() { fmt.Println("Hello, World") }))
}
