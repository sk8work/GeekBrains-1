package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Name string
	Body string
	Time int
}

func main() {
	m1 := Message{"Alice", "Hello", 1234567893218351353}

	marshal, err := json.Marshal(m1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(marshal)
	fmt.Println(string(marshal))
}
