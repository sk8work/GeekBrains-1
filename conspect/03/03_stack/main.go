package main

import (
	"GeekBrains-1/conspect/03/03_stack/stack"
	"fmt"
)

func main() {
	stack.Push("This text")
	stack.Push("Will be in stack")
	stack.Push("Untill the first call in 'pop'")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	stack.Push("Enter some else text")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
