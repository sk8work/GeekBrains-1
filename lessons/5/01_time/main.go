package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")
	time.Sleep(time.Second)
	// time.Sleep(time.Duration(4000000000)) // Здесь наносекунды
	fmt.Println("World")

	// t := time.Now()
	t1 := time.Date(1996, time.April, 16, 14, 15, 02, 0, time.UTC)
	fmt.Println(time.Since(t1))
}
