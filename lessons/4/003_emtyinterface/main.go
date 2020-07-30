package main

import "fmt"

func accept(iface ...interface{}) {
	fmt.Printf("%T\n", iface)
}

func main() {
	accept(uint(8))
	accept("some str")
	accept(struct{}{})
}
