package main

import (
	"fmt"
	"hash/crc32"
)

type hashmap [16]int

func hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func findPosition(key string, length int) int {
	h := hash(key)
	return int(h % uint32(length))
}

func main() {
	sl := make([]int, 4)
	key := "some_key123"
	value := 42
	position := findPosition(key, len(sl))
	fmt.Println(position)
	fmt.Println(sl)
	sl[position] = value
	fmt.Println(position)
	fmt.Println(sl)
}
