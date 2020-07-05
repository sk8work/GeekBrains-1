// Написать функцию, которая последовательно
// выводит на экран 100 первых чисел Фибоначчи,
// начиная от 0.
// Числа Фибоначчи определяются
// соотношениями Fn=Fn-1 + Fn-2.

package main

import "fmt"

func printFibo(n int) {
	num1 := 0
	num2 := 1
	temp := 0
	for i := 0; i <= n; i++ {
		if i == 0 {
			fmt.Println(0)
			continue
		}
		temp = num1
		num1 = num2
		num2 = temp + num1
		fmt.Println(num1)
	}
}

func main() {
	printFibo(93)
}
