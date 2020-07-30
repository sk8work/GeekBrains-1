package main

import "fmt"

type SomeStruct struct {
	SomeInteger             int
	someNoneExportedInteger int
}

func (s *SomeStruct) SetSomeInt(value int) {
	s.SomeInteger = value
}

func (s SomeStruct) SetSomeIntCopy(value int) {
	s.SomeInteger = value
}

func main() {
	ss := SomeStruct{}
	ss.SetSomeInt(42)
	fmt.Println(ss)
	ss.SetSomeIntCopy(35)
	fmt.Println(ss)
}
