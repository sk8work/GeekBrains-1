package main

import "fmt"

func main() {
	sl := make([]int, 0)
	fmt.Println(sl, len(sl), cap(sl))

	sl1 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl1, len(sl1), cap(sl1))
	sl1 = append(sl1, 6)
	fmt.Println(sl1, len(sl1), cap(sl1))
}
