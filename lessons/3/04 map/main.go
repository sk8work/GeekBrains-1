package main

import "fmt"

func main() {
	mp := make(map[string]int)
	mp["some_key"] = 42
	fmt.Println(mp, len(mp))
}
