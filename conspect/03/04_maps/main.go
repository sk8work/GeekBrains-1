package main

import (
	"fmt"
)

func main() {
	// just a sample - var x map[string]int

	// Создаем и изменяем карты, функция delete
	addressBook := make(map[string][]int)
	addressBook["Eugene"] = []int{89520321452}
	fmt.Println(addressBook)
	addressBook["Eugene"] = []int{32135512352}
	addressBook["Eugene"] = append(addressBook["Eugene"], 6543213135)
	addressBook["Alex"] = []int{321353235132}
	fmt.Println(addressBook)
	// delete(addressBook, "Eugene")
	// fmt.Println(addressBook)

	// Итерируемся по значениям
	for name, numbers := range addressBook {
		fmt.Println("Abonent:", name)
		for _, number := range numbers {
			fmt.Printf("\t %v\n", number)
		}
	}
}
