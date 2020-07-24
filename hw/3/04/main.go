// * Внести в телефонный справочник дополнительные данные. Реализовать сохранение
// json-файла на диске с помощью ​пакета ioutil​ при изменении данных.
// Сделал не совсем по заданию, есть нкоторое непонимание процесса. Надо будет вернуться позже!!!!!!

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type JSONPhoneNums struct {
	book struct{}
}
type PhoneNums struct {
	Name string
	Nums interface{}
}

func main() {
	eugene := PhoneNums{}
	alex := PhoneNums{}

	eugene.Name = "Eugene"
	eugene.Nums = []int{9630958098, 9222128133}
	fmt.Println(eugene)

	alex.Name = "Alex"
	alex.Nums = []int{9992322545}
	fmt.Println(alex)

	jsonEugene, err := json.Marshal(eugene)
	if err != nil {
		log.Println(err)
	}
	jsonAlex, err := json.Marshal(alex)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonEugene))
	fmt.Println(string(jsonAlex))

	file, err := os.Create("nums.json")

	if err != nil {
		log.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString(string(jsonEugene) + string(jsonAlex))

	fmt.Println("Done.")
}
