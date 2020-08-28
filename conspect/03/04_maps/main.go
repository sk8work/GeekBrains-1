package main

import (
	"fmt"
)

func main() {
	var x map[string]int
	fmt.Println(x)

	addressBook := make(map[string][]int)
	addressBook["firstAbonent"] = append(addressBook["firstAbonent"], 89991234567)
	addressBook["firstAbonent"] = append(addressBook["firstAbonent"], 87775648912)

	addressBook["secondAbonent"] = append(addressBook["secondAbonent"], 89993215465)
	fmt.Println(addressBook)
	for i, _ := range addressBook {
		fmt.Println(i)
		for _, k := range addressBook[i] {
			fmt.Println("\t", k)
		}
	}
}
