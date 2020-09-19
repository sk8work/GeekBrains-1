package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// fmt.Println(runtime.GOMAXPROCS(10))
	// go func() {
	// 	fmt.Println("Hello world!!!")
	// }()
	// time.Sleep(time.Nanosecond)
	// for i := 0; i <= 100000; i++ {
	// }
	runtime.GOMAXPROCS(1)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(500 * time.Millisecond)
}

// 0857224500
