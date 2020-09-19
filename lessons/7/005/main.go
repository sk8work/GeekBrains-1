package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var a int

func main() {
	for i := 0; i < 10000; i++ {
		go func() {
			mu.Lock()
			a++
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}
