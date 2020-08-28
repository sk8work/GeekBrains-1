// Инкапсуляция
// Инкапсуляция в языке Go доступна на уровне пакетов. Методы, которые находятся в импортируемом             пакете,
// будут доступны, только если их названия начинаются с большой буквы.

package main

import (
	"fmt"
	"math"
)

// Circle - простая структура дял описания круга
type Circle struct {
	radius int
}

// Area - метод, который получает структуру по значению
// и возвращает ее площадь
func (c Circle) Area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

// ChangeRadius - метод, который может изменить
// значение радиуса в конкретной структуре Circle
func (c *Circle) ChangeRadius(radius int) {
	c.radius = radius
}

func main() {
	c1 := Circle{5}
	fmt.Println(c1, &c1)
	c1.ChangeRadius(10)
	fmt.Println(c1, &c1)
	fmt.Println(c1.Area())
	c2 := Circle{6}
	fmt.Println(c2, &c2)
	c2.ChangeRadius(8)
	fmt.Println(c2, &c2)
	fmt.Println(c2.Area())
}
