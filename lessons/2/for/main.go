package main

import "fmt"

func main() {
	a := 0
	fmt.Println("Enter a max value of count: ")
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Println(i)
	}
}
