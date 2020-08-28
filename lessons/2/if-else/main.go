// ветвление

package main

import "fmt"

func f() bool {
	return true
}

func main() {
	i := 2

	if i%2 == 0 {
		fmt.Println("Hello")
	} else {
		fmt.Println("wrld")
	}

	if f() == true {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}
