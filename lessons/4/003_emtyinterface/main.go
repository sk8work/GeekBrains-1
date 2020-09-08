package main

import (
	"fmt"
	"sort"
)

func accept(iface ...interface{}) {
	fmt.Printf("%T\n", iface)
}

// type ToSort struct {
// 	A int
// }

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	accept(uint(8))
	accept("some str")
	accept(struct{}{})

	sl := []int{2, 5, 4, 8, 7, 9, 3, 5}
	fmt.Println(sl)
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] > sl[j]
	})
	fmt.Println(sl)
}
