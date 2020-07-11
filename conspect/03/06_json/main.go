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

	stringMarshal := string(marshal)
	fmt.Println(stringMarshal)
	byteStringMarshal := []byte(stringMarshal)
	fmt.Println(byteStringMarshal)

	m2 := Message{}
	err1 := json.Unmarshal(byteStringMarshal, &m2)
	if err1 != nil {
		log.Println(err1)
		return
	}
	fmt.Println(m2)
	// fmt.Println(string(b))
}
