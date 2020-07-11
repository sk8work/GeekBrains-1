package main

import "fmt"

func main() {
	var sl []int
	fmt.Printf("%v\n", sl)

	sl1 := make([]int, 5, 10)
	fmt.Printf("%v", sl1)
	fmt.Println(len(sl1), cap(sl1))

	// Дополнение существующего среза
	sl1 = append(sl1, 6, 7, 8)
	fmt.Println(sl1, len(sl1), cap(sl1))

	// Создание нового среза на основе старого с дополнительными элементами
	sl2 := append(sl1, 10, 10)
	fmt.Println(sl2, len(sl2), cap(sl2))

	// copy Нужно создать слайс такой же len и cap
	sl3 := make([]int, len(sl1), cap(sl1))
	fmt.Println(sl3, len(sl3), cap(sl3))
	copy(sl3, sl1)
	fmt.Println(sl1, len(sl1), cap(sl1))
	fmt.Println(sl3, len(sl3), cap(sl3))
	sl1[0] = 666
	fmt.Println(sl1, len(sl1), cap(sl1))
	fmt.Println(sl3, len(sl3), cap(sl3))
	sl3[1] = 666
	fmt.Println(sl1, len(sl1), cap(sl1))
	fmt.Println(sl3, len(sl3), cap(sl3))
}
