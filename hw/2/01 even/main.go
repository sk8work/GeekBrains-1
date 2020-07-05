// Написать функцию, которая определяет, четное ли число.

package main

import "fmt"

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func main() {
	myNum := 10
	fmt.Println(myNum, "is even?", isEven(myNum))
}
