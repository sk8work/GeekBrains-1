// Написать функцию, которая определяет, делится ли
// число без остатка на 3

package main

import "fmt"

func isDivBy3(a int) bool {
	if a%3 == 0 {
		return true
	}
	return false
}

func main() {
	myNum := 12
	fmt.Println(myNum, "is div by 3?", isDivBy3(myNum))
}
