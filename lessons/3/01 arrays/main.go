package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println(arr, len(arr), cap(arr))

	for i, v := range arr {
		fmt.Println(i, "\t", v)
	}
	for i := range arr {
		arr[i] = i
	}

	for i, v := range arr {
		fmt.Println(i, "\t", v)
	}

	// arr1 := [5]int{1, 2, 3, 4, 5}
	// fmt.Printf("%T\n", arr1)

	// sl := []int{1, 2, 3, 4, 5}
	// fmt.Printf("%T\n", sl)

	// emptyArr := [5]int{}
	// fmt.Printf("%T\n", emptyArr)

	for j, v := range [10]int{} {
		fmt.Println(j, v)
	}

	sl := make([]struct{}, 0, 0)
	fmt.Printf("%T ,%v, %v", sl, len(sl), cap(sl))
	// sl = append(sl, [arr])

}
