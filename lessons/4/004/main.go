package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Field string
}

func (a A) MarshalJSON() ([]byte, error) {
	return []byte(`{"field": "Not Some Str!!!"}`), nil
}

func main() {
	a := A{
		Field: "Some str",
	}
	data1, _ := json.Marshal(a)
	fmt.Println(string(data1))

	b := A{
		Field: "Some strasgsddgs",
	}
	data2, _ := json.Marshal(b)
	fmt.Println(string(data2))
}
