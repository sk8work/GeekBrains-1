package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Name string `json:"name"`
	Body string
	Time int64 `json: ommitempty`
}

func main() {
	jsonStr := "{\"Name\": \"Alice\", \"Body\": \"Hello\", \"Time\": 1294706395881547000}"

	b := []byte(jsonStr)
	m := Message{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%v", m)

	fmt.Println(m.Name)
}
