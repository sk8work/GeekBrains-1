// массивы и слайсы

package main

import "fmt"

func main() {
	// sl1 := []int{1, 2, 3, 4, 5}
	// fmt.Println(sl1, len(sl1), cap(sl1))

	// for i, v := range sl1 {
	// 	fmt.Println(i, v)
	// }

	// sl1 = append(sl1, 6)
	// fmt.Println(sl1, len(sl1), cap(sl1))

	// for i, v := range sl1 {
	// 	fmt.Println(i, v)
	// }

	sl1 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl1, len(sl1), cap(sl1))

	sl1 = append(sl1, 10)

	fmt.Println(sl1, len(sl1), cap(sl1))

	for i, v := range sl1 {
		fmt.Println(i, v)
	}
}
