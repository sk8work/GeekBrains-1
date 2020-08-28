package main

import (
	"fmt"
)

func funcName(b bool) (bool, string, int) {
	return b, "sorok dva", 42
}

// func calcParams(a, b float64, cb func()) (float64, float64) {
// 	c := math.Sqrt(a*a + b*b)
// 	hep := a + b + c
// 	cb()
// 	return c, hep
// }

func main() {
	a, s, i := funcName(true)
	fmt.Println(a, s, i)
	// 	condition, _, str := funcName(true)
	// 	fmt.Println(condition)
	// 	// fmt.Println(number)
	// 	fmt.Println(str)

	// 	fmt.Println(calcParams(1, 2, func() {}))
	// 	fmt.Println(calcParams(3, 4, func() { fmt.Println("Hello, World") }))
}
