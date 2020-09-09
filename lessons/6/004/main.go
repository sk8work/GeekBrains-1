package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := struct {
		A string
		B string
	}{
		A: "42",
		B: "42",
	}
	typ := reflect.TypeOf(t)
	fmt.Println(typ)
	fmt.Println(typ.String())

	field, ok := typ.FieldByName("B")
	// fmt.Println("ok", ok)
	if ok {
		fmt.Println(field.PkgPath, ok)
	}
}
