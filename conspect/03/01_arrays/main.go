package main

import "fmt"

func main() {
	// Объявление массива со значениями
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	// Объявление пустого массива
	var arr1 [5]int
	fmt.Println(arr1)
	arr1[2] = 10
	fmt.Println(arr1)

	// Итерируемся по массиву
	for i, v := range arr {
		fmt.Printf("%v - это %v элемент массива \n", v, i)
	}
}
