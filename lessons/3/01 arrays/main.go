package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println(arr, len(arr), cap(arr))

	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", arr1)

	sl := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", sl)

	emptyArr := [5]int{}
	fmt.Printf("%T\n", emptyArr)
}
