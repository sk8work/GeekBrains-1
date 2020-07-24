package main

import "fmt"

type Incrementer interface {
	Inc() error
}

type A struct {
	field int
}

// по указателю
func (a *A) Inc() error {
	a.field = a.field + 1
	return nil
}

type B struct {
	field2 uint32
}

func (b *B) Inc() error {
	b.field2 += 2
	if b.field2 > 2 {
		return fmt.Errorf("Got invalid field2: %d", b.field2)
	}
	return nil
}

func increment(inc Incrementer) error {
	return inc.Inc()
}

func main() {
	a := &A{}
	b := &B{}

	increment(a)
	increment(b)

	fmt.Println(a)
	fmt.Println(b)

}
