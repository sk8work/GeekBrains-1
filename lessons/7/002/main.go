package main

import "fmt"

type A struct {
	arr []chan int
}

func (a *A) Start() {
	a.arr = make([]chan int, 4)
	for i := 0; i < 4; i++ {
		a.arr[i] = make(chan int)
	}
	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Println(<-a.arr[i])
		}(i)

	}
}

func (a *A) WriteTo(i int, elem int) {
	a.arr[i] <- elem
}

func (a *A) ReSetup() {
	a.arr = nil
	a.Start()
}

func main() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 10; i++ {
		go func(a int) {
			aa := &A{}
			aa.Start()
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
