package main

import (
	"fmt"
	"reflect"
)

// func somefunc() bool {
// 	return true
// }

// const (
// 	a = iota << 10
// 	b = iota << 10
// 	c = iota << 10
// 	d = iota << 10
// 	e = iota << 10
// )

type SomeStruct struct {
	SomeInt      int
	someOtherInt int
}

func (s *SomeStruct) SetSomeOtherInt(v int) {
	s.someOtherInt = v
}

func main() {

	a := 5
	fmt.Println(reflect.TypeOf(a))
	b := "hello"
	fmt.Println(reflect.TypeOf(b))
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	// s := SomeStruct{}
	// fmt.Println(s)
	// s.SomeInt = 5
	// s.SetSomeOtherInt(10)
	// fmt.Println(s)
	// var a *int = nil
	// fmt.Println(&a)
	// fmt.Println(a)
	// newVar := 5
	// a = &newVar
	// fmt.Println(*a)
	// newVar = 10
	// fmt.Println(*a)

	// arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	// fmt.Println(json.Marshal(arr))
	// str, err := json.Marshal(arr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(str))

	// fmt.Println(packet.MyString)
	// runtime.GOMAXPROCS(1)

	// done := false

	// done = somefunc()
	// for !done {
	// 	fmt.Println("Not done")
	// }
	// fmt.Println("finished")

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	qqq := make([]*int, 0, 0)
	qq1 := 5
	qq2 := 10
	qqq = append(qqq, &qq1)
	qqq = append(qqq, &qq2)
	ptr := reflect.ValueOf(qqq)
	fmt.Println(&ptr)
	fmt.Println(qqq)
	for i, v := range qqq {
		fmt.Println(i, v, *v)
		fmt.Println(reflect.TypeOf(i), reflect.TypeOf(v), reflect.TypeOf(*v))
	}
}
