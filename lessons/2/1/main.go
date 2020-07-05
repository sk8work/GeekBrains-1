// Обход по сущности

package main

import "fmt"

func main() {
	s := "Hello world"
	for i, v := range s {
		fmt.Println(i, v, string(v))
	}
}
