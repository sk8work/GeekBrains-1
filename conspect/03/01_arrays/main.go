package main

import "fmt"

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)
	var arr2 [5]int
	fmt.Println(arr2)
	arr2[2] = 10
	fmt.Println(arr2)

	for i, v := range arr {
		fmt.Println(i, v)
	}
}
