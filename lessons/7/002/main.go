package main

import (
	"fmt"
	"time"
)

func main() {

	// ch := make(chan int, 10)
	// sl := make([]int, 0, cap(ch))
	// for i := 0; i < cap(ch); i++ {
	// 	ch <- i
	// }
	// fmt.Println(ch)
	// for i := 0; i < cap(ch); i++ {
	// 	sl = append(sl, <-ch)
	// }
	// fmt.Println(sl)

	// ch := make(chan struct{})
	// go func() {
	// 	fmt.Println("Hello world")
	// 	ch <- struct{}{}
	// 	// close(ch)
	// }()
	// <-ch

	// ch := make(chan int, 1)
	// ch <- 1
	// close(ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	ch := make(chan *int)
	d := 42
	go func() {
		fmt.Println(<-ch)
		close(ch)
		fmt.Println(<-ch)
	}()
	ch <- &d
	time.Sleep(300 * time.Millisecond)
}
