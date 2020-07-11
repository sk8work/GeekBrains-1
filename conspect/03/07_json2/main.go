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
	jsonString := "{\"Name\": \"Alice\", \"Body\": \"Hello\", \"Time\": 129470639588154700}"
	b := []byte(jsonString)
	m := Message{}

	err1 := json.Unmarshal(b, &m)
	if err1 != nil {
		log.Println(err1)
		return
	}
	fmt.Println(m)
	fmt.Println(string(b))
}
