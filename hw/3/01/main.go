// Описать несколько структур — любой легковой автомобиль и грузовик.
// Структуры должны содержать марку авто, год выпуска, объем
// багажника/кузова, запущен ли двигатель, открыты ли окна, насколько
// заполнен объем багажника.

package main

import (
	"GeekBrains-1/hw/3/01/vehicle_struct"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	vehicle1 := vehicle_struct.Vehicle{
		Year:           1995,
		Name:           "Scania",
		ValueOfBaggage: 300,
		EngineStart:    true,
		IsWindowsOpen:  false,
		IsBagageFull:   false,
	}
	fmt.Println(vehicle1)
	jsonedVehicle1, err := json.Marshal(vehicle1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(jsonedVehicle1)

	v := vehicle_struct.Vehicle{}
	vehicle1ToString := string(jsonedVehicle1)
	fmt.Println(vehicle1ToString)

	err = json.Unmarshal(jsonedVehicle1, &v)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(v)

	// Принцип понятен!! Описывать еще структуры нет необходимости
}
