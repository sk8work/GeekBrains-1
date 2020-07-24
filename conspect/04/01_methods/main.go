// Инкапсуляция
// Инкапсуляция в языке Go доступна на уровне пакетов. Методы, которые находятся в импортируемом             пакете,
// будут доступны, только если их названия начинаются с большой буквы.

package main

import (
	"fmt"
	"math"
)

// Circle - простая структура дял оописаня круга
type Circle struct {
	radius int
}

// Area - метод, который получает структуру по значению
// и возвращает ее площадь
func (c Circle) Area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

// ChangeRadius - метод, кооторый может изменить
// значение радиуса в конкретной структуре Circle
func (c *Circle) ChangeRadius(radius int) {
	c.radius = radius
}

func main() {
	c1 := Circle{5}
	fmt.Println(c1)
	fmt.Println(c1.Area())
	c1.ChangeRadius(2)
	fmt.Println(c1.Area())

}
