package main

import (
	"GeekBrains-1/lessons/3/06_json/structs"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	s := structs.Person{}
	s.Name = "name"
	s.Surname = "surname"
	data, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	fmt.Println(string(data))
}
