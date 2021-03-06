// Интерфейс - это совокупность методов для
// взаимодействия между разными типами данных.

//     Удовлетворение интерфейсу:
// тип данных удовлетворяет интерфейсу в случаях:
// Для структуры описаны методы, названия которых совпадают с названиями в интерфейсе
// Методы имеют входные и выходные данные согласно объявлениюв интерфейсе.

package main

import "fmt"

type Incrementer interface {
	Inc() error
}

type A struct {
	field int
}

func (a *A) Inc() error {
	a.field = a.field + 1
	return nil
}

type B struct {
	field2 uint32
}

func (b *B) Inc() error {
	b.field2 += 3
	if b.field2 > 2 {
		return fmt.Errorf("Got invalid field2: %d", b.field2)
	}
	return nil
}

func increment(inc Incrementer) error {
	return inc.Inc()
}

type IncrementerAbstractFactory struct {
}

func (iaf *IncrementerAbstractFactory) create(typ int) Incrementer {
	if typ == 1 {
		return &A{}
	}
	return &B{}
}

func main() {
	a := A{}
	b := &B{}

	increment(&a)
	increment(b)

	fmt.Println(a)
	fmt.Println(b)

	factory := &IncrementerAbstractFactory{}
	aa := factory.create(5)
	aa.Inc()
	fmt.Println(aa)

	iaf := &IncrementerAbstractFactory{}
	bb := iaf.create(2)
	err := bb.Inc()
	if err != nil {
		fmt.Println(err)
	}
}
