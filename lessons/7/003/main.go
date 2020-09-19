package main

import (
	"fmt"
)

func main() {
	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println("Hello world")
	// 	}()
	// }
	// wg.Wait()

	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				done <- struct{}{}
			}()
			fmt.Println("Hello world")
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
