package main

import (
	"fmt"
	"os"
)

// пустому интерфейсу соответствуют все типы данных
// func​ ​Println​(a ...​interface​{}) (n ​int​, err error) {
// 	return​ Fprintln(os.Stdout, a...)
// }

func print(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}

func main() {
	print("Hello")
}
