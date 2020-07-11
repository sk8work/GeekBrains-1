package main

import (
	"GeekBrains-1/conspect/03/03_stack/stack"
	"fmt"
)

func main() {
	stack.Push("text1")
	stack.Push("text2")
	stack.Push("text3")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	stack.Push("text4")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
