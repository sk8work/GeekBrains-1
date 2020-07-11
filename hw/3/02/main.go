// Инициализировать несколько экземпляров структур.
// Применить к ним различные действия.
// Вывести значения свойств экземпляров в консоль.

package main

import (
	"GeekBrains-1/hw/3/02/car_struct"
	"fmt"
)

func main() {
	car1 := car_struct.Car{
		Year:           2013,
		Name:           "Ford",
		ValueOfBaggage: 75,
		EngineStart:    false,
		IsWindowsOpen:  false,
		IsBagageFull:   true,
	}
	fmt.Println(car1)

	car1.EngineStart = true
	fmt.Println(car1)

	// Принцип ясен
}
