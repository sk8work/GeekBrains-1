// * Реализовать очередь. Это структура данных, работающая по принципу FIFO
// (First Input first output, или «первым зашел — первым вышел»).

package main

import (
	"GeekBrains-1/hw/3/03/fifo"
	"fmt"
)

func main() {
	fifo.Push("Text1")
	fifo.Push("Text2")
	fifo.Push("Text3")
	fmt.Println(fifo.Unshift())
	fmt.Println(fifo.Unshift())
	fifo.Push("Text4")
	fmt.Println(fifo.Unshift())
	fmt.Println(fifo.Unshift())
	fmt.Println(fifo.Unshift())

	fmt.Println(fifo.Unshift())

}
