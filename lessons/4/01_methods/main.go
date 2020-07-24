package main

import "fmt"

type SomeStruct struct {
	SomeInteger        int
	someNonExportedInt int
}

// Геттеры мы устанавливаеv по указателю!!!!
func (s *SomeStruct) SetSomeInteger(value int) {
	s.SomeInteger = value
}

func (s SomeStruct) SetSomeIntegerCopy(value int) {
	s.SomeInteger = value
}

func main() {
	ss := SomeStruct{}
	ss.SetSomeInteger(42)
	ss.SetSomeIntegerCopy(35)
	fmt.Println(ss)
}
