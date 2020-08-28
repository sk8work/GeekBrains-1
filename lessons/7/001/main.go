package main

import (
	"fmt"
	"runtime"
)

func main() {
	// fmt.Println(runtime.GOMAXPROCS(8))
	// runtime.GOMAXPROCS(8)
	// go func() {
	// 	fmt.Print("Hello, World!")
	// }()
	// time.Sleep(time.Nanosecond)

	// for i := 1; i < 7; i++ {
	// 	go factorial(i)
	// }
	// fmt.Println("The End")

	myOs := runtime.GOOS
	switch myOs {
	case "windows":
		fmt.Println("WIN")
	}

}

// func factorial(n int) {
// 	if n < 1 {
// 		fmt.Println("Invalid input number")
// 		return
// 	}
// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
// 	}
// 	fmt.Println(n, "-", result)

// }
