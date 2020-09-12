package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(500 * time.Millisecond)
}

// 0857224500
