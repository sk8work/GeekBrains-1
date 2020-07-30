package main

import (
	"fmt"
	"sort"
)

// Len, Less, Swap

type ToSort struct {
	A int
}

type ToSortArray []ToSort

func (t ToSortArray) Len() int {
	return len(t)
}

func (t ToSortArray) Less(i, j int) bool {
	return t[i].A < t[j].A
}

func (t ToSortArray) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func main() {
	// t := make([]ToSort, 0, 4)
	// for i := 0; i < 4; i++ {
	// 	t = append(t, ToSort{
	// 		A: i + 5,
	// 	})
	// }
	t := []ToSort{}
	t = append(t, ToSort{A: 5})
	t = append(t, ToSort{A: 1})
	t = append(t, ToSort{A: 55})
	t = append(t, ToSort{A: 12})
	t = append(t, ToSort{A: 6})
	t = append(t, ToSort{A: 3})
	fmt.Println(t)
	sort.Sort(ToSortArray(t))
	fmt.Println(t)
}
