package main

import "fmt"

type phones []int

func (p phones) ViewList() {
	for i, phone := range p {
		fmt.Printf("\t %v) %v \n", i, phone)
	}
}

func main() {
	adressBook := make(map[string]phones)

	adressBook["Миша"] = phones{86432183}
	adressBook["Никита"] = phones{321351353, 45135138}

	for name, ph := range adressBook {
		fmt.Println(name)
		ph.ViewList()
	}
	fmt.Println(adressBook)
}
