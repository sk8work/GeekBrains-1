package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(ch)
	// &p := 0
	p := <-ch
	fmt.Println(p)
	p = <-ch
	fmt.Println(p)
	p = <-ch
	fmt.Println(p)
}
