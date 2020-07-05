// ветвление

package main

import "fmt"

func f() bool {
	return true
}

func main() {
	i := 2

	if i%2 == 0 && i == 2 {
		fmt.Println("Hello")
	} else {
		fmt.Println("World")
	}

	if f() {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
