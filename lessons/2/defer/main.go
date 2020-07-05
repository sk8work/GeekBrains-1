// Отложенная функция

package main

import "fmt"

func main() {
	defer fmt.Println("Hello world")
	fmt.Println("Привет, мир")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
