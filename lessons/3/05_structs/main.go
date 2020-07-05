package main

import (
	"GeekBrains-1/lessons/3/05_structs/structs"
	"fmt"
)

func main() {
	s := structs.Person{}
	fmt.Println(s)
	fmt.Printf("%+v", s)

	s.Name = "Vasya"
	s.SetSurname("Pupkin")
	fmt.Println(s)
	fmt.Printf("%+v", s)
}
